package playlist

import (
	"fmt"
	goplaylist "mini-clean"
	"mini-clean/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostgresRepository struct {
	db *gorm.DB
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&entities.Playlist{})
	err := db.SetupJoinTable(&entities.Playlist{}, "Musics", &entities.PlaylistMusic{})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&entities.PlaylistMusic{})
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	Migrate(db)
	return &PostgresRepository{
		db: db,
	}
}

func (repo *PostgresRepository) ExistCollab(userId uint64, playlistId uint64) (playlist *entities.Playlist, err error) {
	fmt.Println(userId, playlistId)
	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return nil, goplaylist.ErrInternalServer
	}
	ent := &entities.Collaboration{}
	err = opr.Debug().Where("playlist_id = ?", playlistId).First(ent, "user_id = ?", userId).Error
	fmt.Println(err, ent)
	if err != nil {
		err = goplaylist.ErrNotFound
		return
	}

	opr.Commit()

	return
}

func (repo *PostgresRepository) FindById(id uint64) (playlist *entities.Playlist, err error) {

	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return nil, goplaylist.ErrInternalServer
	}

	err = opr.Debug().Preload("Users", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("id", "name", "email")
	}).Preload("Musics").First(&playlist, id).Error

	if err != nil {
		err = goplaylist.ErrNotFound
		return
	}

	opr.Commit()

	return
}

func (repo *PostgresRepository) FindAll() (playlist []entities.Playlist, err error) {
	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return nil, goplaylist.ErrInternalServer
	}

	err = opr.Preload("Musics").Find(&playlist).Error
	if err != nil {
		err = goplaylist.ErrNotFound
		return
	}

	opr.Commit()
	return
}

func (repo *PostgresRepository) FindByQuery(key string, value interface{}) (playlist []entities.Playlist, err error) {
	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return playlist, goplaylist.ErrInternalServer
	}

	err = opr.Debug().Find(&playlist, key+" = ?", value).Error
	if err != nil || len(playlist) == 0 {
		err = goplaylist.ErrNotFound
		return
	}

	opr.Commit()

	return
}

func (repo *PostgresRepository) Insert(data entities.Playlist) (id uint64, err error) {

	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		err = goplaylist.ErrInternalServer
		return
	}

	err = opr.Create(&data).Error

	if err != nil {
		// internal server error
		err = goplaylist.ErrInternalServer
		return
	}
	id = data.ID
	opr.Commit()

	return
}

func (repo *PostgresRepository) Update(data entities.Playlist) (playlist *entities.Playlist, err error) {

	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return nil, goplaylist.ErrInternalServer
	}

	opr.First(&playlist, data.ID)
	if err != nil {
		err = goplaylist.ErrNotFound
		return
	}
	opr.Model(&playlist).Omit("ID", "email").Updates(map[string]interface{}{"name": data.Name})
	if err != nil {
		// internal server error
		err = goplaylist.ErrInternalServer
		return
	}

	opr.Commit()

	return
}

func (repo *PostgresRepository) Delete(id uint64) (err error) {
	err = repo.db.Select(clause.Associations).Where("id = ?", id).Delete(&entities.Playlist{ID: id}).Error
	if err != nil {
		err = goplaylist.ErrNotFound
		return
	}
	return
}

func (repo *PostgresRepository) AddPlaylistMusic(data entities.PlaylistMusic) (err error) {
	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return goplaylist.ErrInternalServer
	}

	err = opr.Debug().Create(&data).Error

	if err != nil {
		err = goplaylist.ErrNotFound
		return
	}

	opr.Commit()

	return
}

func (repo *PostgresRepository) FindPlaylistMusicById(playlistId uint64) (playlistMusics entities.Playlist, err error) {
	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		err = goplaylist.ErrInternalServer
		return
	}

	err = opr.Preload("Musics").First(&playlistMusics, playlistId).Error

	if err != nil {
		err = goplaylist.ErrNotFound
		return
	}

	opr.Commit()

	return
}

func (repo *PostgresRepository) DeletePlaylistMusicById(musicId uint64, playlistId uint64) (err error) {
	err = repo.db.Debug().Where("playlist_id = ?", playlistId).Where("music_id = ?", musicId).Delete(&entities.PlaylistMusic{}).Error
	if err != nil {
		err = goplaylist.ErrNotFound
		return
	}
	return
}

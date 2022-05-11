package playlist

import (
	"fmt"
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

func (repo *PostgresRepository) FindById(id uint64) (playlist *entities.Playlist, err error) {

	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return nil, err
	}

	opr.Preload("Musics").Preload("Users").First(&playlist, id)

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
		return nil, err
	}

	err = opr.Preload("Musics").Preload("Users").Find(&playlist).Error
	if err != nil {
		return
	}

	opr.Commit()
	return
}

func (repo *PostgresRepository) FindByQuery(key string, value interface{}) (playlist entities.Playlist, err error) {
	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return playlist, err
	}

	err = opr.Preload("Musics").Preload("Users").Where(key+" = ?", value).Find(&playlist).Error
	if err != nil {
		return
	}

	opr.Commit()

	return
}

func (repo *PostgresRepository) Insert(data entities.Playlist) (err error) {

	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return err
	}

	err = opr.Create(&data).Error

	if err != nil {
		return
	}

	opr.Commit()

	return
}

func (repo *PostgresRepository) Update(data entities.Playlist) (playlist *entities.Playlist, err error) {

	opr := repo.db.Begin()

	if err != nil {
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return nil, err
	}

	opr.First(&playlist, data.ID)

	opr.Model(&playlist).Omit("ID", "email").Updates(map[string]interface{}{"name": data.Name})

	opr.Commit()

	return
}

func (repo *PostgresRepository) Delete(id uint64) (err error) {
	err = repo.db.Select(clause.Associations).Where("id = ?", id).Delete(&entities.Playlist{ID: id}).Error
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
		return
	}

	err = opr.Create(&data).Error

	if err != nil {
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
		return
	}

	err = opr.First(&playlistMusics, playlistId).Preload("Musics").Error

	if err != nil {
		return
	}

	opr.Commit()

	return
}

func (repo *PostgresRepository) DeletePlaylistMusicById(musicId uint64, playlistId uint64) (err error) {
	err = repo.db.Where("playlist_id = ?", playlistId).Where("music_id = ?", musicId).Delete(&entities.PlaylistMusic{}).Error
	return
}

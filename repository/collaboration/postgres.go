package collaboration

import (
	"mini-clean/entities"
	goplaylist "mini-clean/error"

	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func Migrate(db *gorm.DB) {
	err := db.SetupJoinTable(&entities.Playlist{}, "Users", &entities.Collaboration{})
	if err != nil {

	}
	db.AutoMigrate(&entities.Collaboration{})
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	Migrate(db)
	return &PostgresRepository{
		db: db,
	}
}

func (repo *PostgresRepository) Exist(userId uint64, playlistId uint64) (collaboration *entities.Collaboration, err error) {
	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return nil, err
	}

	err = opr.Where("user_id = ?", userId).Where("playlist_id").Find(&collaboration).Error
	if err != nil {
		err = goplaylist.ErrNotFound
		return
	}

	opr.Commit()

	return
}

func (repo *PostgresRepository) FindById(id uint64) (collaboration *entities.Collaboration, err error) {

	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return nil, err
	}

	err = opr.First(&collaboration, id).Error
	if err != nil {
		err = goplaylist.ErrNotFound
		return
	}

	opr.Commit()

	return
}

func (repo *PostgresRepository) FindAll() (users []entities.Collaboration, err error) {
	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return nil, goplaylist.ErrInternalServer
	}

	err = opr.Find(&users).Error
	if err != nil {
		err = goplaylist.ErrNotFound
		return
	}
	opr.Commit()
	return
}

func (repo *PostgresRepository) FindByQuery(key string, value interface{}) (collaboration entities.Collaboration, err error) {
	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return collaboration, err
	}

	err = opr.Where(key+" = ?", value).Find(&collaboration).Error
	if err != nil {
		err = goplaylist.ErrNotFound
		return
	}
	opr.Commit()

	return
}

func (repo *PostgresRepository) Insert(data entities.Collaboration) (err error) {

	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		err = goplaylist.ErrInternalServer
		return err
	}

	err = opr.Create(&data).Error
	if err != nil {
		err = goplaylist.ErrInternalServer
		return
	}

	opr.Commit()

	return
}

func (repo *PostgresRepository) Delete(userId uint64, playlistId uint64) (err error) {
	var check entities.Collaboration
	err = repo.db.First(&check, "user_id = ?", userId).Error
	if err != nil {
		err = goplaylist.ErrNotFound
		return
	}
	err = repo.db.Where("playlist_id = ?", playlistId).Where("user_id = ?", userId).Delete(&entities.Collaboration{}).Error
	if err != nil {
		err = goplaylist.ErrInternalServer
		return
	}
	return
}

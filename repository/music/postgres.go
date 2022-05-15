package music

import (
	"fmt"
	"mini-clean/entities"
	goplaylist "mini-clean/error"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostgresRepository struct {
	db *gorm.DB
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&entities.Music{})
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	Migrate(db)
	return &PostgresRepository{
		db: db,
	}
}

func (repo *PostgresRepository) FindById(id uint64) (music *entities.Music, err error) {

	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return nil, goplaylist.ErrInternalServer
	}

	err = opr.First(&music, id).Error
	if err != nil {
		err = goplaylist.ErrNotFound
		return
	}

	opr.Commit()

	return
}

func (repo *PostgresRepository) FindAll() (musics []entities.Music, err error) {
	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return nil, goplaylist.ErrInternalServer
	}

	err = opr.Find(&musics).Error
	if err != nil {
		err = goplaylist.ErrInternalServer
		return
	}

	opr.Commit()
	return
}

func (repo *PostgresRepository) FindByQuery(key string, value interface{}) (music *entities.Music, err error) {
	opr := repo.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return nil, goplaylist.ErrInternalServer
	}

	err = opr.First(&music, key+" = ?", value).Error
	fmt.Println(music, key, value)
	if err != nil {
		err = goplaylist.ErrNotFound
		return
	}
	opr.Commit()

	return
}

func (repo *PostgresRepository) Insert(data entities.Music) (id uint64, err error) {

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
		err = goplaylist.ErrInternalServer
		return
	}
	opr.Commit()
	id = data.ID
	return
}

func (repo *PostgresRepository) Update(data entities.Music) (music *entities.Music, err error) {
	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return nil, goplaylist.ErrInternalServer
	}

	err = opr.First(&music, data.ID).Error

	if err != nil {
		return
	}

	err = opr.Model(&music).Omit("ID").Updates(map[string]interface{}{"title": data.Title}).Error
	if err != nil {
		err = goplaylist.ErrInternalServer
		return
	}
	opr.Commit()

	return
}

func (repo *PostgresRepository) Delete(id uint64) (err error) {
	music, er := repo.FindById(id)
	if er != nil {
		err = goplaylist.ErrNotFound
		return
	}
	err = repo.db.Select(clause.Associations).Where("id = ?", id).Delete(music).Error
	if err != nil {
		err = goplaylist.ErrInternalServer
		return
	}
	return
}

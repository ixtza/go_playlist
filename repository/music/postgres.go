package music

import (
	"fmt"
	"mini-clean/entities"

	"gorm.io/gorm"
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
		return nil, err
	}

	opr.First(&music, id)

	opr.Commit()

	return
}

func (repo *PostgresRepository) FindAll() (users []entities.Music, err error) {
	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return nil, err
	}

	opr.Find(&users)

	opr.Commit()
	return
}

func (repo *PostgresRepository) FindByQuery(key string, value interface{}) (music entities.Music, err error) {
	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return music, err
	}

	opr.Where(key+" = ?", value).Find(&music)

	opr.Commit()

	return
}

func (repo *PostgresRepository) Insert(data entities.Music) (err error) {

	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return err
	}

	opr.Create(&data)

	opr.Commit()

	return
}

func (repo *PostgresRepository) Update(data entities.Music) (music *entities.Music, err error) {
	fmt.Println(data)
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

	opr.First(&music, data.ID)

	opr.Model(&music).Omit("ID").Updates(map[string]interface{}{"title": data.Title})

	opr.Commit()

	return
}

func (repo *PostgresRepository) Delete(id uint64) (err error) {
	err = repo.db.Where("id = ?", id).Delete(&entities.Music{}).Error
	return
}

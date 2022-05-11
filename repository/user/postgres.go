package user

import (
	"mini-clean/entities"

	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&entities.User{})
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	Migrate(db)
	return &PostgresRepository{
		db: db,
	}
}

func (repo *PostgresRepository) FindById(id uint64) (user *entities.User, err error) {

	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return nil, err
	}

	err = opr.First(&user, id).Error

	opr.Commit()

	return
}

func (repo *PostgresRepository) FindAll() (users []entities.User, err error) {
	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return nil, err
	}

	err = opr.Find(&users).Error

	opr.Commit()
	return
}

func (repo *PostgresRepository) FindByQuery(key string, value interface{}) (user *entities.User, err error) {
	opr := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			opr.Rollback()
		}
	}()

	if err = opr.Error; err != nil {
		return user, err
	}

	err = opr.Where(key+" = ?", value).Find(&user).Error

	opr.Commit()
	return
}

func (repo *PostgresRepository) Insert(data entities.User) (err error) {

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

	opr.Commit()

	return
}

func (repo *PostgresRepository) Update(data entities.User) (user *entities.User, err error) {
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

	err = opr.First(&user, data.ID).Error

	err = opr.Model(&user).Omit("ID", "email").Updates(map[string]interface{}{"name": data.Name, "password": data.Password}).Error

	opr.Commit()

	return
}

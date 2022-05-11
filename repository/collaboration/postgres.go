package collaboration

import (
	"fmt"
	"mini-clean/entities"

	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func Migrate(db *gorm.DB) {
	err := db.SetupJoinTable(&entities.Playlist{}, "Users", &entities.Collaboration{})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&entities.Collaboration{})
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	Migrate(db)
	return &PostgresRepository{
		db: db,
	}
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

	opr.First(&collaboration, id)

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
		return nil, err
	}

	opr.Find(&users)

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

	opr.Where(key+" = ?", value).Find(&collaboration)

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
		return err
	}

	opr.Create(&data)

	opr.Commit()

	return
}

func (repo *PostgresRepository) Delete(userId uint64, collaborationId uint64) (err error) {
	err = repo.db.Where("collaboration_id = ?", collaborationId).Where("user_id = ?", userId).Delete(&entities.Collaboration{}).Error
	return
}

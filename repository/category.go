package repository

import (
	"repo/modal"

	"gorm.io/gorm"
)

type CategoryRepo struct {
	db *gorm.DB
}

func (rootRepo *Repositories) Category() CategoryRepo {
	return CategoryRepo{db: rootRepo.Db}
}
func (c CategoryRepo) Save(category *modal.Category) error {
	err := c.db.Model(&modal.Category{}).Save(&category)
	return err.Error
}

package repository

import (
	"repo/modal"
	"repo/modal/request"

	"gorm.io/gorm"
)

type CategoryRepo struct {
	db *gorm.DB
}

func (rootRepo *Repositories) Category() CategoryRepo {
	return CategoryRepo{db: rootRepo.Db}
}
func (c CategoryRepo) Save(request request.CategoryInsert) (*modal.Category, error) {
	categoryies := modal.Category{}
	categoryies.Name = request.Name
	save := c.db.Debug().Save(&categoryies)
	return &categoryies, save.Error
}

package services

import (
	"repo/modal"
	"repo/modal/request"
	"repo/repository"

	"github.com/labstack/echo/v4"
)

type CategoryServices struct{}

func CategoryS() CategoryServices {
	return CategoryServices{}
}
func (cs CategoryServices) Save(c *echo.Context) error {
	var rq request.CategoryInsert
	category := modal.Category{
		Name: rq.Name,
	}
	err := repository.Get().Category().Save(category)
	return err
}

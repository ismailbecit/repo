package controller

import (
	"net/http"
	"repo/modal/request"
	"repo/repository"

	"github.com/labstack/echo/v4"
)

func CategoryInsert(c echo.Context) error {
	var rq request.CategoryInsert
	if err := c.Bind(&rq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	newUser, _ := repository.Get().Category().Save(rq)
	return c.JSON(http.StatusOK, newUser)
}

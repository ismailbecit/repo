package controller

import (
	"net/http"
	"repo/services"

	"github.com/labstack/echo/v4"
)

func CategoryInsert(c echo.Context) error {
	err := services.CategoryS().Save(&c)
	return c.JSON(http.StatusOK, err)
}

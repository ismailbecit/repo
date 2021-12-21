package router

import (
	"repo/config"
	"repo/controller"

	"github.com/labstack/echo/v4"
)

func init() {
	config.Connet()
}
func Router() {

	e := echo.New()
	category := e.Group("/category")
	category.POST("/insert", controller.CategoryInsert)

	e.Start(":8080")
}

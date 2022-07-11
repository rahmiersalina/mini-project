package routers

import (
	"mini-project/controllers"

	"github.com/labstack/echo/v4"
)

func Router() *echo.Echo {
	e := echo.New()
	e.GET("/makanan", controllers.GetMakananscontrollers)
	e.GET("/makanan/:id", controllers.GetMakanancontrollers)
	e.POST("/makanan", controllers.InsertMakanan)

	return e
}


package routers

import (
	"net/http"

	"github.com/AnggaArdhinata/indochat/src/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	api := e.Group("api/v1")

	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Golang Back-End APP")

	})

	//? Product End Point
	api.GET("/product", controllers.GetAllProduct)
	api.POST("/product", controllers.StoreProduct)

	//? Order End Point
	api.GET("/order", controllers.GetOrder)

	return e
}

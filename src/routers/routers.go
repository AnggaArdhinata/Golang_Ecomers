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
	api.PUT("/product/:id", controllers.UpdateProduct)
	api.DELETE("/product/:id", controllers.DeleteProduct)

	//? Order End Point
	api.GET("/order", controllers.GetOrder)
	api.POST("/order", controllers.StoreOrder)
	api.PUT("/order/:id", controllers.UpdateOrder)
	api.DELETE("/order/:id", controllers.DeleteOrder)

	//? Customer End Point
	api.GET("/customer", controllers.GetAllCustomer)
	api.POST("/customer", controllers.StoreCustomer)
	api.PUT("/customer/:id", controllers.UpdateCustomer)
	api.DELETE("/customer/:id", controllers.DeleteCustomer)

	//? Category End Point
	api.GET("/category", controllers.GetAllCategory)
	api.POST("/category", controllers.StoreCategory)
	api.PUT("/category/:id", controllers.UpdateCategory)
	api.DELETE("/category/:id", controllers.DeleteCategory)

	return e
}

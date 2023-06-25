package routers

import (
	"net/http"

	"github.com/AnggaArdhinata/indochat/src/controllers"
	"github.com/AnggaArdhinata/indochat/src/middlewares"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	api := e.Group("api/v1")
	
	product := api.Group("/product", middlewares.RateLimit)
	order := api.Group("/order", middlewares.RateLimit)
	customer := api.Group("/customer", middlewares.RateLimit)
	category := api.Group("/category", middlewares.RateLimit)

	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Golang Back-End APP")

	})

	//? Product End Point
	product.GET("", controllers.GetAllProduct)
	product.POST("", controllers.StoreProduct)
	product.PUT("/:id", controllers.UpdateProduct)
	product.DELETE("/:id", controllers.DeleteProduct)

	//? Order End Pointcustomer
	order.GET("", controllers.GetOrder)
	order.POST("", controllers.StoreOrder)
	order.PUT("/:id", controllers.UpdateOrder)
	order.GET("/verify/:id", controllers.UpdatePayment)
	order.DELETE("/:id", controllers.DeleteOrder)

	//? Customer End Point
	customer.GET("", controllers.GetAllCustomer)
	customer.POST("", controllers.StoreCustomer)
	customer.PUT("/:id", controllers.UpdateCustomer)
	customer.DELETE("/:id", controllers.DeleteCustomer)

	//? Category End Point
	category.GET("", controllers.GetAllCategory)
	category.POST("", controllers.StoreCategory)
	category.PUT("/:id", controllers.UpdateCategory)
	category.DELETE("/:id", controllers.DeleteCategory)

	return e
}

package routers

import (
	"net/http"

	"github.com/AnggaArdhinata/indochat/src/controllers"
	"github.com/AnggaArdhinata/indochat/src/middlewares"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	api := e.Group("api/v1", middlewares.RateLimit)

	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Golang Back-End APP")

	})
	
	product := api.Group("/product")
	order := api.Group("/order")
	customer := api.Group("/customer")
	category := api.Group("/category")

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
	order.GET("/generate/xls", controllers.GetXls)
	order.DELETE("/:id", controllers.DeleteOrder)

	//? Customer End Point
	customer.GET("", controllers.GetAllCustomer)
	customer.GET("/email", controllers.GetByEmail)
	customer.POST("", controllers.StoreCustomer)
	customer.PUT("/:id", controllers.UpdateCustomer)
	customer.DELETE("/:id", controllers.DeleteCustomer)
	customer.DELETE("/delete", controllers.DeleteCustomerByEmail)

	//? Category End Point
	category.GET("", controllers.GetAllCategory)
	category.POST("", controllers.StoreCategory)
	category.PUT("/:id", controllers.UpdateCategory)
	category.DELETE("/:id", controllers.DeleteCategory)

	return e
}

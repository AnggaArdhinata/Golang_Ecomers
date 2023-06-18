package controllers

import (
	"net/http"

	"github.com/AnggaArdhinata/indochat/src/models"
	"github.com/labstack/echo/v4"
)

func GetOrder(c echo.Context) error {
	result, err := models.GetOrder()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

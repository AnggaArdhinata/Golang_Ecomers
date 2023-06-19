package controllers

import (
	"net/http"
	"strconv"

	"github.com/AnggaArdhinata/indochat/src/models"
	"github.com/labstack/echo/v4"
)

func StoreCategory(c echo.Context) error {
	name := c.FormValue("name")

	result, err := models.StoreCategory(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)

}

func GetAllCategory(c echo.Context) error {
	result, err := models.GetAllCategory()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateCategory(c echo.Context) error {
	str_id := c.Param("id")
	name := c.FormValue("name")

	id, err := strconv.Atoi(str_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateCategory(id, name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteCategory(c echo.Context) error {
	id_str := c.Param("id")

	id, err := strconv.Atoi(id_str)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteCategory(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
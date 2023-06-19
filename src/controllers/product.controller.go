package controllers

import (
	"net/http"
	"strconv"

	"github.com/AnggaArdhinata/indochat/src/models"
	"github.com/labstack/echo/v4"
)

func StoreProduct(c echo.Context) error {
	name := c.FormValue("name")
	str_category_id := c.FormValue("category_id")
	str_price := c.FormValue("price")
	description := c.FormValue("description")
	image := c.FormValue("image")

	category_id, err := strconv.Atoi(str_category_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	price, err := strconv.Atoi(str_price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StoreProduct(name, category_id, price, description, image)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)

}

func GetAllProduct(c echo.Context) error {
	result, err := models.GetAllProduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateProduct(c echo.Context) error {
	str_id := c.Param("id")
	name := c.FormValue("name")
	str_category_id := c.FormValue("category_id")
	str_price := c.FormValue("price")
	description := c.FormValue("description")
	image := c.FormValue("image")

	id, err := strconv.Atoi(str_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	category_id, err := strconv.Atoi(str_category_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	price, err := strconv.Atoi(str_price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateProduct(id, name, category_id, price, description, image)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteProduct(c echo.Context) error {
	id_str := c.Param("id")

	id, err := strconv.Atoi(id_str)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

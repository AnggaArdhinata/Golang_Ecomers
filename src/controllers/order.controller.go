package controllers

import (
	"net/http"
	"strconv"

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

func StoreOrder(c echo.Context) error {
	str_cust_id := c.FormValue("cust_id")
	str_product_id := c.FormValue("product_id")
	discount_code := c.FormValue("discount_code")

	cust_id, err := strconv.Atoi(str_cust_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	product_id, err := strconv.Atoi(str_product_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StoreOrder(cust_id, product_id, discount_code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)

}

func UpdateOrder(c echo.Context) error {
	str_id := c.Param("id")
	str_cust_id := c.FormValue("cust_id")
	str_product_id := c.FormValue("product_id")
	str_isPaid := c.FormValue("ispaid")
	status := c.FormValue("status")
	discount_code := c.FormValue("discount_code")

	id, err := strconv.Atoi(str_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cust_id, err := strconv.Atoi(str_cust_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	product_id, err := strconv.Atoi(str_product_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	isPaid, err := strconv.ParseBool(str_isPaid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateOrder(id, cust_id, product_id, isPaid, status, discount_code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteOrder(c echo.Context) error {
	id_str := c.Param("id")

	id, err := strconv.Atoi(id_str)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteOrder(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
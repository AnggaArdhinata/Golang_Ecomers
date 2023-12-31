package controllers

import (
	"net/http"
	"strconv"

	"github.com/AnggaArdhinata/indochat/src/models"
	"github.com/labstack/echo/v4"
)

func StoreCustomer(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	res, err := models.GetCustomerByEmail(email)
	if res.Status == 200 {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "email already registered !"})
	}
	result, err := models.StoreCustomer(name, email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)

}

func GetAllCustomer(c echo.Context) error {
	result, err := models.GetAllCustomer()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetByEmail(c echo.Context) error {

	email := c.QueryParam("email")

	result, err := models.GetCustomerByEmail(email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "email not found !"})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateCustomer(c echo.Context) error {
	str_id := c.Param("id")
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	id, err := strconv.Atoi(str_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateCustomer(id, name, email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteCustomer(c echo.Context) error {
	id_str := c.Param("id")

	id, err := strconv.Atoi(id_str)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteCustomer(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteCustomerByEmail(c echo.Context) error {
	var res models.Response
	email := c.QueryParam("email")

	checkEmail, err := models.GetCustomerByEmail(email)
	if checkEmail.Status != 200 {
		res.Status = http.StatusInternalServerError
		res.Message = "internal server error"
		res.Data = models.Msg{Msg: "email not found !"}
		return c.JSON(http.StatusInternalServerError, res)
	}
	result, err := models.DeleteByEmail(email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

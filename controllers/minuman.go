package controllers

import (
	"mini-project/config"
	"mini-project/helper"
	"mini-project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all minumans
func GetMinumanscontrollers(c echo.Context) error {
	var minumans []models.Makanan
	if err := config.DB.Table("minuman").Find(&minumans).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all minuman", minumans))
}

// get minumans by id
func GetMinumancontrollers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	minuman := models.Minuman{}
	if err := config.DB.Table("minuman").First(&minuman, id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "minuman not found",
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get minuman", minuman))
}

func InsertMinuman(c echo.Context) error {
	var input models.InputNewMinuman
	var minuman models.Minuman

	err := c.Bind(&input)
	if err != nil {
		return err
	}

	minuman.Nama_minuman = input.NamaMinuman
	minuman.Jenis = input.Jenis
	minuman.Harga = input.Harga
	minuman.Id_minuman = "2"

	err = config.DB.Table("minuman").Create(&minuman).Error
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "success")

}

package controllers

import (
	"mini-project/config"
	"mini-project/helper"
	"mini-project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all makanans
func GetMakananscontrollers(c echo.Context) error {
	var makanans []models.Makanan
	if err := config.DB.Table("makanan").Find(&makanans).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all makanan", makanans))
}

// get makanan by id
func GetMakanancontrollers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	makanan := models.Makanan{}
	if err := config.DB.Table("makanan").First(&makanan, id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "makanan not found",
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get makanan", makanan))
}

func InsertMakanan(c echo.Context) error {
	var input models.InputNewMakanan
	var makanan models.Makanan

	err := c.Bind(&input)
	if err != nil {
		return err
	}

	makanan.Nama_makanan = input.NamaMakanan
	makanan.Jenis = input.Jenis
	makanan.Harga = input.Harga
	makanan.Id_makanan = "2"

	err = config.DB.Table("makanan").Create(&makanan).Error
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "success")

}

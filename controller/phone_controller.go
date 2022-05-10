package controller

import (
	"kentang/config"
	"kentang/models"
	"net/http"

	"kentang/formatter"

	"github.com/labstack/echo/v4"
)

func GetPhonesController(c echo.Context) error {
	DB := config.Connect()
	var phones []models.Phone

	if err := DB.Find(&phones).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, formatter.SuccessResponse(phones))
}

func GetPhoneController(c echo.Context) error {
	DB := config.Connect()
	id := c.Param("id")
	phone := models.Phone{}

	DB.Where("ID = ?", id).First(&phone)

	if phone.ID == 0 {
		return c.JSON(http.StatusNotFound, formatter.NotFoundResponse(nil))
	}

	return c.JSON(http.StatusOK, formatter.SuccessResponse(phone))
}

func CreatePhoneController(c echo.Context) error {
	DB := config.Connect()
	phone := models.Phone{}
	c.Bind(&phone)

	if err := DB.Save(&phone).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, formatter.SuccessResponse(phone))
}

func DeletePhoneController(c echo.Context) error {
	DB := config.Connect()
	id := c.Param("id")

	DB.Delete(&models.Phone{}, id)

	return c.JSON(http.StatusOK, formatter.SuccessResponse(nil))
}

func UpdatePhoneController(c echo.Context) error {
	DB := config.Connect()
	id := c.Param("id")
	phone := models.Phone{}

	DB.Where("ID = ?", id).First(&phone)

	if phone.ID == 0 {
		return c.JSON(http.StatusNotFound, formatter.NotFoundResponse(nil))
	}

	payload := models.Phone{}
	c.Bind(&payload)

	phone.Name = payload.Name
	phone.Price = payload.Price
	phone.Design = payload.Design
	phone.Display = payload.Display
	phone.Perfomance = payload.Perfomance
	phone.Cameras = payload.Cameras
	phone.Audio = payload.Audio
	phone.Battery = payload.Battery
	phone.Features = payload.Features
	DB.Save(&phone)

	return c.JSON(http.StatusOK, formatter.SuccessResponse(phone))
}

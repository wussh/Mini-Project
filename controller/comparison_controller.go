package controller

import (
	"kentang/config"
	"kentang/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetComparisonController(c echo.Context) error {
	DB := config.Connect()
	comparison := models.Comparison{}
	c.Bind(&comparison)
	phonesatu := models.Phone{}
	phonedua := models.Phone{}
	DB.Where("name = ?", comparison.PhoneSatu).First(&phonesatu)
	DB.Where("name = ?", comparison.PhoneDua).First(&phonedua)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"phonesatu": phonesatu,
		"phonedua":  phonedua,
	})
}

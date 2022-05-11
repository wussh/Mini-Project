package controller

import (
	"kentang/config"
	"kentang/formatter"
	"kentang/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateCommentController(c echo.Context) error {
	DB := config.Connect()
	comment := models.Comment{}
	c.Bind(&comment)

	if err := DB.Save(&comment).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "success create comment",
		"usercomment": comment,
	})
}

func UpdateCommentController(c echo.Context) error {
	DB := config.Connect()
	id := c.Param("id")
	comment := models.Comment{}

	DB.Where("ID = ?", id).First(&comment)

	if comment.ID == 0 {
		return c.JSON(http.StatusNotFound, formatter.NotFoundResponse(nil))
	}

	payload := models.Comment{}
	c.Bind(&payload)

	comment.Context = payload.Context
	DB.Save(&comment)

	return c.JSON(http.StatusOK, formatter.SuccessResponse(comment))
}

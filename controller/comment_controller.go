package controller

import (
	"fmt"
	"kentang/config"
	"kentang/formatter"
	"kentang/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CommentRequest struct {
	Context string `json:"context" form:"context"`
	PhoneId int    `json:"phoneid" form:"phoneid"`
}

func CreateCommentController(c echo.Context) error {
	DB := config.Connect()
	var commentRequest CommentRequest
	c.Bind(&commentRequest)
	fmt.Printf("%#v\n", commentRequest)
	phoneId := commentRequest.PhoneId
	var phone models.Phone
	DB.Where("id = ?", phoneId).Find(&phone)
	fmt.Printf("%#v\n", phone)
	comment := models.Comment{
		Phone:   phone,
		Context: commentRequest.Context,
	}
	DB.Save(&comment)
	return c.JSON(http.StatusOK, comment)
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
func GetPhoneCommentController(c echo.Context) error {
	DB := config.Connect()
	var comment []models.Comment

	if err := DB.Find(&comment).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, formatter.SuccessResponse(comment))
}

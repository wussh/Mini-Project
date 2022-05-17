package controller

import (
	"fmt"
	"kentang/config"
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

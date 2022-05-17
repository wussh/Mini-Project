package routes

import (
	c "kentang/controller"
	"kentang/key"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	// Welcome Page
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Review HP")
	})

	e.POST("/auth/login", c.LoginController)

	e.POST("/handphone", c.CreatePhoneController)

	e.POST("/users", c.CreateUserController)

	e.GET("/handphone", c.GetPhonesController)
	e.GET("/handphone/:id", c.GetPhoneController)
	e.POST("/comparison", c.GetComparisonController)

	jwtAuth := e.Group("/jwt/redirected")
	jwtAuth.Use(middleware.JWT([]byte(key.SECRET_JWT)))

	jwtAuth.GET("/users", c.GetUsersController)
	jwtAuth.GET("/users/:id", c.GetUserController)
	jwtAuth.DELETE("/users/:id", c.DeleteUserController)
	jwtAuth.PUT("/users/:id", c.UpdateUserController)
	jwtAuth.POST("/comment", c.CreateCommentController)
	return e
}

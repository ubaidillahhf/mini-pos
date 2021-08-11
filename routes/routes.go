package routes

import (
	"pos/constants"
	"pos/controllers"
	"pos/middlewares"
	"pos/validations"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	middlewares.LogMiddlewares(e)
	validations.CustomValidation(e)

	eAuth := e.Group("")
	eAuth.Use(middleware.JWT([]byte(constants.JWT_SECRET)))

	/** FEATURE USERS */
	e.POST("/api/v1/register", controllers.RegisterControllers)
	e.POST("/api/v1/login", controllers.LoginControllers)
	eAuth.GET("/api/v1/users", controllers.GetUserControllers)
	eAuth.GET("/api/v1/users/:userId", controllers.DetailUserControllers)
	eAuth.PUT("/api/v1/users", controllers.EditUserControllers)
	eAuth.DELETE("/api/v1/users", controllers.DeleteUserControllers)

	return e
}

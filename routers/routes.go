package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/VJ-Vijay77/echo-dock-k8s/controllers"
)

func Routers(e *echo.Echo) {
	e.GET("/",controllers.Home)
	e.POST("/users",controllers.AddUser)
	e.GET("/users/:id",controllers.GetUser)
	e.PUT("/users/:id",controllers.UpdateUser)
	e.DELETE("/users/:id",controllers.DeleteUser)
}
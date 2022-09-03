package main

import (
	"os"

	"github.com/VJ-Vijay77/echo-dock-k8s/conf"
	"github.com/VJ-Vijay77/echo-dock-k8s/controllers"
	"github.com/VJ-Vijay77/echo-dock-k8s/routers"
	"github.com/labstack/echo/v4"
)

func main() {
	conf.LoadEnv()
	Port := os.Getenv("PORT")
	e := echo.New()
	routers.Routers(e)
	controllers.SchemaInitialise()
	e.Logger.Fatal(e.Start(":"+Port))
}
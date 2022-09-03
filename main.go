package main

import (
	"log"

	"github.com/VJ-Vijay77/echo-dock-k8s/controllers"
	"github.com/VJ-Vijay77/echo-dock-k8s/routers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("no env")
	}
	Port := "8080"
	e := echo.New()
	routers.Routers(e)
	controllers.SchemaInitialise()
	e.Logger.Fatal(e.Start(":" + Port))
}

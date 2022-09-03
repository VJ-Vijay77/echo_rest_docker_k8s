package controllers

import "github.com/labstack/echo/v4"




func Home(c echo.Context) error {

	return c.JSON(200,"Welcome to the Home Page...")

}

func AddUser(c echo.Context) error {
	return nil
}

func GetUser(c echo.Context) error {
	return nil
}

func UpdateUser(c echo.Context) error {
	return nil
}

func DeleteUser(c echo.Context) error {
	return nil
}
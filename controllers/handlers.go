package controllers

import (
	"log"

	"github.com/VJ-Vijay77/echo-dock-k8s/database"
	"github.com/VJ-Vijay77/echo-dock-k8s/models"
	"github.com/labstack/echo/v4"
)



func Home(c echo.Context) error {

	return c.JSON(200, "Welcome to the Home Page...")

}

func AddUser(c echo.Context) error {
	db := database.ConnectDB()
	var users models.Users
	err := c.Bind(&users)
	if err != nil {
		return c.JSON(400, "Couldnt get data")
	}

	user1 := `INSERT INTO person(firstname,lastname,email,password,age) VALUES($1,$2,$3,$4,$5) `
	result, err := db.Exec(user1, users.FirstName, users.LastName, users.Email, users.Password, users.Age)
	if err != nil {
		log.Println("database operation declined")
	}
	res, _ := result.RowsAffected()
	last,_ := result.LastInsertId()

	return c.JSON(200, map[string]interface{}{
		"rows affected":res,
		"last":last,
	})
}

func GetUser(c echo.Context) error {
	db := database.ConnectDB()

	var usr models.Users
	name := c.Param("id")
	err := db.Get(&usr, "SELECT * FROM person WHERE firstname=$1", name)
	if err != nil {
		return c.String(400, "Couldnt find data from database")
	}

	return c.JSON(200, map[string]interface{}{
		"name":     usr.FirstName,
		"lastname": usr.LastName,
		"email":    usr.Email,
		"age":      usr.Age,
		"Password": "Encrypted",
	})
}

func UpdateUser(c echo.Context) error {
	db := database.ConnectDB()

	var usr models.Users

	name := c.Param("id")

	err := c.Bind(&usr)
	if err != nil {
		return c.JSON(400, "Couldnt get data")
	}

	query := `UPDATE person SET lastname=$1,email=$2,age=$3 WHERE firstname=$4`
	result, err := db.Exec(query, usr.LastName, usr.Email, usr.Age, name)
	if err != nil {
		log.Println("database operation declined")
		return c.JSON(400, "Declined")
	}
	res, _ := result.RowsAffected()
	var data models.Users
	err = db.Get(&data, "SELECT * FROM person WHERE firstname=$1", name)
	if err != nil {
		log.Println("couldnt get data")
	}
	return c.JSON(200, map[string]interface{}{
		"Rows Affected": res,
		"FirstName":     data.FirstName,
		"LastName":      data.LastName,
		"Email":         data.Email,
		"Age":           data.Age,
	})
}

func DeleteUser(c echo.Context) error {
	db := database.ConnectDB()
	name := c.Param("id")

	res,err := db.Exec("DELETE FROM person WHERE firstname=$1",name)
	if err != nil{
		return c.String(400,"Declined")
	}
	count,_ := res.RowsAffected()
	last,_ := res.LastInsertId()


	return c.JSON(200,map[string]interface{}{
		"Rows affected":count,
		"last":last,
		"Status":"Successful",
	})
}

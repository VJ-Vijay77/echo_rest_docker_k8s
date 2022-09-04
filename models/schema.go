package models


var UserTable = `CREATE TABLE person(
	firstname text,
	lastname text,
	email text,
	password text,
	age integer);`


type Users struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
	Password string `json:"password"`
	Age int64 `json:"age"`
}
package controllers

import (
	"github.com/VJ-Vijay77/echo-dock-k8s/database" 
	"github.com/VJ-Vijay77/echo-dock-k8s/models" 

)
func SchemaInitialise() {
	db := database.ConnectDB()
	db.MustExec(models.UserTable) 
	
}
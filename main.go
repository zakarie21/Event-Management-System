package main

import (
	"RESTApi/db"
	"RESTApi/routes"

	"github.com/gin-gonic/gin"
) 




func main() {

	db.InitDB()
	server := gin.Default()
	routes.RouterInitialisation(server)
	server.Run("localhost:8080")
}




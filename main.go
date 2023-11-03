package main

import (
	"auth-api/dbutils"
	"auth-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	dbutils.ConnectToDatabase()

	engine := gin.Default()

	engine.POST("/new", handlers.NewUser)
	engine.POST("/change", handlers.UpdatePassword)
	engine.POST("/delete", handlers.DeleteMember)

	engine.Run("localhost:9090")
}

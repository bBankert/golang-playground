package main

import (
	"example.com/config"
	"example.com/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitializeDatabase()

	defer config.CloseDatabaseConnection(db)

	httpServer := gin.Default()

	routes.RegisterRoutes(httpServer, db)

	httpServer.Run(":8080")
}

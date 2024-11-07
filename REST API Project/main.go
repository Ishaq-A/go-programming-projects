package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create Server Engine (Engine Instance)
	server := gin.Default()
	
	// Initialise Database
	db.InitDB()
	
	// Register Routes
	routes.RegisterRoutes(server)
	
	// Start Server
	server.Run(":8080") // localhost:8080
}


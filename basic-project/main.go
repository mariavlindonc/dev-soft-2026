package main

import (
	"basic-project/db"

	"basic-project/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.StartDB()
	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default() // I let my project know that I want to use the Gin framework by importing it at the top of the file.

	router.GET("/albums", controllers.GetAll)
	router.POST("/album", controllers.PostAlbum)
	router.GET("/albums/:id", controllers.GetAlbumByID)
	// router.PUT("/albums/:id", editAlbumByID)
	// router.DELETE("/albums/:id", deleteAlbumByID)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows) if nothing is written in the Run() method.
	router.Run("localhost:8080") // This starts the server and listens for incoming requests on the specified address and port.
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Year   int     `json:"year"`
	Price  float64 `json:"price"`
}

// The struct tags (e.g., `json:"id"`) specify how the fields should be encoded/decoded when converting to/from JSON.
// It specifies the format of the JSON key that corresponds to the struct field when the struct is serialized to JSON or deserialized from JSON.

var albums = []album{
	{ID: "1", Title: "The Number of the Beast", Artist: "Iron Maiden", Year: 1982, Price: 25.19},
	{ID: "2", Title: "Youthanasia", Artist: "Medadeth", Year: 1994, Price: 13.65},
	{ID: "3", Title: "Master of Puppets", Artist: "Metallica", Year: 1986, Price: 20.97},
}

// GET ALL
func getAll(c *gin.Context) {
	c.JSON(http.StatusOK, albums) // This line sends a JSON response with the list of albums and an HTTP status code of 200 (OK).
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id") // This line retrieves the value of the "id" parameter from the URL path and stores it in the variable id.
	for index := range albums {
		if albums[index].ID == id {
			c.IndentedJSON(http.StatusOK, albums[index]) // If a matching album is found, this line sends a JSON response with the album details and an HTTP status code of 200 (OK).
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "album not found"}) // If no matching album is found after iterating through the list, this line sends a JSON response with an error message and an HTTP status code of 404 (Not Found).
}

func postAlbum(c *gin.Context) {
	var newAlbum album
	err := c.BindJSON(&newAlbum)
	// This line attempts to bind the incoming JSON payload from the request body to the newAlbum variable.
	// If the JSON is not properly formatted or does not match the structure of the album struct, an error will occur.
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // If there is an error during JSON binding, this line sends a JSON response with an error message and an HTTP status code of 400 (Bad Request).
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func pingPong(c *gin.Context) { // When this URL is accessed, the function will be executed. The 'c' parameter is the context of the request, which helps us avoid having to deal with the raw HTTP request and response objects directly.
	// Return JSON response
	c.JSON(http.StatusOK, gin.H{
		"message": "pong", // This is the response that will be sent back to the client.
	})
}

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default() // I let my project know that I want to use the Gin framework by importing it at the top of the file.

	// Define a simple GET endpoint
	router.GET("/ping", pingPong)
	router.GET("/albums", getAll)
	router.POST("/album", postAlbum)
	router.GET("/albums/:id", getAlbumByID)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows) if nothing is written in the Run() method.
	router.Run("localhost:8080") // This starts the server and listens for incoming requests on the specified address and port.
}

package controllers

import (
	"basic-project/db"
	"basic-project/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET ALL
func GetAll(c *gin.Context) {
	var albums []dtos.Album

	db.DB.Find(&albums)

	c.JSON(http.StatusOK, albums) // This line sends a JSON response with the list of albums and an HTTP status code of 200 (OK).
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id") // This line retrieves the value of the "id" parameter from the URL path and stores it in the variable id.
	var album []dtos.Album

	result := db.DB.First(&album, "id = ?", id)

	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, "album not found")
	}

	c.JSON(http.StatusOK, album)
}

func PostAlbum(c *gin.Context) {
	var newAlbum dtos.Album
	err := c.BindJSON(&newAlbum)
	// This line attempts to bind the incoming JSON payload from the request body to the newAlbum variable.
	// If the JSON is not properly formatted or does not match the structure of the album struct, an error will occur.
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // If there is an error during JSON binding, this line sends a JSON response with an error message and an HTTP status code of 400 (Bad Request).
		return
	}

	db.DB.Create(newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// func editAlbumByID(c *gin.Context) {
// 	id := c.Param("id") // This line retrieves the value of the "id" parameter from the URL path and stores it in the variable id.
// 	var updatedAlbum album
// 	err := c.BindJSON(&updatedAlbum)
// 	if err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // If there is an error during JSON binding, this line sends a JSON response with an error message and an HTTP status code of 400 (Bad Request).
// 		return
// 	}
// 	for index := range albums {
// 		if albums[index].ID == id {
// 			albums[index] = updatedAlbum
// 			c.IndentedJSON(http.StatusOK, updatedAlbum)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "album not found"}) // If no matching album is found after iterating through the list, this line sends a JSON response with an error message and an HTTP status code of 404 (Not Found).
// }

// func deleteAlbumByID(c *gin.Context) {
// 	id := c.Param("id") // This line retrieves the value of the "id" parameter from the URL path and stores it in the variable id.
// 	for index := range albums {
// 		if albums[index].ID == id {
// 			albums = append(albums[:index], albums[index+1:]...)
// 			c.IndentedJSON(http.StatusOK, "album deleted.")
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "album not found"}) // If no matching album is found after iterating through the list, this line sends a JSON response with an error message and an HTTP status code of 404 (Not Found).
// }

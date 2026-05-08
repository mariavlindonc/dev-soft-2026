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
	var album dtos.Album

	result := db.DB.First(&album, "id = ?", id)

	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "album not found"})
		return
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

	db.DB.Create(&newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func EditAlbumByID(c *gin.Context) {
	id := c.Param("id")

	var album dtos.Album
	if err := db.DB.First(&album, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "album not found"})
		return
	}

	var updatedAlbum dtos.Album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Model(&album).Updates(updatedAlbum).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, album)
}

func DeleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	var album dtos.Album

	result := db.DB.Delete(&album, "id = ?", id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "album not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "album deleted"})
}

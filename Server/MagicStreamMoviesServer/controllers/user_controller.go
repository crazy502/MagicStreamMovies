package controllers

import (
	"net/http"

	"github.com/crazy502/MagicStreamMovies/Server/MagicStreamMoviesServer/models"
	"github.com/gin-gonic/gin"
)

func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
			return
		}
	}
}

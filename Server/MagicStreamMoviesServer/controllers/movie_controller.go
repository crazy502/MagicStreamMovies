package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/crazy502/MagicStreamMovies/Server/MagicStreamMoviesServer/database"
	"github.com/crazy502/MagicStreamMovies/Server/MagicStreamMoviesServer/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/gin-gonic/gin"
)

var movieCollection *mongo.Collection = database.OpenCollection("movies")

func GetMovies() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Check if movieCollection is nil before using it
		if movieCollection == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database collection is not initialized"})
			return
		}

		var movies []models.Movie

		cursor, err := movieCollection.Find(ctx, bson.M{})

		if err != nil {
			log.Printf("Failed to fetch movies: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies: " + err.Error()})
			return
		}
		defer cursor.Close(ctx)

		if err = cursor.All(ctx, &movies); err != nil {
			log.Printf("Failed to decode movies: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode movies: " + err.Error()})
			return
		}

		log.Printf("Successfully fetched %d movies", len(movies))
		c.JSON(http.StatusOK, movies)
	}
}
package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/crazy502/MagicStreamMovies/Server/MagicStreamMoviesServer/database"
	"github.com/crazy502/MagicStreamMovies/Server/MagicStreamMoviesServer/models"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/gin-gonic/gin"
)

var movieCollection *mongo.Collection = database.OpenCollection("movies")

var validate = validator.New()

// 检索整个电影集合的数据
func GetMovies() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if movieCollection == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database collection is not initialized"})
			return
		}

		var movies []models.Movie

		cursor, err := movieCollection.Find(ctx, bson.M{})

		if err != nil {
			// log.Printf("Failed to fetch movies: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies: " + err.Error()})
			// return
		}
		defer cursor.Close(ctx)

		if err = cursor.All(ctx, &movies); err != nil {
			// log.Printf("Failed to decode movies: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode movies: " + err.Error()})
			// return
		}

		// log.Printf("Successfully fetched %d movies", len(movies))
		c.JSON(http.StatusOK, movies)
	}
}

// 检索单个电影
// 查询将基于相关电影的IMDb
// gin.HandlerFunc连接到Gin框架
func GetMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		//取消此处的超时，目的是为了清理资源
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		//延迟取消执行，直到外围函数返回，确保了上下文，即使函数由于错误提前退出，也能得到正确清理并释放资源
		defer cancel()

		movieID := c.Param("imdb_id")

		if movieID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Movie ID is required"})
			return
		}

		var movie models.Movie

		err := movieCollection.FindOne(ctx, bson.M{"imdb_id": movieID}).Decode(&movie)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
			return
		}

		c.JSON(http.StatusOK, movie)
	}
}

func AddMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var movie models.Movie

		if err := c.ShouldBindJSON(&movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		if err := validate.Struct(movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
			return
		}

		result, err := movieCollection.InsertOne(ctx, movie)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add movie"})
			return
		}

		c.JSON(http.StatusCreated, result)
	}
}

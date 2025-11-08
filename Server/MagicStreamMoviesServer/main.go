package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	// "github.com/crazy502/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
	controller "github.com/crazy502/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
)

func main() {

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello,MagicStreamMovies!")
	})

	router.GET("/movies", controller.GetMovies())
	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())
	router.POST("/register", controller.RegisterUser())
	router.POST("/login", controller.LoginUser())

	if err := router.Run(":8081"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}

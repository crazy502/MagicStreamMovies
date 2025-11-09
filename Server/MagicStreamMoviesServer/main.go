package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/crazy502/MagicStreamMovies/Server/MagicStreamMoviesServer/routes"
)

func main() {

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello,MagicStreamMovies!")
	})

	routes.SetupUnProtectedRoutes(router)
	routes.SetupProtectedRoutes(router)

	if err := router.Run(":8081"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}

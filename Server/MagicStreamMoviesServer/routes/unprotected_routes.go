package routes

import (
	"github.com/gin-gonic/gin"

	controller "github.com/crazy502/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
)

func SetupUnProtectedRoutes(router *gin.Engine) {

	router.GET("/movies", controller.GetMovies())
	router.POST("/register", controller.RegisterUser())
	router.POST("/login", controller.LoginUser())

}

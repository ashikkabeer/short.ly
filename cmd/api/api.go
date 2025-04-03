package api

import "github.com/gin-gonic/gin"

func SetupAPI() {
	router := gin.Default()

	router.POST("/short")
	router.GET("/")
	router.GET("/:slug") // redirect url
	router.GET("/analytics")
}
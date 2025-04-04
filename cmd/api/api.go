package api

import (
	"github.com/ashikkabeer/short.ly/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupAPI() *gin.Engine {
	router := gin.Default()
	router.POST("/short", handler.GenerateShortURL)
	router.GET("/:short_url", handler.RetrieveOriginalUrl)
	// 	router.GET("/analytics")
	return router
}

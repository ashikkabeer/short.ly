package handler

import (
	"log"

	"github.com/ashikkabeer/short.ly/short.ly/internal/service"
	"github.com/gin-gonic/gin"
)

type MessageBody struct {
	url string
}

func GenerateShortURL(c *gin.Context) {
	var requestBody MessageBody
	url := c.ShouldBindJSON(requestBody)
	log.Println(url)
	// got the url

	url, err := service.GenerateShortURL(requestBody.url)
	// pass it to the service that creates a short url

}
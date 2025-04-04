package handler

import (
	"net/http"
	"regexp"

	"github.com/ashikkabeer/short.ly/config/db"
	"github.com/ashikkabeer/short.ly/internal/repository"
	"github.com/ashikkabeer/short.ly/internal/service"
	"github.com/gin-gonic/gin"
)

type RequestBody struct {
    Url string `json:"url"`
}

var urlService *service.URLService

// InitHandler initializes the handler package by setting up the URLService with a repository instance
func InitHandler() error {
	repo, err := repository.NewPGURLRepository(db.DB)
	if err != nil {
		return err
	}
	urlService = service.NewURLService(repo)
	return nil
}


func GenerateShortURL(c *gin.Context) {
    var req RequestBody
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    if req.Url == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Missing URL"})
        return
    }
    // Validate URL format using regex
    urlPattern := `^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-zA-Z0-9]+([\-\.]{1}[a-zA-Z0-9]+)*\.[a-zA-Z]{2,5}(:[0-9]{1,5})?(\/.*)?$`
    match, err := regexp.MatchString(urlPattern, req.Url)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "URL validation failed"})
        return
    }
    if !match {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL format"})
        return
    }

    // Call the GenerateShortURL method
    ip := c.ClientIP() // Get the client's IP address
    shortURL, err := urlService.GenerateShortURL(req.Url, ip)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"shortUrl": shortURL})
}

func RetrieveOriginalUrl(c *gin.Context) {
	shortCode := c.Param("short_url")
	if shortCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing URL"})
		return
	}

	// Call the service layer to get the original url
	url, err := urlService.RetrieveOriginalUrl(shortCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"originalURL": url})
	return
}
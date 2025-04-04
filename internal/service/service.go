package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/ashikkabeer/short.ly/internal/cache"
	"github.com/ashikkabeer/short.ly/internal/repository"
)

const cacheThreshold = 10 // Define the threshold for hot URLs

type URLService struct {
	repo repository.URLRepository
}
var ctx = context.Background()

func NewURLService(repo repository.URLRepository) *URLService {
	return &URLService{repo: repo}
}

func (s *URLService) GenerateShortURL(url string, ip string) (string, error) {
	hash := GenerateHASH(url)
	err := s.repo.Save(ctx, hash, url, ip)
	if err != nil {
		return "", fmt.Errorf("failed to save URL: %v", err)
	}
	return hash, nil
}

func (s *URLService) RetrieveOriginalUrl(shortCode string) (string, error) {
	// Check if the URL is cached
	originalUrl, _ := cache.Get(shortCode)
	if originalUrl != "" {
		return originalUrl, nil
	}

	// Retrieve the URL from the database
	originalUrl, err := s.repo.Find(ctx, shortCode)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve URL: %v", err)
	}

	// Check if the URL is hot
	accessCount, _ := cache.GetAccessCount(shortCode)
	if accessCount >= cacheThreshold {
		cache.Create(shortCode, originalUrl) // Cache for 24 hours
	}

	return originalUrl, nil
}

func GenerateHASH(url string) string {
	// Combine URL and timestamp with a random number for more uniqueness
	timestamp := time.Now().UnixNano()
	randomPart := rand.Int63()
	input := fmt.Sprintf("%s%d%d", url, timestamp, randomPart)
	
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hash := hasher.Sum(nil)
	hashString := hex.EncodeToString(hash)
	return hashString[:8]
}
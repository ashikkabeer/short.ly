package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/ashikkabeer/short.ly/internal/repository"
)

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
	ctx := context.Background()
	originalUrl, err := s.repo.Find(ctx, shortCode)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve URL: %v", err)
	}
	return originalUrl, nil
}

func GenerateHASH(url string) string {
	// Use time as part of the input to ensure different hashes
	timestamp := time.Now().UnixNano()
	// Combine URL and timestamp with a random number for more uniqueness
	randomPart := rand.Int63()
	input := fmt.Sprintf("%s%d%d", url, timestamp, randomPart)
	
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hash := hasher.Sum(nil)
	hashString := hex.EncodeToString(hash)
	return hashString[:8]
}
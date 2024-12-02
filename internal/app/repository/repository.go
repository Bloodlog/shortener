package repository

import (
	"errors"
	"sync"
)

type URLRepository struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewURLRepository() *URLRepository {
	return &URLRepository{
		data: make(map[string]string),
	}
}

func (repo *URLRepository) Set(shortID string, destURL string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.data[shortID] = destURL
}

func (repo *URLRepository) Get(shortID string) (string, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	destURL, exists := repo.data[shortID]
	if !exists {
		return "", errors.New("shortID not found")
	}
	return destURL, nil
}

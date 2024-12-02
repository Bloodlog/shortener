package service

import (
	"errors"
	"fmt"
	"net/url"
	"shortener/internal/app/repository"
)

type AddURLRequest struct {
	DestURL string
}

type AddURLResponse struct {
	ShortURL string
}

var (
	ErrGetFoobarInvalidRequest = errors.New("invalid get foobar request")
)

func Add(req AddURLRequest, repository *repository.URLRepository) (*AddURLResponse, error) {
	if !isValidURL(req.DestURL) {
		return nil, fmt.Errorf("%w: url address not valid, got %s", ErrGetFoobarInvalidRequest, req.DestURL)
	}
	shortID := "EwHXdJfB"
	repository.Set(shortID, req.DestURL)

	return &AddURLResponse{
		ShortURL: fmt.Sprintf("http://localhost:8080/%s", shortID),
	}, nil
}

func isValidURL(rawURL string) bool {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return false
	}

	return parsedURL.Scheme == "http" || parsedURL.Scheme == "https" && parsedURL.Host != ""
}

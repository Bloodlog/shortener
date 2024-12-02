package service

import (
	"shortener/internal/app/repository"
)

type GetURLRequest struct {
	ID string
}

type GetURLResponse struct {
	DestURL string
}

func Get(req GetURLRequest, repository *repository.URLRepository) (*GetURLResponse, error) {
	destURL, err := repository.Get(req.ID)

	return &GetURLResponse{
		DestURL: destURL,
	}, err
}

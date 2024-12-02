package handlers

import (
	"errors"
	"io"
	"net/http"
	"shortener/internal/app/repository"
	"shortener/internal/app/service"
	"strings"
)

func Handle(repository *repository.URLRepository) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "text/plain")
		if request.Method == http.MethodPost {
			body, err := io.ReadAll(request.Body)
			if err != nil {
				http.Error(response, "Failed to read request body", http.StatusInternalServerError)
				return
			}
			defer request.Body.Close()

			addURLResponse, err := service.Add(service.AddURLRequest{DestUrl: string(body)}, repository)
			if err != nil {
				if errors.Is(err, service.ErrGetFoobarInvalidRequest) {
					http.Error(response, err.Error(), http.StatusBadRequest)
					return
				}
				http.Error(response, "Failed to read request body", http.StatusInternalServerError)
				return
			}

			response.WriteHeader(http.StatusCreated)
			_, _ = response.Write([]byte(addURLResponse.ShortURL))

			return
		}

		if request.Method == http.MethodGet {
			parts := strings.Split(request.RequestURI, "/")
			if len(parts) < 2 {
				response.WriteHeader(http.StatusNotFound)
				return
			}

			responseService, err := service.Get(service.GetURLRequest{Id: parts[1]}, repository)
			if err != nil {
				response.WriteHeader(http.StatusNotFound)
				return
			}
			http.Redirect(response, request, responseService.DestURL, http.StatusTemporaryRedirect)
			return
		}
		response.WriteHeader(http.StatusBadRequest)
	}
}

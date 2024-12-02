package main

import (
	"log"
	"shortener/internal/app/repository"
	"shortener/internal/app/router"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	rep := repository.NewURLRepository()

	return router.Run(rep)
}

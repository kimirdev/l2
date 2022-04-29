package main

import (
	"dev11/internal/repository"
	"dev11/internal/server/httpserver"
	"dev11/internal/service"
	"log"
)

func main() {

	db := repository.NewCalendarCache()

	service := service.NewService(db)

	server := httpserver.New(service)

	if err := server.Run(); err != nil {
		log.Println(err)
	}
}

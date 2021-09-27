package main

import (
	"log"

	"github.com/rhavyx/react_go_watch-movies/backend_golang/src/models"
)

type configuration struct {
	port    int
	env     string
	version string
	db      struct {
		dsn string
	}
}

type application struct {
	config configuration
	logger *log.Logger
	models models.Models
}

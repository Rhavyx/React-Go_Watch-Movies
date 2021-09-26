package main

import (
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"
)

const version = "1.0.0"

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type config_type struct {
	port int
	env  string
}

func Status(c *gin.Context) {

	var config config_type

	flag.IntVar(&config.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&config.env, "env", "development", "Application environment (development|production)")
	flag.Parse()

	currentStatus := AppStatus{
		Status:      "Available",
		Environment: config.env,
		Version:     version,
	}

	c.JSON(http.StatusOK, currentStatus)

}

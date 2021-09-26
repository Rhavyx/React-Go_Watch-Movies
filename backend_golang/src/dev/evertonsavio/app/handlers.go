package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const version = "1.0.0"

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func (app *application) Status(c *gin.Context) {

	currentStatus := AppStatus{
		Status:      "Available",
		Environment: app.config.env,
		Version:     version,
	}

	c.JSON(http.StatusOK, currentStatus)

}

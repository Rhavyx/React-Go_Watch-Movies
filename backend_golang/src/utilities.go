package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) errorJson(c *gin.Context, err error) {

	type jsonError struct {
		Message string `json:"message"`
	}
	theError := jsonError{
		Message: err.Error(),
	}

	c.JSON(http.StatusBadRequest, theError)

}

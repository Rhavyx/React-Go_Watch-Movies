package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	log.Println("init...")
	router = gin.Default()
}

func (app *application) StartApp() {

	router.GET("/status", app.Status)

	if err := router.Run(":4000"); err != nil {
		panic(err)
	}

	log.Panicln("Running")

}

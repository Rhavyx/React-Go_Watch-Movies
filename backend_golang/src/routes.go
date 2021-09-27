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

func (app *application) startApp() {

	router.GET("/status", app.statusHandler)
	router.GET("/movie/:movie_id", app.getOneMovie)
	router.GET("/movies", app.getAllMovies)

	if err := router.Run(":4000"); err != nil {
		panic(err)
	}

	log.Panicln("Running")

}

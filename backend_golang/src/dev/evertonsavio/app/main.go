package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const version = "1.0.0"

type config_type struct {
	port int
	env  string
}

func main() {
	var config config_type

	flag.IntVar(&config.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&config.env, "env", "development", "Application environment (development|production)")
	flag.Parse()

	fmt.Println("running")

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "{\"heath\": true }")
	})
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.port), nil)
	if err != nil {
		log.Println(err)
	}

}

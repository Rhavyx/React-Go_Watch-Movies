package main

import (
	"flag"
	"log"
	"os"
)

type configuration struct {
	port    int
	env     string
	version string
}

type application struct {
	config configuration
	logger *log.Logger
}

func main() {

	var cfg configuration
	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production)")
	flag.StringVar(&cfg.version, "version", "1.0.0", "application version")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	app.StartApp()
}

package main

import "flag"

type config_type struct {
	port int
	env  string
}

var config config_type

func main() {

	flag.IntVar(&config.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&config.env, "env", "development", "Application environment (development|production)")
	flag.Parse()

	StartApp()
}

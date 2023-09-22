package main

import (
	"flag"
	"log"
	"shrink-url/internal/config"
	"shrink-url/internal/server"
)

const (
	defaultPort = "80"
)

func main() {
	var port string
	flag.StringVar(&port, "port", defaultPort, "server listening port")
	flag.Parse()

	config.LoadConfig()

	router := server.NewRouter()
	server := server.NewServer(port, router)

	log.Printf("Starting server on port %s ...\n", port)
	if err := server.Start(); err != nil {
		log.Fatal("failed to start server")
	}
	defer server.Close()
}

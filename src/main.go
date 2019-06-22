package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func getPort() string {
	port := os.Getenv("SERVER_PORT")

	if port == "" {
		port = "8080"
	}

	return port
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	port := getPort()

	log.Printf("Starting server on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

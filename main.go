package main

import (
	"fmt"
	"net/http"
	"os"
	"urlShortner/executors"
	"urlShortner/storage"
)

func main() {
	store := storage.NewStore()
	handler := executors.NewHandler(store)
	http.HandleFunc("/shorten", handler.ShortenURL)
	http.HandleFunc("/", handler.RedirectURL)
	port := os.Getenv("port")
	if port == "" {
		port = ":8080"
	}
	fmt.Println("SERVER IS RUNNING....")
	if err := http.ListenAndServe(port, nil); err != nil {
		println("Error starting server:", err)
	}
}

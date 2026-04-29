package main

import (
	"fmt"
	"net/http"
	"os"
	"urlShortner/db"
	"urlShortner/executors"
	auth_login "urlShortner/executors/Auth"
	"urlShortner/middleware"
	"urlShortner/storage"
)

func main() {
	store := storage.NewStore()
	handler := executors.NewHandler(store)
	db.CreateTables()
	fmt.Print("CREATE END :")
	http.HandleFunc("/login", auth_login.LoginHandler)
	http.Handle("/shorten", middleware.AuthMiddleware(http.HandlerFunc(handler.ShortenURL)))
	http.Handle("/", middleware.AuthMiddleware(http.HandlerFunc(handler.RedirectURL)))
	port := os.Getenv("port")
	if port == "" {
		port = ":8080"
	}
	fmt.Println("SERVER IS RUNNING....")
	if err := http.ListenAndServe(port, nil); err != nil {
		println("Error starting server:", err)
	}
}

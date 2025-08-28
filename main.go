package main

import (
	"fmt"
	"net/http"
)

func main() {

	port := "8080"

	router := http.NewServeMux()

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	router.HandleFunc("GET /", handleHome)

	server.ListenAndServe()
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome home"))
}

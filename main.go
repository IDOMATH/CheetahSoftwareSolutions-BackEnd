package main

import (
	"encoding/json"
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

	fmt.Println("running")
	server.ListenAndServe()
}

type msg struct {
	Message string `json:"message"`
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	m, _ := json.Marshal(msg{Message: "Welcome home!"})
	w.Write(m)
}

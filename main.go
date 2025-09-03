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

type jaysahn struct {
	message string
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	m, _ := json.Marshal(jaysahn{message: "Welcome home"})
	w.Write(m)
}

package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	itemHandler := &ItemHandler{}

	router.HandleFunc("POST /item", itemHandler.Create)
	router.HandleFunc("GET /item/{id}", itemHandler.GetById)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Starting server on port :8080")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

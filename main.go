package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	itemHandler := &ItemHandler{}

	itemRouterV1 := http.NewServeMux()
	itemRouterV1.HandleFunc("GET /item/{id}", itemHandler.GetById)

	itemRouterV2 := http.NewServeMux()
	itemRouterV2.HandleFunc("GET /item/{id}", itemHandler.GetById)
	itemRouterV2.HandleFunc("POST /item", itemHandler.Create)

	router.Handle("/v1/", http.StripPrefix("/v1", itemRouterV1))
	router.Handle("/v2/", http.StripPrefix("/v2", itemRouterV2))

	stack := middlewareStack(
		LogMiddleware,
		AuthMiddleware,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	log.Println("Starting server on port :8080")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

type middleware func(http.Handler) http.Handler

func middlewareStack(ms ...middleware) middleware {
	return func(next http.Handler) http.Handler {
		for i := len(ms) - 1; i >= 0; i-- {
			next = ms[i](next)
		}
		return next
	}
}

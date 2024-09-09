package main

import (
	"log"
	"net/http"
)

type Middlewares func(http.HandlerFunc) http.HandlerFunc

var middlewares = []Middlewares{
	TokenAuthMiddleware,
}

func main() {

	var handler http.HandlerFunc = handleClientProfile

	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	http.HandleFunc("/user/profile", handler)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/Taha-flutter/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	host, err := os.Hostname()
	if err != nil {
		log.Fatalf("Error getting hostname: %v", err)
	}
	port := "8080"
	fmt.Printf("Server is running on http://%s:%s\n", host, port)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

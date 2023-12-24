package main

import (
	"github.com/ajay0221/go-bookstore/pkg/config"
	"github.com/ajay0221/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	config.Connect()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:9010", router))
}

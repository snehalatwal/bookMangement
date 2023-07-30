package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/snehalatwal/bookMangement/pkg/routes"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}

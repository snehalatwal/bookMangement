package routes

import (
	"github.com/gorilla/mux"
	"github.com/snehalatwal/bookMangement/pkg/controller"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {

	router.HandleFunc("/book/", controller.CreateBooks).Methods("POST")
	router.HandleFunc("/book/", controller.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controller.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controller.UpdateBookById).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controller.DeleteBookById).Methods("DELETE")
}

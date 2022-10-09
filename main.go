package main

import (
	"FoodApi/src/handlers"
	"FoodApi/src/models"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type allProducts []models.Product

func handleRequest() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/product", handlers.CreateProduct).Methods(http.MethodPost)
	router.HandleFunc("/product", handlers.GetAllProducts).Methods(http.MethodGet)
	router.HandleFunc("/product/{id}", handlers.GetOneProduct).Methods(http.MethodGet)
	router.HandleFunc("/product/{id}", handlers.UpdateProduct).Methods(http.MethodPatch)
	router.HandleFunc("/product/{id}", handlers.DeleteProduct).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
	return router
}

func main() {
	handleRequest()
	log.Println("API is running!")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Food API")
}

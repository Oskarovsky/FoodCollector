package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome DJ!")
}

type product struct {
	ID       string  `json:"ID"`
	Name     string  `json:"Name"`
	Energy   int32   `json:"Energy"`
	Protein  float32 `json:"Protein"`
	Fat      float32 `json:"Fat"`
	Carbs    float32 `json:"Carbs"`
	Price    float32 `json:"Price"`
	Provider string  `json:"Provider"`
}

type allProducts []product

var products = allProducts{
	{
		ID:       "1",
		Name:     "Apple",
		Energy:   310,
		Protein:  2.10,
		Fat:      0.01,
		Carbs:    1.99,
		Price:    0.99,
		Provider: "Tesco",
	},
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct product
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the product attributes in order to update")
	}

	json.Unmarshal(reqBody, &newProduct)
	products = append(products, newProduct)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newProduct)

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/product", createProduct)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

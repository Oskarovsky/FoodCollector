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

func getOneProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	for _, singleProduct := range products {
		if singleProduct.ID == productId {
			json.NewEncoder(w).Encode(singleProduct)
		}
	}
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	var updatedProduct product

	reqBody, error := io.ReadAll(r.Body)
	if error != nil {
		fmt.Fprintf(w, "Kindly enter data with the product attributes in order to update")
	}
	json.Unmarshal(reqBody, &updatedProduct)

	for i, singleProduct := range products {
		if singleProduct.ID == productId {
			singleProduct.Name = updatedProduct.Name
			singleProduct.Energy = updatedProduct.Energy
			singleProduct.Protein = updatedProduct.Protein
			singleProduct.Price = updatedProduct.Price
			singleProduct.Carbs = updatedProduct.Carbs
			singleProduct.Fat = updatedProduct.Fat
			singleProduct.Provider = updatedProduct.Provider
			products = append(products[:i], singleProduct)
			json.NewEncoder(w).Encode(singleProduct)
		}
	}
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]

	for i, singleProduct := range products {
		if singleProduct.ID == productId {
			products = append(products[:i], products[i+1:]...)
			fmt.Fprintf(w, "The product with id %v has been deleted successfully", productId)
		}
	}
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/product", createProduct).Methods("POST")
	router.HandleFunc("/product", getAllProducts).Methods("GET")
	router.HandleFunc("/product/{id}", getOneProduct).Methods("GET")
	router.HandleFunc("/product/{id}", updateProduct).Methods("PATCH")
	router.HandleFunc("/product/{id}", deleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

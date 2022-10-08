package handlers

import (
	"FoodApi/src/mocks"
	"FoodApi/src/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Products)
}

func GetOneProduct(w http.ResponseWriter, r *http.Request) {
	productId, _ := strconv.Atoi(mux.Vars(r)["id"])
	for _, singleProduct := range mocks.Products {
		if singleProduct.ID == productId {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(singleProduct)
			break
		}
	}
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the product attributes in order to update")
		log.Fatalln(err)
	}

	var newProduct models.Product
	json.Unmarshal(reqBody, &newProduct)
	newProduct.ID = rand.Intn(100)
	mocks.Products = append(mocks.Products, newProduct)
	w.WriteHeader(http.StatusCreated)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId, _ := strconv.Atoi(mux.Vars(r)["id"])

	defer r.Body.Close()
	reqBody, error := io.ReadAll(r.Body)

	if error != nil {
		fmt.Fprintf(w, "Kindly enter data with the product attributes in order to update")
	}

	var updatedProduct models.Product
	json.Unmarshal(reqBody, &updatedProduct)

	for i, singleProduct := range mocks.Products {
		if singleProduct.ID == productId {
			singleProduct.Name = updatedProduct.Name
			singleProduct.Energy = updatedProduct.Energy
			singleProduct.Protein = updatedProduct.Protein
			singleProduct.Price = updatedProduct.Price
			singleProduct.Carbs = updatedProduct.Carbs
			singleProduct.Fat = updatedProduct.Fat
			singleProduct.Provider = updatedProduct.Provider

			mocks.Products = append(mocks.Products[:i], singleProduct)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(singleProduct)
			break
		}
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId, _ := strconv.Atoi(mux.Vars(r)["id"])
	for i, singleProduct := range mocks.Products {
		if singleProduct.ID == productId {
			mocks.Products = append(mocks.Products[:i], mocks.Products[i+1:]...)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "The product with id %v has been deleted successfully", productId)
			break
		}
	}
}

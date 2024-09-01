package api

import (
	model "AahaFeltBackend/models"
	"strconv"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func makeHandler(fn func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// Handle GET /products
func (s *ApiServer) handleGetProducts(w http.ResponseWriter, r *http.Request) error {
	products, err := s.store.GetProducts()
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, products)
}

// Handle POST /products
func (s *ApiServer) handlePostProducts(w http.ResponseWriter, r *http.Request) error {
	product := model.Product{}

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		return err
	}

	pro := model.NewProduct(
		product.ID, product.Weight, product.Price, product.MostPopular, product.Bestseller,
		product.Material, product.Stock, product.NewArrival, product.Designer, product.Company,
		product.HotCollection, product.Colors, product.Category, product.Description, product.Reviews,
		product.Stars, product.Name, product.Notes, product.Featured, product.Sale, product.Trending,
		product.Shipping, product.Origin, product.Image, product.Images, product.Exclusive, product.NewInMarket,
	)

	if err := s.store.AddProducts(*pro); err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, pro)
}

func (s *ApiServer) handleGetProductsById(w http.ResponseWriter, r *http.Request) error {

	vars := mux.Vars(r)["id"]

	id, err := strconv.Atoi(vars)
	if err != nil {
		return err // return error if conversion fails
	}
	products, err := s.store.GetProductsById(id)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, products)
}

func (s *ApiServer) UpdateProductHandler(w http.ResponseWriter, r *http.Request) error {

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return nil
	}

	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return nil
	}

	err = s.store.UpdateProductById(id, product)
	if err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return nil
	}

	w.WriteHeader(http.StatusOK)

	return writeJSON(w, http.StatusOK, "Product updated successfully")

}

func (s *ApiServer) handleDeleteProduct(w http.ResponseWriter, r *http.Request) error {

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return nil
	}

	err = s.store.DeleteProductById(id)
	if err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return nil
	}

	w.WriteHeader(http.StatusOK)

	return writeJSON(w, http.StatusOK, "Product deleted successfully")
}

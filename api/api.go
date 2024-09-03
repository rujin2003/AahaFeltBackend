package api

import (
	"AahaFeltBackend/storage"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	address string
	store   storage.Storage
}

func NewApiServer(address string, store storage.Storage) *ApiServer {
	return &ApiServer{
		address: address,
		store:   store,
	}
}

func (s *ApiServer) Start() {
	//MARK: Products
	router := mux.NewRouter()
	router.HandleFunc("/products", makeHandler(s.handleGetProducts)).Methods("GET")
	router.HandleFunc("/products", makeHandler(s.handlePostProducts)).Methods("POST")
	router.HandleFunc("/products/{id}", makeHandler(s.handleGetProductsById)).Methods("GET")
	router.HandleFunc("/products/{id}", makeHandler(s.UpdateProductHandler)).Methods("POST")
	router.HandleFunc("/products/{id}", makeHandler(s.handleDeleteProduct)).Methods("DELETE")

	// MARK: Gallery
	router.HandleFunc("/gallery-images", makeHandler(s.addImageHandler)).Methods("POST")
	router.HandleFunc("/gallery-images/{id}", makeHandler(s.getImageHandler)).Methods("GET")
	router.HandleFunc("/gallery-images", makeHandler(s.getAllImageLinksHandler)).Methods("GET")
	router.HandleFunc("/gallery-images/{id}", makeHandler(s.deleteImageHandler)).Methods("DELETE")

	// MARK: Product Images
	router.HandleFunc("/productimage", makeHandler(s.AddProductImagesHandler)).Methods("POST")
	router.HandleFunc("/productimage/{product_name}", makeHandler(s.GetProductImagesByNameHandler)).Methods("GET")

	router.HandleFunc("/productimg/{id}", makeHandler(s.getImageHandler)).Methods("GET")
	router.HandleFunc("/productimage/{product_name}", makeHandler(s.DeleteProductImagesByNameHandler)).Methods("DELETE")

	fmt.Printf("Server is starting on %s...\n", s.address)
	if err := http.ListenAndServe(s.address, router); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}

	http.ListenAndServe(s.address, router)
}

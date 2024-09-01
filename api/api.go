package api

import (
	"AahaFeltBackend/storage"

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

	router := mux.NewRouter()
	router.Use(s.corsMiddleware)
	router.HandleFunc("/products", makeHandler(s.handleGetProducts)).Methods("GET")
	router.HandleFunc("/products", makeHandler(s.handlePostProducts)).Methods("POST")
	router.HandleFunc("/products/{id}", makeHandler(s.handleGetProductsById)).Methods("GET")
	router.HandleFunc("/products/{id}", makeHandler(s.UpdateProductHandler)).Methods("POST")
	router.HandleFunc("/products/{id}", makeHandler(s.handleDeleteProduct)).Methods("DELETE")
	router.HandleFunc("/images", makeHandler(s.addImageHandler)).Methods("GET")
	router.HandleFunc("/images", makeHandler(s.getImageHandler)).Methods("POST")

	http.ListenAndServe(s.address, router)
}

func (s *ApiServer) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

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
	router.HandleFunc("/products", makeHandler(s.handleProducts)).Methods("GET")
	http.ListenAndServe(s.address, router)
}

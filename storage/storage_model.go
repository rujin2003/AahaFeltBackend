package storage

import (
	model "AahaFeltBackend/models"
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
)

type Storage interface {
	AddProducts(product model.Product) error
	GetProducts() ([]model.Product, error)
	GetProductsById(id int) (*model.Product, error)
	UpdateProductById(id int, product model.Product) error
	DeleteProductById(id int) error
	AddImage(r *http.Request) (int, error)
	// gallery
	GetImageURL(id int) (string, error)
	GetImageByID(id int) (*model.GalleryImage, error)

	Close()
	Init() error
}

type PostgresStorage struct {
	db *sql.DB
}

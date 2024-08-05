package storage

import (
	modle "AahaFeltBackend/models"
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	AddProducts(product modle.Product) error
	GetProducts() ([]modle.Product, error)
	Close()
	Init() error
}

type PostgresStorage struct {
	db *sql.DB
}

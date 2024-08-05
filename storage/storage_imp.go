package storage

import (
	modle "AahaFeltBackend/models"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

func NewPostgresStorage() (*PostgresStorage, error) {
	connStr := "user=postgres password=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Check if the database exists
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = 'bank')").Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf("failed to check if database exists: %w", err)
	}

	if !exists {
		// Create the database if it does not exist
		_, err = db.Exec("CREATE DATABASE bank")
		if err != nil {
			return nil, fmt.Errorf("failed to create database: %w", err)
		}
	}

	// Connect to the newly created or existing database
	db, err = sql.Open("postgres", connStr+" dbname=bank")
	if err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}

// Init initializes the database with the required tables
func (s *PostgresStorage) Init() error {
	_, err := s.db.Exec(`
        CREATE TABLE IF NOT EXISTS products (
            id SERIAL PRIMARY KEY,
            weight TEXT,
            price TEXT,
            most_popular BOOLEAN,
            bestseller BOOLEAN,
            material TEXT,
            stock INT,
            new_arrival BOOLEAN,
            designer TEXT,
            company TEXT,
            hot_collection BOOLEAN,
            colors TEXT[],
            category TEXT,
            description TEXT,
            reviews INT,
            stars FLOAT,
            name TEXT,
            notes TEXT,
            featured BOOLEAN,
            sale BOOLEAN,
            trending BOOLEAN,
            shipping TEXT,
            origin TEXT,
            image TEXT,
            images TEXT[],
            exclusive BOOLEAN,
            new_in_market BOOLEAN
        )
    `)
	return err
}

// Close closes the database connection.
func (s *PostgresStorage) Close() {
	s.db.Close()
}

// AddProducts adds a product to the database
func (s *PostgresStorage) AddProducts(product modle.Product) error {
	_, err := s.db.Exec(`
        INSERT INTO products (weight, price, most_popular, bestseller, material, stock, new_arrival, designer, company, hot_collection, colors, category, description, reviews, stars, name, notes, featured, sale, trending, shipping, origin, image, images, exclusive, new_in_market)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26)`,
		product.Weight, product.Price, product.MostPopular, product.Bestseller, product.Material, product.Stock, product.NewArrival, product.Designer, product.Company, product.HotCollection, product.Colors, product.Category, product.Description, product.Reviews, product.Stars, product.Name, product.Notes, product.Featured, product.Sale, product.Trending, product.Shipping, product.Origin, product.Image, product.Images, product.Exclusive, product.NewInMarket)
	return err
}

// GetProducts retrieves all products from the database
func (s *PostgresStorage) GetProducts() ([]modle.Product, error) {
	rows, err := s.db.Query(`SELECT id, weight, price, most_popular, bestseller, material, stock, new_arrival, designer, company, hot_collection, colors, category, description, reviews, stars, name, notes, featured, sale, trending, shipping, origin, image, images, exclusive, new_in_market FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []modle.Product
	for rows.Next() {
		var product modle.Product
		var colors, images []string
		err := rows.Scan(&product.ID, &product.Weight, &product.Price, &product.MostPopular, &product.Bestseller, &product.Material, &product.Stock, &product.NewArrival, &product.Designer, &product.Company, &product.HotCollection, pq.Array(&colors), &product.Category, &product.Description, &product.Reviews, &product.Stars, &product.Name, &product.Notes, &product.Featured, &product.Sale, &product.Trending, &product.Shipping, &product.Origin, &product.Image, pq.Array(&images), &product.Exclusive, &product.NewInMarket)
		if err != nil {
			return nil, err
		}
		product.Colors = colors
		product.Images = images
		products = append(products, product)
	}
	return products, nil
}

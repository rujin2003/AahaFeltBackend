package storage

import (
	model "AahaFeltBackend/models"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (s *PostgresStorage) InitGallery() error {
	_, err := s.db.Exec(`
        CREATE TABLE IF NOT EXISTS gallery (
            id SERIAL PRIMARY KEY,
            image_base64 TEXT NOT NULL,
            description TEXT
        );
    `)
	return err
}

func (s *PostgresStorage) AddImage(r *http.Request) (int, error) {
	// Parse the form data
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return 0, err
	}

	// Get the file from the request
	file, _, err := r.FormFile("image")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return 0, err
	}

	imageBase64 := base64.StdEncoding.EncodeToString(fileBytes)

	var id int
	description := r.FormValue("description")
	err = s.db.QueryRow(`
        INSERT INTO gallery (image_base64, description)
        VALUES ($1, $2)
        RETURNING id`,
		imageBase64, description).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// GetImageURL returns the URL for an image
func (s *PostgresStorage) GetImageURL(id int) (string, error) {
	return fmt.Sprintf("http://localhost:8080/images/%d", id), nil
}

// GetImageByID retrieves an image by its ID and returns it as base64
func (s *PostgresStorage) GetImageByID(id int) (*model.GalleryImage, error) {
	var image model.GalleryImage
	err := s.db.QueryRow(`SELECT id, image_base64, description FROM gallery WHERE id=$1`, id).
		Scan(&image.ID, &image.ImageBase64, &image.Description)
	if err != nil {
		return nil, err
	}
	return &image, nil
}

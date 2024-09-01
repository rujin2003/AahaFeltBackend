package api

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *ApiServer) addImageHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := s.store.AddImage(r)
	if err != nil {
		http.Error(w, "Failed to add image", http.StatusInternalServerError)
		return nil
	}

	url, err := s.store.GetImageURL(id)
	if err != nil {
		http.Error(w, "Failed to generate image URL", http.StatusInternalServerError)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"image_url": url})
	return writeJSON(w, http.StatusOK, "image added successfully")
}

func (s *ApiServer) getImageHandler(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid image ID", http.StatusBadRequest)
		return nil
	}

	image, err := s.store.GetImageByID(id)
	if err != nil {
		http.Error(w, "Image not found", http.StatusNotFound)
		return nil
	}

	// Decode the base64 image and serve it as an HTTP response
	imageBytes, err := base64.StdEncoding.DecodeString(image.ImageBase64)
	if err != nil {
		http.Error(w, "Failed to decode image", http.StatusInternalServerError)
		return nil
	}

	w.Header().Set("Content-Disposition", "inline; filename=image.jpg")
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(imageBytes)
	return writeJSON(w, http.StatusOK, "image added successfully")
}

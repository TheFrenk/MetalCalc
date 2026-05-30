package handlers

import (
	"MetalCalc/backend/internal/models"
	"MetalCalc/backend/internal/storage"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func GetShapesHandler(w http.ResponseWriter, r *http.Request) {
	list := storage.Global.GetShapes()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func AddShapeHandler(w http.ResponseWriter, r *http.Request) {
	var sh models.CustomShape
	if err := json.NewDecoder(r.Body).Decode(&sh); err != nil {
		http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
		return
	}
	if sh.Name == "" || sh.Key == "" {
		http.Error(w, `{"error":"name and key are required"}`, http.StatusBadRequest)
		return
	}
	created := storage.Global.AddShape(sh)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func DeleteShapeHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := storage.Global.DeleteShape(id); err != nil {
		http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

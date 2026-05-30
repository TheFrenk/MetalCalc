package handlers

import (
	"MetalCalc/backend/internal/models"
	"MetalCalc/backend/internal/storage"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func GetMaterialsHandler(w http.ResponseWriter, r *http.Request) {
	list := storage.Global.GetMaterials()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func AddMaterialHandler(w http.ResponseWriter, r *http.Request) {
	var m models.CustomMaterial
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
		return
	}
	if m.Name == "" || m.Key == "" {
		http.Error(w, `{"error":"name and key are required"}`, http.StatusBadRequest)
		return
	}
	created := storage.Global.AddMaterial(m)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func DeleteMaterialHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := storage.Global.DeleteMaterial(id); err != nil {
		http.Error(w, `{"error":"material not found"}`, http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

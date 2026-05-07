package handlers

import (
	"encoding/json"
	"net/http"

	"MetalCalc/backend/internal/calculations/forging"
)

type ForgeRequest struct {
	Material         string  `json:"material"`
	Shape            string  `json:"shape"`
	DimensionA       float64 `json:"dimension_a"`
	DimensionB       float64 `json:"dimension_b"`
	InitialHeight    float64 `json:"initial_height"`
	FinalHeight      float64 `json:"final_height"`
	Temperature      float64 `json:"temperature"`
	FrictionCoeff    float64 `json:"friction_coeff"`
	DeformationSpeed float64 `json:"deformation_speed"`
}

func ForgingHandler(w http.ResponseWriter, r *http.Request) {
	var req ForgeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
		return
	}

	params := forging.Params{
		Material:         req.Material,
		Shape:            req.Shape,
		DimensionA:       req.DimensionA,
		DimensionB:       req.DimensionB,
		InitialHeight:    req.InitialHeight,
		FinalHeight:      req.FinalHeight,
		Temperature:      req.Temperature,
		FrictionCoeff:    req.FrictionCoeff,
		DeformationSpeed: req.DeformationSpeed,
	}

	result := forging.Calculate(params)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

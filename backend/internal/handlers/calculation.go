package handlers

import (
	"MetalCalc/backend/internal/calculations/forging"
	"MetalCalc/backend/internal/pdf"
	"encoding/json"
	"fmt"
	"net/http"
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

func ForgingPDFHandler(w http.ResponseWriter, r *http.Request) {
	var req ForgeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "invalid json"}`, http.StatusBadRequest)
		return
	}

	params := forging.Params{
		Material: req.Material, Shape: req.Shape, DimensionA: req.DimensionA,
		DimensionB: req.DimensionB, InitialHeight: req.InitialHeight,
		FinalHeight: req.FinalHeight, Temperature: req.Temperature,
		FrictionCoeff: req.FrictionCoeff, DeformationSpeed: req.DeformationSpeed,
	}
	result := forging.Calculate(params)

	reportData := pdf.ReportData{
		Title: "Forging Proccess Calculation Report",
		Params: []pdf.ParamRow{
			{Label: "Material", Value: req.Material},
			{Label: "Shape", Value: req.Shape},
			{Label: "Dimension A", Value: fmt.Sprintf("%.1f", req.DimensionA), Unit: "mm"},
			{Label: "Dimension B", Value: fmt.Sprintf("%.1f", req.DimensionB), Unit: "mm"},
			{Label: "Initial Height", Value: fmt.Sprintf("%.1f", req.InitialHeight), Unit: "mm"},
			{Label: "Final Height", Value: fmt.Sprintf("%.1f", req.FinalHeight), Unit: "mm"},
			{Label: "Temperature", Value: fmt.Sprintf("%.0f", req.Temperature), Unit: "°C"},
			{Label: "Friction Coeff", Value: fmt.Sprintf("%.2f", req.Temperature)},
			{Label: "Deformation Speed", Value: fmt.Sprintf("%.2f", req.DeformationSpeed), Unit: "mm/s"},
		},
		Results: []pdf.ParamRow{
			{Label: "Forging Force", Value: fmt.Sprintf("%.2f", result.ForgingForce), Unit: "N"},
			{Label: "Forging Pressure", Value: fmt.Sprintf("%.2f", result.ForgingPressure), Unit: "MPa"},
			{Label: "Work Done", Value: fmt.Sprintf("%.2f", result.WorkDone), Unit: "J"},
			{Label: "Power", Value: fmt.Sprintf("%.2f", result.Power), Unit: "kW"},
			{Label: "Deformation Speed", Value: fmt.Sprintf("%.4f", result.DeformationSpeed), Unit: "s⁻¹"},
			{Label: "Strain Degree", Value: fmt.Sprintf("%.4f", result.StrainDegree)},
			{Label: "Workpiece Mass", Value: fmt.Sprintf("%.3f", result.WorkpieceMass), Unit: "Kg"},
			{Label: "Contact Area", Value: fmt.Sprintf("%.2f", result.ContactArea), Unit: "mm²"},
			{Label: "Height Reduction", Value: fmt.Sprintf("%.2f", result.HeightReduction), Unit: "mm"},
		},
	}

	pdfBytes, err := pdf.GenerateForgingReport(reportData)
	if err != nil {
		http.Error(w, `{"error": "pdf generation failed"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=\"forging.pdf\"")
	w.Write(pdfBytes)
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

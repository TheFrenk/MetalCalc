package handlers

import (
	"MetalCalc/backend/internal/calculations/forging"
	"MetalCalc/backend/internal/pdf"
	"MetalCalc/backend/internal/storage"
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
	Passes           int     `json:"passes"`
}

func ForgingPDFHandler(w http.ResponseWriter, r *http.Request) {
	var req ForgeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "invalid json"}`, http.StatusBadRequest)
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
		Passes:           req.Passes,
	}
	result := forging.Calculate(params)

	reportData := pdf.ReportData{
		Title: "Звіт розрахунку процесу кування",
		Params: []pdf.ParamRow{
			{Label: "Матеріал", Value: translateMaterial(req.Material)},
			{Label: "Форма", Value: translateShape(req.Shape)},
			{Label: "Розмір A", Value: fmt.Sprintf("%.1f", req.DimensionA), Unit: "мм"},
			{Label: "Розмір B", Value: fmt.Sprintf("%.1f", req.DimensionB), Unit: "мм"},
			{Label: "Початкова висота", Value: fmt.Sprintf("%.1f", req.InitialHeight), Unit: "мм"},
			{Label: "Кінцева висота", Value: fmt.Sprintf("%.1f", req.FinalHeight), Unit: "мм"},
			{Label: "Температура", Value: fmt.Sprintf("%.0f", req.Temperature), Unit: "°C"},
			{Label: "Коеф. тертя", Value: fmt.Sprintf("%.2f", req.FrictionCoeff)},
			{Label: "Швидкість деформації", Value: fmt.Sprintf("%.2f", req.DeformationSpeed), Unit: "мм/с"},
		},
		Results: []pdf.ParamRow{
			{Label: "Зусилля кування", Value: fmt.Sprintf("%.2f", result.ForgingForce), Unit: "Н"},
			{Label: "Тиск кування", Value: fmt.Sprintf("%.2f", result.ForgingPressure), Unit: "МПа"},
			{Label: "Виконана робота", Value: fmt.Sprintf("%.2f", result.WorkDone), Unit: "Дж"},
			{Label: "Потужність", Value: fmt.Sprintf("%.2f", result.Power), Unit: "кВт"},
			{Label: "Швидкість деформ.", Value: fmt.Sprintf("%.4f", result.DeformationSpeed), Unit: "с⁻¹"},
			{Label: "Ступінь деформації", Value: fmt.Sprintf("%.4f", result.StrainDegree)},
			{Label: "Маса заготовки", Value: fmt.Sprintf("%.3f", result.WorkpieceMass), Unit: "кг"},
			{Label: "Площа контакту", Value: fmt.Sprintf("%.2f", result.ContactArea), Unit: "мм²"},
			{Label: "Зменшення висоти", Value: fmt.Sprintf("%.2f", result.HeightReduction), Unit: "мм"},
		},
	}

	pdfBytes, err := pdf.GenerateForgingReport(reportData, "./fonts")
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
		Passes:           req.Passes,
	}

	result := forging.Calculate(params)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func translateMaterial(s string) string {
	for _, m := range storage.Global.GetMaterials() {
		if m.Key == s {
			return m.Name
		}
	}
	return s
}

func translateShape(s string) string {
	for _, sh := range storage.Global.GetShapes() {
		if sh.Key == s {
			return sh.Name
		}
	}
	return s
}

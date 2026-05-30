package models

type CustomMaterial struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Key           string  `json:"key"`
	Density       float64 `json:"density"`
	YieldStrength float64 `json:"yield_strength"`
}

type CustomShape struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

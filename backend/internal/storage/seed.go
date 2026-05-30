package storage

import "MetalCalc/backend/internal/models"

func Seed() {
	materials := []models.CustomMaterial{
		{Key: "steel", Name: "Сталь"},
		{Key: "aluminum", Name: "Алюміній"},
		{Key: "copper", Name: "Мідь"},
		{Key: "titanium", Name: "Титан"},
	}
	for _, m := range materials {
		Global.AddMaterial(m)
	}

	shapes := []models.CustomShape{
		{Key: "cylinder", Name: "Циліндр"},
		{Key: "rectangle", Name: "Прямокутник"},
	}
	for _, shape := range shapes {
		Global.AddShape(shape)
	}
}

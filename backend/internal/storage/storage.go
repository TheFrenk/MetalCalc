package storage

import (
	"MetalCalc/backend/internal/models"
	"errors"
	"github.com/google/uuid"
	"sync"
)

type Store struct {
	mu        sync.RWMutex
	Materials map[string]models.CustomMaterial
	Shapes    map[string]models.CustomShape
}

var Global = &Store{
	Materials: make(map[string]models.CustomMaterial),
	Shapes:    make(map[string]models.CustomShape),
}

func (s *Store) AddMaterial(m models.CustomMaterial) models.CustomMaterial {
	s.mu.Lock()
	defer s.mu.Unlock()
	m.ID = uuid.NewString()
	s.Materials[m.ID] = m
	return m
}

func (s *Store) GetMaterials() []models.CustomMaterial {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]models.CustomMaterial, 0, len(s.Materials))
	for _, m := range s.Materials {
		list = append(list, m)
	}

	return list
}

func (s *Store) DeleteMaterial(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.Materials[id]; !ok {
		return errors.New("Material not found")
	}
	delete(s.Materials, id)
	return nil
}

func (s *Store) AddShape(sh models.CustomShape) models.CustomShape {
	s.mu.Lock()
	defer s.mu.Unlock()
	sh.ID = uuid.NewString()
	s.Shapes[sh.ID] = sh
	return sh
}

func (s *Store) GetShapes() []models.CustomShape {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]models.CustomShape, 0, len(s.Shapes))
	for _, sh := range s.Shapes {
		list = append(list, sh)
	}
	return list
}

func (s *Store) DeleteShape(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.Shapes[id]; !ok {
		return errors.New("Shape not found")
	}
	delete(s.Shapes, id)
	return nil
}

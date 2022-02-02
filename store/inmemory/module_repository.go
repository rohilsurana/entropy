package inmemory

import (
	"github.com/odpf/entropy/domain"
	"github.com/odpf/entropy/store"
)

type ModuleRepository struct {
	collection map[string]domain.Module
}

func NewModuleRepository() *ModuleRepository {
	return &ModuleRepository{}
}

func (mr *ModuleRepository) Register(urn string, module domain.Module) error {
	if _, exists := mr.collection[urn]; exists {
		return store.ModuleAlreadyExistsError
	}
	mr.collection[urn] = module
	return nil
}

func (mr *ModuleRepository) Get(urn string) (domain.Module, error) {
	if module, exists := mr.collection[urn]; exists {
		return module, nil
	}
	return nil, store.ModuleNotFoundError
}

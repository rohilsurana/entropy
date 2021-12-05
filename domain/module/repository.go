package module

import (
	"github.com/odpf/entropy/domain/model"
)

type Repository struct {
	modules map[string]model.Module
}

func NewRepository() *Repository {
	return &Repository{
		modules: map[string]model.Module{},
	}
}

func (mr *Repository) RegisterModule(m model.Module) {
	mr.modules[m.ID()] = m
}

func (mr *Repository) Get(kind string) model.Module {
	return mr.modules[kind]
}

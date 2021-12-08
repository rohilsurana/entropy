package service

import (
	"github.com/odpf/entropy/domain/module"
	"github.com/odpf/entropy/domain/resource"
	"github.com/odpf/entropy/domain/resource/firehose"
	"github.com/odpf/entropy/pkg/store"
)

type Container struct {
	MR *module.Repository
}

func Init(db *store.DB) (*Container, error) {
	mr := module.NewRepository()
	rc := &resource.Repository{DB: db}
	tm := firehose.Firehose{
		RC: rc,
	}
	mr.RegisterModule(tm)
	return &Container{
		MR: mr,
	}, nil
}

func (container *Container) MigrateAll(db *store.DB) error {
	return nil
}

package firehose

import (
	"strings"

	"github.com/odpf/entropy/domain/model"
	"github.com/odpf/entropy/domain/resource"
)

type Firehose struct {
	RC *resource.Repository
}

const ModuleID = "firehose"

func (t Firehose) ID() string {
	return ModuleID
}

func (t Firehose) Sync(name string, parent string, configs map[string]interface{}) (*model.Resource, error) {
	res := &model.Resource{}
	res.Name = name
	res.URN = strings.Join([]string{parent, name, ModuleID}, "-")
	res.Parent = parent
	res.Kind = ModuleID
	res.Configs = configs
	return t.RC.Create(res)
}

func (t Firehose) Update(urn string, configs map[string]interface{}) (*model.Resource, error) {
	res, err := t.Get(urn)
	if err != nil {
		return nil, err
	}

	res.Configs = configs // take care of merging new configs into old configs
	return t.RC.Update(res)
}

func (t Firehose) Get(urn string) (*model.Resource, error) {
	return t.RC.GetResourceByURN(urn)
}

func (t Firehose) List(parent string) ([]*model.Resource, error) {
	return t.RC.GetResources(ModuleID, parent)
}

func (t Firehose) Delete(urn string) error {
	return nil
}

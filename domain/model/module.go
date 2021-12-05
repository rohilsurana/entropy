package model

type Module interface {
	ID() string
	Create(name string, parent string, config map[string]interface{}) (*Resource, error)
	Update(urn string, configs map[string]interface{}) (*Resource, error)
	Get(urn string) (*Resource, error)
	List(parent string) ([]*Resource, error)
	Delete(urn string) error
}

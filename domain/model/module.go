package model

type Module interface {
	ID() string
	Create(name string, parent string, config map[string]interface{}) (*Resource, error)
	Update(urn string, configs map[string]interface{}) (*Resource, error)
	Get(urn string) (*Resource, error)
	List(parent string) ([]*Resource, error)
	Delete(urn string) error


	Apply(urn string, configs map[string]interface{}) (state map[string]interface{}, status string, err error)
	ValidateConfigs() (err error)
	EnrichConfigs(payload map[string]interface{}) (configs map[string]interface{}, err error)
	Actions() map[string]Action

}

type Action interface {
	ID() string
	Apply(configs map[string]interface{}) (finalConfigs map[string]interface{}, err error)
}

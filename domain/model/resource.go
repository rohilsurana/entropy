package model

type Resource struct {
	ID      string
	URN     string
	Name    string
	Parent  string
	Kind    string
	Configs map[string]interface{}
}

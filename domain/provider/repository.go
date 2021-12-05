package provider

import (
	"os"

	"github.com/odpf/entropy/domain/provider/helm"
)

type Repository struct{}

type Provider interface {
	ID() string
}

func NewRepository() Repository {
	return Repository{}
}

func (r Repository) GetProvider(id string, path string) Provider {
	switch path {
	default:
		var envKubeAPIServer = os.Getenv("TEST_K8S_API_SERVER")
		var envKubeSAToken = os.Getenv("TEST_K8S_SA_TOKEN")

		providerConfig := helm.DefaultProviderConfig()
		providerConfig.Kubernetes.Host = envKubeAPIServer
		providerConfig.Kubernetes.Insecure = true
		providerConfig.Kubernetes.Token = envKubeSAToken
		return helm.NewProvider(providerConfig)
	}
}

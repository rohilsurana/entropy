package handlersv1

import (
	"context"

	"github.com/odpf/entropy/domain/model"
	"github.com/odpf/entropy/service"
	entropy "go.buf.build/odpf/gwv/rohilsurana/proton/odpf/entropy/v1beta1"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
)

type APIServer struct {
	entropy.UnimplementedResourceServiceServer
	container *service.Container
}

func NewApiServer(container *service.Container) *APIServer {
	return &APIServer{
		container: container,
	}
}

func (server APIServer) ListResources(ctx context.Context, request *entropy.ListResourcesRequest) (*entropy.ListResourcesResponse, error) {
	module := server.container.MR.Get(request.Kind)

	res, err := module.List(request.Parent)
	if err != nil {
		return nil, err
	}

	var resources []*entropy.Resource
	for _, item := range res {
		resource, err := resourceToProto(item)
		if err != nil {
			return nil, err
		}
		resources = append(resources, resource)
	}

	response := entropy.ListResourcesResponse{
		Resources: resources,
	}
	return &response, nil
}

func (server APIServer) GetResource(ctx context.Context, request *entropy.GetResourceRequest) (*entropy.GetResourceResponse, error) {
	module := server.container.MR.Get(request.Kind)

	res, err := module.Get(request.Urn)
	if err != nil {
		return nil, err
	}
	resource, err := resourceToProto(res)
	if err != nil {
		return nil, err
	}

	response := entropy.GetResourceResponse{
		Resource: resource,
	}
	return &response, nil
}

func (server APIServer) CreateResource(ctx context.Context, request *entropy.CreateResourceRequest) (*entropy.CreateResourceResponse, error) {
	module := server.container.MR.Get(request.Resource.Kind)
	res, err := module.Create(request.Resource.Name, request.Resource.Parent, request.Resource.Configs.GetStructValue().AsMap())
	if err != nil {
		return nil, err
	}
	resource, err := resourceToProto(res)
	if err != nil {
		return nil, err
	}

	response := entropy.CreateResourceResponse{
		Resource: resource,
	}
	return &response, nil
}

func (server APIServer) UpdateResource(ctx context.Context, request *entropy.UpdateResourceRequest) (*entropy.UpdateResourceResponse, error) {
	module := server.container.MR.Get(request.Resource.Kind)
	res, err := module.Update(request.Resource.Urn, request.Resource.Configs.GetStructValue().AsMap())
	if err != nil {
		return nil, err
	}
	resource, err := resourceToProto(res)
	if err != nil {
		return nil, err
	}

	response := entropy.UpdateResourceResponse{
		Resource: resource,
	}
	return &response, nil
}

func (server APIServer) DeleteResource(ctx context.Context, request *entropy.DeleteResourceRequest) (*entropy.DeleteResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteResource not implemented")
}

func resourceToProto(res *model.Resource) (*entropy.Resource, error) {
	conf, err := structpb.NewValue(res.Configs)
	if err != nil {
		return nil, err
	}
	return &entropy.Resource{
		Id:      res.ID,
		Urn:     res.URN,
		Name:    res.Name,
		Parent:  res.Parent,
		Kind:    res.Kind,
		Configs: conf,
		Status:  ":D",
	}, nil
}

package handlersv1

import (
	"context"
	"errors"
	"github.com/odpf/entropy/domain"
	"github.com/odpf/entropy/mocks"
	"github.com/odpf/entropy/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	entropyv1beta1 "go.buf.build/odpf/gwv/odpf/proton/odpf/entropy/v1beta1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reflect"
	"testing"
	"time"
)

func TestAPIServer_CreateResource(t *testing.T) {
	t.Run("test create new resource", func(t *testing.T) {
		createdAt := time.Now()
		updatedAt := createdAt.Add(time.Minute)
		configsStructValue, _ := structpb.NewValue(map[string]interface{}{
			"replicas": "10",
		})
		want := &entropyv1beta1.CreateResourceResponse{
			Resource: &entropyv1beta1.Resource{
				Urn:       "p-testdata-gl-testname-log",
				Name:      "testname",
				Parent:    "p-testdata-gl",
				Kind:      "log",
				Configs:   configsStructValue,
				Labels:    nil,
				Status:    entropyv1beta1.Resource_STATUS_COMPLETED,
				CreatedAt: timestamppb.New(createdAt),
				UpdatedAt: timestamppb.New(updatedAt),
			},
		}
		wantErr := error(nil)

		ctx := context.Background()
		request := &entropyv1beta1.CreateResourceRequest{
			Resource: &entropyv1beta1.Resource{
				Name:    "testname",
				Parent:  "p-testdata-gl",
				Kind:    "log",
				Configs: configsStructValue,
				Labels:  nil,
			},
		}

		resourceService := &mocks.ResourceService{}
		resourceService.EXPECT().CreateResource(mock.Anything, mock.Anything).Run(func(ctx context.Context, res *domain.Resource) {
			assert.Equal(t, "p-testdata-gl-testname-log", res.Urn)
		}).Return(&domain.Resource{
			Urn:    "p-testdata-gl-testname-log",
			Name:   "testname",
			Parent: "p-testdata-gl",
			Kind:   "log",
			Configs: map[string]interface{}{
				"replicas": "10",
			},
			Labels:    nil,
			Status:    domain.ResourceStatusPending,
			CreatedAt: createdAt,
			UpdatedAt: createdAt,
		}, nil).Once()

		resourceService.EXPECT().UpdateResource(mock.Anything, mock.Anything).Run(func(ctx context.Context, res *domain.Resource) {
			assert.Equal(t, "p-testdata-gl-testname-log", res.Urn)
			assert.Equal(t, domain.ResourceStatusCompleted, res.Status)
		}).Return(&domain.Resource{
			Urn:    "p-testdata-gl-testname-log",
			Name:   "testname",
			Parent: "p-testdata-gl",
			Kind:   "log",
			Configs: map[string]interface{}{
				"replicas": "10",
			},
			Labels:    nil,
			Status:    domain.ResourceStatusCompleted,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}, nil).Once()

		moduleService := &mocks.ModuleService{}
		moduleService.EXPECT().Sync(mock.Anything, mock.Anything).Return(&domain.Resource{
			Urn:    "p-testdata-gl-testname-log",
			Name:   "testname",
			Parent: "p-testdata-gl",
			Kind:   "log",
			Configs: map[string]interface{}{
				"replicas": "10",
			},
			Labels:    nil,
			Status:    domain.ResourceStatusCompleted,
			CreatedAt: createdAt,
			UpdatedAt: createdAt,
		}, nil)

		server := NewApiServer(resourceService, moduleService)
		got, err := server.CreateResource(ctx, request)
		if !errors.Is(err, wantErr) {
			t.Errorf("CreateResource() error = %v, wantErr %v", err, wantErr)
			return
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("CreateResource() got = %v, want %v", got, want)
		}
	})

	t.Run("test create duplicate resource", func(t *testing.T) {
		configsStructValue, _ := structpb.NewValue(map[string]interface{}{
			"replicas": "10",
		})
		want := (*entropyv1beta1.CreateResourceResponse)(nil)
		wantErr := status.Error(codes.AlreadyExists, "resource already exists")

		ctx := context.Background()
		request := &entropyv1beta1.CreateResourceRequest{
			Resource: &entropyv1beta1.Resource{
				Name:    "testname",
				Parent:  "p-testdata-gl",
				Kind:    "log",
				Configs: configsStructValue,
				Labels:  nil,
			},
		}

		resourceService := &mocks.ResourceService{}

		resourceService.EXPECT().
			CreateResource(mock.Anything, mock.Anything).
			Return(nil, store.ResourceAlreadyExistsError).
			Once()

		moduleService := &mocks.ModuleService{}

		server := NewApiServer(resourceService, moduleService)
		got, err := server.CreateResource(ctx, request)
		if !errors.Is(err, wantErr) {
			t.Errorf("CreateResource() error = %v, wantErr %v", err, wantErr)
			return
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("CreateResource() got = %v, want %v", got, want)
		}
	})

	t.Run("test create resource of unknown kind", func(t *testing.T) {
		createdAt := time.Now()
		updatedAt := createdAt.Add(time.Minute)
		configsStructValue, _ := structpb.NewValue(map[string]interface{}{
			"replicas": "10",
		})
		want := (*entropyv1beta1.CreateResourceResponse)(nil)
		wantErr := status.Error(codes.Internal, "failed to find module to deploy this kind")

		ctx := context.Background()
		request := &entropyv1beta1.CreateResourceRequest{
			Resource: &entropyv1beta1.Resource{
				Name:    "testname",
				Parent:  "p-testdata-gl",
				Kind:    "unknown",
				Configs: configsStructValue,
				Labels:  nil,
			},
		}

		resourceService := &mocks.ResourceService{}

		resourceService.EXPECT().CreateResource(mock.Anything, mock.Anything).Return(&domain.Resource{
			Urn:    "p-testdata-gl-testname-unknown",
			Name:   "testname",
			Parent: "p-testdata-gl",
			Kind:   "unkown",
			Configs: map[string]interface{}{
				"replicas": "10",
			},
			Labels:    nil,
			Status:    domain.ResourceStatusPending,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}, nil).Once()

		moduleService := &mocks.ModuleService{}
		moduleService.EXPECT().Sync(mock.Anything, mock.Anything).Return(&domain.Resource{
			Urn:    "p-testdata-gl-testname-unknown",
			Name:   "testname",
			Parent: "p-testdata-gl",
			Kind:   "unknown",
			Configs: map[string]interface{}{
				"replicas": "10",
			},
			Labels:    nil,
			Status:    domain.ResourceStatusError,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}, store.ModuleNotFoundError)

		server := NewApiServer(resourceService, moduleService)
		got, err := server.CreateResource(ctx, request)
		if !errors.Is(err, wantErr) {
			t.Errorf("CreateResource() error = %v, wantErr %v", err, wantErr)
			return
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("CreateResource() got = %v, want %v", got, want)
		}
	})
}

func TestAPIServer_UpdateResource(t *testing.T) {
	t.Run("test update existing resource", func(t *testing.T) {
		createdAt := time.Now()
		updatedAt := createdAt.Add(time.Minute)
		configsStructValue, _ := structpb.NewValue(map[string]interface{}{
			"replicas": "10",
		})
		want := &entropyv1beta1.UpdateResourceResponse{
			Resource: &entropyv1beta1.Resource{
				Urn:       "p-testdata-gl-testname-log",
				Name:      "testname",
				Parent:    "p-testdata-gl",
				Kind:      "log",
				Configs:   configsStructValue,
				Labels:    nil,
				Status:    entropyv1beta1.Resource_STATUS_COMPLETED,
				CreatedAt: timestamppb.New(createdAt),
				UpdatedAt: timestamppb.New(updatedAt),
			},
		}
		wantErr := error(nil)

		ctx := context.Background()
		request := &entropyv1beta1.UpdateResourceRequest{
			Urn:     "p-testdata-gl-testname-log",
			Configs: configsStructValue,
		}

		resourceService := &mocks.ResourceService{}
		resourceService.EXPECT().
			GetResource(mock.Anything, "p-testdata-gl-testname-log").
			Return(&domain.Resource{
				Urn:    "p-testdata-gl-testname-log",
				Name:   "testname",
				Parent: "p-testdata-gl",
				Kind:   "log",
				Configs: map[string]interface{}{
					"replicas": "9",
				},
				Labels:    nil,
				Status:    domain.ResourceStatusCompleted,
				CreatedAt: createdAt,
				UpdatedAt: createdAt,
			}, nil).Once()

		resourceService.EXPECT().
			UpdateResource(mock.Anything, mock.Anything).
			Run(func(ctx context.Context, res *domain.Resource) {
				assert.Equal(t, domain.ResourceStatusPending, res.Status)
			}).
			Return(&domain.Resource{
				Urn:    "p-testdata-gl-testname-log",
				Name:   "testname",
				Parent: "p-testdata-gl",
				Kind:   "log",
				Configs: map[string]interface{}{
					"replicas": "10",
				},
				Labels:    nil,
				Status:    domain.ResourceStatusPending,
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			}, nil).Once()

		resourceService.EXPECT().
			UpdateResource(mock.Anything, mock.Anything).
			Run(func(ctx context.Context, res *domain.Resource) {
				assert.Equal(t, domain.ResourceStatusCompleted, res.Status)
			}).
			Return(&domain.Resource{
				Urn:    "p-testdata-gl-testname-log",
				Name:   "testname",
				Parent: "p-testdata-gl",
				Kind:   "log",
				Configs: map[string]interface{}{
					"replicas": "10",
				},
				Labels:    nil,
				Status:    domain.ResourceStatusCompleted,
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			}, nil).Once()

		moduleService := &mocks.ModuleService{}
		moduleService.EXPECT().Sync(mock.Anything, mock.Anything).Return(&domain.Resource{
			Urn:    "p-testdata-gl-testname-log",
			Name:   "testname",
			Parent: "p-testdata-gl",
			Kind:   "log",
			Configs: map[string]interface{}{
				"replicas": "10",
			},
			Labels:    nil,
			Status:    domain.ResourceStatusCompleted,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}, nil)

		server := NewApiServer(resourceService, moduleService)
		got, err := server.UpdateResource(ctx, request)
		if !errors.Is(err, wantErr) {
			t.Errorf("UpdateResource() error = %v, wantErr %v", err, wantErr)
			return
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("UpdateResource() got = %v, want %v", got, want)
		}
	})

	t.Run("test update non-existing resource", func(t *testing.T) {
		configsStructValue, _ := structpb.NewValue(map[string]interface{}{
			"replicas": "10",
		})
		want := (*entropyv1beta1.UpdateResourceResponse)(nil)
		wantErr := status.Error(codes.NotFound, "could not find resource with given urn")

		ctx := context.Background()
		request := &entropyv1beta1.UpdateResourceRequest{
			Urn:     "p-testdata-gl-testname-log",
			Configs: configsStructValue,
		}

		resourceService := &mocks.ResourceService{}

		resourceService.EXPECT().
			GetResource(mock.Anything, mock.Anything).
			Return(nil, store.ResourceNotFoundError).Once()

		moduleService := &mocks.ModuleService{}

		server := NewApiServer(resourceService, moduleService)
		got, err := server.UpdateResource(ctx, request)
		if !errors.Is(err, wantErr) {
			t.Errorf("UpdateResource() error = %v, wantErr %v", err, wantErr)
			return
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("UpdateResource() got = %v, want %v", got, want)
		}
	})

	t.Run("test update resource with unknown kind", func(t *testing.T) {
		configsStructValue, _ := structpb.NewValue(map[string]interface{}{
			"replicas": "10",
		})
		want := (*entropyv1beta1.UpdateResourceResponse)(nil)
		wantErr := status.Error(codes.Internal, "failed to find module to deploy this kind")

		ctx := context.Background()
		request := &entropyv1beta1.UpdateResourceRequest{
			Urn:     "p-testdata-gl-testname-log",
			Configs: configsStructValue,
		}

		resourceService := &mocks.ResourceService{}
		resourceService.EXPECT().
			UpdateResource(mock.Anything, mock.Anything).
			Return(&domain.Resource{
				Urn: "p-testdata-gl-testname-log",
			}, nil).Once()
		resourceService.EXPECT().
			GetResource(mock.Anything, mock.Anything).
			Return(&domain.Resource{
				Urn: "p-testdata-gl-testname-log",
			}, nil).Once()

		moduleService := &mocks.ModuleService{}
		moduleService.EXPECT().Sync(mock.Anything, mock.Anything).Return(nil, store.ModuleNotFoundError)

		server := NewApiServer(resourceService, moduleService)
		got, err := server.UpdateResource(ctx, request)
		if !errors.Is(err, wantErr) {
			t.Errorf("UpdateResource() error = %v, wantErr %v", err, wantErr)
			return
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("UpdateResource() got = %v, want %v", got, want)
		}
	})
}

func TestAPIServer_GetResource(t *testing.T) {
	t.Run("test get resource", func(t *testing.T) {
		r := &domain.Resource{
			Urn:       "p-testdata-gl-testname-mock",
			Name:      "testname",
			Parent:    "p-testdata-gl",
			Kind:      "mock",
			Configs:   map[string]interface{}{},
			Labels:    map[string]string{},
			Status:    domain.ResourceStatusCompleted,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		rProto, _ := resourceToProto(r)
		argsRequest := &entropyv1beta1.GetResourceRequest{
			Urn: "p-testdata-gl-testname-mock",
		}
		want := &entropyv1beta1.GetResourceResponse{
			Resource: rProto,
		}
		wantErr := error(nil)

		mockResourceService := &mocks.ResourceService{}
		mockResourceService.EXPECT().GetResource(mock.Anything, mock.Anything).Return(r, nil).Once()

		mockModuleService := &mocks.ModuleService{}

		server := APIServer{
			resourceService: mockResourceService,
			moduleService:   mockModuleService,
		}
		got, err := server.GetResource(context.TODO(), argsRequest)
		if !errors.Is(err, wantErr) {
			t.Errorf("GetResource() error = %v, wantErr %v", err, wantErr)
			return
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetResource() got = %v, want %v", got, want)
		}
	})

	t.Run("test get non existent resource", func(t *testing.T) {
		argsRequest := &entropyv1beta1.GetResourceRequest{
			Urn: "p-testdata-gl-testname-mock",
		}
		want := (*entropyv1beta1.GetResourceResponse)(nil)
		wantErr := status.Error(codes.NotFound, "could not find resource with given urn")

		mockResourceService := &mocks.ResourceService{}
		mockResourceService.EXPECT().GetResource(mock.Anything, mock.Anything).Return(nil, store.ResourceNotFoundError).Once()

		mockModuleService := &mocks.ModuleService{}

		server := APIServer{
			resourceService: mockResourceService,
			moduleService:   mockModuleService,
		}
		got, err := server.GetResource(context.TODO(), argsRequest)
		if !errors.Is(err, wantErr) {
			t.Errorf("GetResource() error = %v, wantErr %v", err, wantErr)
			return
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetResource() got = %v, want %v", got, want)
		}
	})
}

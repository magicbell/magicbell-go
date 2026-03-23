package broadcasts

import (
	"context"
	restClient "github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest/hooks"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest/httptransport"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/configmanager"
	"github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
	"github.com/magicbell/magicbell-go/pkg/project-client/shared"
	"time"
)

// BroadcastsService provides methods to interact with BroadcastsService-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type BroadcastsService struct {
	manager *configmanager.ConfigManager
	hook    hooks.Hook
}

func NewBroadcastsService() *BroadcastsService {
	return &BroadcastsService{
		manager: configmanager.NewConfigManager(clientconfig.Config{}),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *BroadcastsService) WithConfigManager(manager *configmanager.ConfigManager) *BroadcastsService {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *BroadcastsService) WithHook(hook hooks.Hook) *BroadcastsService {
	api.hook = hook
	return api
}

func (api *BroadcastsService) getConfig() *clientconfig.Config {
	return api.manager.GetBroadcasts()
}

func (api *BroadcastsService) getHook() hooks.Hook {
	return api.hook
}

func (api *BroadcastsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *BroadcastsService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *BroadcastsService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// Retrieves a paginated list of broadcasts for the project. Returns basic information about each broadcast including its creation time and status.
func (api *BroadcastsService) ListBroadcasts(ctx context.Context, params ListBroadcastsRequestParams) (*shared.ClientResponse[BroadcastCollection], *shared.ClientError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/broadcasts").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[BroadcastCollection, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[[]byte](err)
	}

	return shared.NewClientResponse[BroadcastCollection](resp), nil
}

// Creates a new broadcast. When a broadcast is created, it generates individual notifications for relevant users within the project.
func (api *BroadcastsService) CreateBroadcast(ctx context.Context, broadcast Broadcast) (*shared.ClientResponse[Broadcast], *shared.ClientError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/broadcasts").
		WithConfig(config).
		WithBody(broadcast).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Broadcast, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[[]byte](err)
	}

	return shared.NewClientResponse[Broadcast](resp), nil
}

// Retrieves detailed information about a specific broadcast by its ID. Includes the broadcast's configuration and current status.
func (api *BroadcastsService) FetchBroadcast(ctx context.Context, broadcastId string) (*shared.ClientResponse[Broadcast], *shared.ClientError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/broadcasts/{broadcast_id}").
		WithConfig(config).
		AddPathParam("broadcast_id", broadcastId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Broadcast, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[[]byte](err)
	}

	return shared.NewClientResponse[Broadcast](resp), nil
}

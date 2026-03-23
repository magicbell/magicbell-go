package events

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

// EventsService provides methods to interact with EventsService-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type EventsService struct {
	manager *configmanager.ConfigManager
	hook    hooks.Hook
}

func NewEventsService() *EventsService {
	return &EventsService{
		manager: configmanager.NewConfigManager(clientconfig.Config{}),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *EventsService) WithConfigManager(manager *configmanager.ConfigManager) *EventsService {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *EventsService) WithHook(hook hooks.Hook) *EventsService {
	api.hook = hook
	return api
}

func (api *EventsService) getConfig() *clientconfig.Config {
	return api.manager.GetEvents()
}

func (api *EventsService) getHook() hooks.Hook {
	return api.hook
}

func (api *EventsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *EventsService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *EventsService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// Retrieves a paginated list of events for the project.
func (api *EventsService) ListEvents(ctx context.Context, params ListEventsRequestParams) (*shared.ClientResponse[EventCollection], *shared.ClientError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/events").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[EventCollection, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[[]byte](err)
	}

	return shared.NewClientResponse[EventCollection](resp), nil
}

// Fetches a project event by its ID.
func (api *EventsService) FetchEvent(ctx context.Context, eventId string) (*shared.ClientResponse[Event], *shared.ClientError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/events/{event_id}").
		WithConfig(config).
		AddPathParam("event_id", eventId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Event, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[[]byte](err)
	}

	return shared.NewClientResponse[Event](resp), nil
}

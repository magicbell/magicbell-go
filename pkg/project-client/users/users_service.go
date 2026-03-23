package users

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

// UsersService provides methods to interact with UsersService-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type UsersService struct {
	manager *configmanager.ConfigManager
	hook    hooks.Hook
}

func NewUsersService() *UsersService {
	return &UsersService{
		manager: configmanager.NewConfigManager(clientconfig.Config{}),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *UsersService) WithConfigManager(manager *configmanager.ConfigManager) *UsersService {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *UsersService) WithHook(hook hooks.Hook) *UsersService {
	api.hook = hook
	return api
}

func (api *UsersService) getConfig() *clientconfig.Config {
	return api.manager.GetUsers()
}

func (api *UsersService) getHook() hooks.Hook {
	return api.hook
}

func (api *UsersService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *UsersService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *UsersService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// Lists all users in the project.
func (api *UsersService) ListUsers(ctx context.Context, params ListUsersRequestParams) (*shared.ClientResponse[UserCollection], *shared.ClientError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[UserCollection, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[[]byte](err)
	}

	return shared.NewClientResponse[UserCollection](resp), nil
}

// Creates or updates a user with the provided details. The user will be associated with the project specified in the request context.
func (api *UsersService) SaveUser(ctx context.Context, user shared.User) (*shared.ClientResponse[shared.User], *shared.ClientError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/users").
		WithConfig(config).
		WithBody(user).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[shared.User, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[[]byte](err)
	}

	return shared.NewClientResponse[shared.User](resp), nil
}

// Removes a user and all associated data from the project.
func (api *UsersService) DeleteUser(ctx context.Context, userId string) (*shared.ClientResponse[any], *shared.ClientError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/users/{user_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[[]byte](err)
	}

	return shared.NewClientResponse[any](resp), nil
}

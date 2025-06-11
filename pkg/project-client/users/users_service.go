package users

import (
	"context"
	restClient "github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest/httptransport"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/configmanager"
	"github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
	"github.com/magicbell/magicbell-go/pkg/project-client/shared"
	"time"
)

type UsersService struct {
	manager *configmanager.ConfigManager
}

func NewUsersService() *UsersService {
	return &UsersService{
		manager: configmanager.NewConfigManager(clientconfig.Config{}),
	}
}

func (api *UsersService) WithConfigManager(manager *configmanager.ConfigManager) *UsersService {
	api.manager = manager
	return api
}

func (api *UsersService) getConfig() *clientconfig.Config {
	return api.manager.GetUsers()
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
func (api *UsersService) ListUsers(ctx context.Context, params ListUsersRequestParams) (*shared.ClientResponse[UserCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[UserCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[UserCollection](err)
	}

	return shared.NewClientResponse[UserCollection](resp), nil
}

// Creates or updates a user with the provided details. The user will be associated with the project specified in the request context.
func (api *UsersService) SaveUser(ctx context.Context, user shared.User) (*shared.ClientResponse[shared.User], *shared.ClientError) {
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

	client := restClient.NewRestClient[shared.User](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[shared.User](err)
	}

	return shared.NewClientResponse[shared.User](resp), nil
}

// Removes a user and all associated data from the project.
func (api *UsersService) DeleteUser(ctx context.Context, userId string) (*shared.ClientResponse[any], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/users/{user_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[any](err)
	}

	return shared.NewClientResponse[any](resp), nil
}

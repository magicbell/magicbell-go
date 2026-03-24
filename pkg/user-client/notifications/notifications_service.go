package notifications

import (
	"context"
	restClient "github.com/magicbell/magicbell-go/pkg/user-client/internal/clients/rest"
	"github.com/magicbell/magicbell-go/pkg/user-client/internal/clients/rest/hooks"
	"github.com/magicbell/magicbell-go/pkg/user-client/internal/clients/rest/httptransport"
	"github.com/magicbell/magicbell-go/pkg/user-client/internal/configmanager"
	"github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
	"github.com/magicbell/magicbell-go/pkg/user-client/shared"
	"time"
)

// NotificationsService provides methods to interact with NotificationsService-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type NotificationsService struct {
	manager *configmanager.ConfigManager
	hook    hooks.Hook
}

func NewNotificationsService() *NotificationsService {
	return &NotificationsService{
		manager: configmanager.NewConfigManager(clientconfig.Config{}),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *NotificationsService) WithConfigManager(manager *configmanager.ConfigManager) *NotificationsService {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *NotificationsService) WithHook(hook hooks.Hook) *NotificationsService {
	api.hook = hook
	return api
}

func (api *NotificationsService) getConfig() *clientconfig.Config {
	return api.manager.GetNotifications()
}

func (api *NotificationsService) getHook() hooks.Hook {
	return api.hook
}

func (api *NotificationsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *NotificationsService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *NotificationsService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// Lists all notifications for a user.
func (api *NotificationsService) ListNotifications(ctx context.Context, params ListNotificationsRequestParams) (*shared.ClientResponse[NotificationCollection], *shared.ClientError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/notifications").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[NotificationCollection, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[[]byte](err)
	}

	return shared.NewClientResponse[NotificationCollection](resp), nil
}

// Archive all notifications.
func (api *NotificationsService) ArchiveAllNotifications(ctx context.Context, params ArchiveAllNotificationsRequestParams) (*shared.ClientResponse[any], *shared.ClientError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/notifications/archive").
		WithConfig(config).
		WithOptions(params).
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

// Marks all notifications as read.
func (api *NotificationsService) MarkAllNotificationsRead(ctx context.Context, params MarkAllNotificationsReadRequestParams) (*shared.ClientResponse[any], *shared.ClientError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/notifications/read").
		WithConfig(config).
		WithOptions(params).
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

// Returns the count of unread notifications for a user. Supports filtering by category and topic.
func (api *NotificationsService) FetchUnreadNotificationsCount(ctx context.Context, params FetchUnreadNotificationsCountRequestParams) (*shared.ClientResponse[CountResponse], *shared.ClientError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/notifications/unread/count").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[CountResponse, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[[]byte](err)
	}

	return shared.NewClientResponse[CountResponse](resp), nil
}

// Gets a notification by ID.
func (api *NotificationsService) FetchNotification(ctx context.Context, notificationId string) (*shared.ClientResponse[Notification], *shared.ClientError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/notifications/{notification_id}").
		WithConfig(config).
		AddPathParam("notification_id", notificationId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Notification, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[[]byte](err)
	}

	return shared.NewClientResponse[Notification](resp), nil
}

// Archive a notification.
func (api *NotificationsService) ArchiveNotification(ctx context.Context, notificationId string) (*shared.ClientResponse[any], *shared.ClientError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/notifications/{notification_id}/archive").
		WithConfig(config).
		AddPathParam("notification_id", notificationId).
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

// Marks a notification as read.
func (api *NotificationsService) MarkNotificationRead(ctx context.Context, notificationId string) (*shared.ClientResponse[any], *shared.ClientError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/notifications/{notification_id}/read").
		WithConfig(config).
		AddPathParam("notification_id", notificationId).
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

// Unarchives a notification.
func (api *NotificationsService) UnarchiveNotification(ctx context.Context, notificationId string) (*shared.ClientResponse[any], *shared.ClientError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/notifications/{notification_id}/unarchive").
		WithConfig(config).
		AddPathParam("notification_id", notificationId).
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

// Marks a notification as unread.
func (api *NotificationsService) MarkNotificationUnread(ctx context.Context, notificationId string) (*shared.ClientResponse[any], *shared.ClientError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/notifications/{notification_id}/unread").
		WithConfig(config).
		AddPathParam("notification_id", notificationId).
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

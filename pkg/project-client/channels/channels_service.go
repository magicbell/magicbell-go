package channels

import (
	"context"
	restClient "github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest/httptransport"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/configmanager"
	"github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
	"github.com/magicbell/magicbell-go/pkg/project-client/shared"
	"time"
)

type ChannelsService struct {
	manager *configmanager.ConfigManager
}

func NewChannelsService() *ChannelsService {
	return &ChannelsService{
		manager: configmanager.NewConfigManager(clientconfig.Config{}),
	}
}

func (api *ChannelsService) WithConfigManager(manager *configmanager.ConfigManager) *ChannelsService {
	api.manager = manager
	return api
}

func (api *ChannelsService) getConfig() *clientconfig.Config {
	return api.manager.GetChannels()
}

func (api *ChannelsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *ChannelsService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *ChannelsService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// Save the channels configuration for a given key.
func (api *ChannelsService) SaveChannelsConfig(ctx context.Context, categoryDeliveryConfig CategoryDeliveryConfig) (*shared.ClientResponse[CategoryDeliveryConfig], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/channels").
		WithConfig(config).
		WithBody(categoryDeliveryConfig).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[CategoryDeliveryConfig](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[CategoryDeliveryConfig](err)
	}

	return shared.NewClientResponse[CategoryDeliveryConfig](resp), nil
}

// Fetches the channels config for a given key.
func (api *ChannelsService) FetchChannelsConfig(ctx context.Context, key string) (*shared.ClientResponse[CategoryDeliveryConfig], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/{key}").
		WithConfig(config).
		AddPathParam("key", key).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[CategoryDeliveryConfig](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[CategoryDeliveryConfig](err)
	}

	return shared.NewClientResponse[CategoryDeliveryConfig](resp), nil
}

// Lists all Inbox tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.
func (api *ChannelsService) ListUserInboxTokens(ctx context.Context, userId string, params ListUserInboxTokensRequestParams) (*shared.ClientResponse[InboxTokenResponseCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/in_app/inbox/tokens").
		WithConfig(config).
		AddPathParam("user_id", userId).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[InboxTokenResponseCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[InboxTokenResponseCollection](err)
	}

	return shared.NewClientResponse[InboxTokenResponseCollection](resp), nil
}

// Fetches a specific Inbox token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.
func (api *ChannelsService) FetchUserInboxToken(ctx context.Context, userId string, tokenId string) (*shared.ClientResponse[InboxTokenResponse], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/in_app/inbox/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[InboxTokenResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[InboxTokenResponse](err)
	}

	return shared.NewClientResponse[InboxTokenResponse](resp), nil
}

// Deletes a specific user's Inbox token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.
func (api *ChannelsService) DeleteUserInboxToken(ctx context.Context, userId string, tokenId string) (*shared.ClientResponse[DiscardResult], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/users/{user_id}/channels/in_app/inbox/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DiscardResult](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[DiscardResult](err)
	}

	return shared.NewClientResponse[DiscardResult](resp), nil
}

// Lists all APNs tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.
func (api *ChannelsService) ListUserApnsTokens(ctx context.Context, userId string, params ListUserApnsTokensRequestParams) (*shared.ClientResponse[ApnsTokenCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/mobile_push/apns/tokens").
		WithConfig(config).
		AddPathParam("user_id", userId).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ApnsTokenCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[ApnsTokenCollection](err)
	}

	return shared.NewClientResponse[ApnsTokenCollection](resp), nil
}

// Fetches a specific APNs token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.
func (api *ChannelsService) FetchUserApnsToken(ctx context.Context, userId string, tokenId string) (*shared.ClientResponse[ApnsToken], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/mobile_push/apns/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ApnsToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[ApnsToken](err)
	}

	return shared.NewClientResponse[ApnsToken](resp), nil
}

// Deletes a specific user's APNs token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.
func (api *ChannelsService) DeleteUserApnsToken(ctx context.Context, userId string, tokenId string) (*shared.ClientResponse[DiscardResult], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/users/{user_id}/channels/mobile_push/apns/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DiscardResult](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[DiscardResult](err)
	}

	return shared.NewClientResponse[DiscardResult](resp), nil
}

// Lists all Expo tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.
func (api *ChannelsService) ListUserExpoTokens(ctx context.Context, userId string, params ListUserExpoTokensRequestParams) (*shared.ClientResponse[ExpoTokenCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/mobile_push/expo/tokens").
		WithConfig(config).
		AddPathParam("user_id", userId).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ExpoTokenCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[ExpoTokenCollection](err)
	}

	return shared.NewClientResponse[ExpoTokenCollection](resp), nil
}

// Fetches a specific Expo token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.
func (api *ChannelsService) FetchUserExpoToken(ctx context.Context, userId string, tokenId string) (*shared.ClientResponse[ExpoToken], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/mobile_push/expo/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ExpoToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[ExpoToken](err)
	}

	return shared.NewClientResponse[ExpoToken](resp), nil
}

// Deletes a specific user's Expo token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.
func (api *ChannelsService) DeleteUserExpoToken(ctx context.Context, userId string, tokenId string) (*shared.ClientResponse[DiscardResult], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/users/{user_id}/channels/mobile_push/expo/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DiscardResult](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[DiscardResult](err)
	}

	return shared.NewClientResponse[DiscardResult](resp), nil
}

// Lists all FCM tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.
func (api *ChannelsService) ListUserFcmTokens(ctx context.Context, userId string, params ListUserFcmTokensRequestParams) (*shared.ClientResponse[FcmTokenCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/mobile_push/fcm/tokens").
		WithConfig(config).
		AddPathParam("user_id", userId).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[FcmTokenCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[FcmTokenCollection](err)
	}

	return shared.NewClientResponse[FcmTokenCollection](resp), nil
}

// Fetches a specific FCM token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.
func (api *ChannelsService) FetchUserFcmToken(ctx context.Context, userId string, tokenId string) (*shared.ClientResponse[FcmToken], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/mobile_push/fcm/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[FcmToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[FcmToken](err)
	}

	return shared.NewClientResponse[FcmToken](resp), nil
}

// Deletes a specific user's FCM token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.
func (api *ChannelsService) DeleteUserFcmToken(ctx context.Context, userId string, tokenId string) (*shared.ClientResponse[DiscardResult], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/users/{user_id}/channels/mobile_push/fcm/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DiscardResult](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[DiscardResult](err)
	}

	return shared.NewClientResponse[DiscardResult](resp), nil
}

// Lists all Slack tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.
func (api *ChannelsService) ListUserSlackTokens(ctx context.Context, userId string, params ListUserSlackTokensRequestParams) (*shared.ClientResponse[SlackTokenCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/slack/tokens").
		WithConfig(config).
		AddPathParam("user_id", userId).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SlackTokenCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[SlackTokenCollection](err)
	}

	return shared.NewClientResponse[SlackTokenCollection](resp), nil
}

// Fetches a specific Slack token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.
func (api *ChannelsService) FetchUserSlackToken(ctx context.Context, userId string, tokenId string) (*shared.ClientResponse[SlackToken], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/slack/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SlackToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[SlackToken](err)
	}

	return shared.NewClientResponse[SlackToken](resp), nil
}

// Deletes a specific user's Slack token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.
func (api *ChannelsService) DeleteUserSlackToken(ctx context.Context, userId string, tokenId string) (*shared.ClientResponse[DiscardResult], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/users/{user_id}/channels/slack/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DiscardResult](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[DiscardResult](err)
	}

	return shared.NewClientResponse[DiscardResult](resp), nil
}

// Lists all Teams tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.
func (api *ChannelsService) ListUserTeamsTokens(ctx context.Context, userId string, params ListUserTeamsTokensRequestParams) (*shared.ClientResponse[TeamsTokenCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/teams/tokens").
		WithConfig(config).
		AddPathParam("user_id", userId).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[TeamsTokenCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[TeamsTokenCollection](err)
	}

	return shared.NewClientResponse[TeamsTokenCollection](resp), nil
}

// Fetches a specific Teams token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.
func (api *ChannelsService) FetchUserTeamsToken(ctx context.Context, userId string, tokenId string) (*shared.ClientResponse[TeamsToken], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/teams/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[TeamsToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[TeamsToken](err)
	}

	return shared.NewClientResponse[TeamsToken](resp), nil
}

// Deletes a specific user's Teams token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.
func (api *ChannelsService) DeleteUserTeamsToken(ctx context.Context, userId string, tokenId string) (*shared.ClientResponse[DiscardResult], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/users/{user_id}/channels/teams/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DiscardResult](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[DiscardResult](err)
	}

	return shared.NewClientResponse[DiscardResult](resp), nil
}

// Lists all Web Push tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.
func (api *ChannelsService) ListUserWebPushTokens(ctx context.Context, userId string, params ListUserWebPushTokensRequestParams) (*shared.ClientResponse[WebPushTokenCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/web_push/tokens").
		WithConfig(config).
		AddPathParam("user_id", userId).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[WebPushTokenCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[WebPushTokenCollection](err)
	}

	return shared.NewClientResponse[WebPushTokenCollection](resp), nil
}

// Fetches a specific Web Push token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.
func (api *ChannelsService) FetchUserWebPushToken(ctx context.Context, userId string, tokenId string) (*shared.ClientResponse[WebPushToken], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/web_push/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[WebPushToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[WebPushToken](err)
	}

	return shared.NewClientResponse[WebPushToken](resp), nil
}

// Deletes a specific user's Web Push token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.
func (api *ChannelsService) DeleteUserWebPushToken(ctx context.Context, userId string, tokenId string) (*shared.ClientResponse[DiscardResult], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/users/{user_id}/channels/web_push/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DiscardResult](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[DiscardResult](err)
	}

	return shared.NewClientResponse[DiscardResult](resp), nil
}

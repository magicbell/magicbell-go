package channels

import (
	"context"
	restClient "github.com/magicbell/magicbell-go/pkg/user-client/internal/clients/rest"
	"github.com/magicbell/magicbell-go/pkg/user-client/internal/clients/rest/httptransport"
	"github.com/magicbell/magicbell-go/pkg/user-client/internal/configmanager"
	"github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
	"github.com/magicbell/magicbell-go/pkg/user-client/shared"
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

// Lists all Inbox tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.
func (api *ChannelsService) ListInboxTokens(ctx context.Context, params ListInboxTokensRequestParams) (*shared.ClientResponse[InboxTokenResponseCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/in_app/inbox/tokens").
		WithConfig(config).
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

// Saves the Inbox token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.
func (api *ChannelsService) SaveInboxToken(ctx context.Context, inboxToken InboxToken) (*shared.ClientResponse[InboxToken], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/channels/in_app/inbox/tokens").
		WithConfig(config).
		WithBody(inboxToken).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[InboxToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[InboxToken](err)
	}

	return shared.NewClientResponse[InboxToken](resp), nil
}

// Fetches details of a specific Inbox token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.
func (api *ChannelsService) FetchInboxToken(ctx context.Context, tokenId string) (*shared.ClientResponse[InboxTokenResponse], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/in_app/inbox/tokens/{token_id}").
		WithConfig(config).
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

// Deletes one of the authenticated user's Inbox tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.
func (api *ChannelsService) DeleteInboxToken(ctx context.Context, tokenId string) (*shared.ClientResponse[DiscardResult], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/channels/in_app/inbox/tokens/{token_id}").
		WithConfig(config).
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

// Lists all APNs tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.
func (api *ChannelsService) ListApnsTokens(ctx context.Context, params ListApnsTokensRequestParams) (*shared.ClientResponse[ApnsTokenCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/mobile_push/apns/tokens").
		WithConfig(config).
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

// Saves the APNs token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.
func (api *ChannelsService) SaveApnsToken(ctx context.Context, apnsTokenPayload ApnsTokenPayload) (*shared.ClientResponse[ApnsTokenPayload], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/channels/mobile_push/apns/tokens").
		WithConfig(config).
		WithBody(apnsTokenPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ApnsTokenPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[ApnsTokenPayload](err)
	}

	return shared.NewClientResponse[ApnsTokenPayload](resp), nil
}

// Fetches details of a specific APNs token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.
func (api *ChannelsService) FetchApnsToken(ctx context.Context, tokenId string) (*shared.ClientResponse[ApnsToken], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/mobile_push/apns/tokens/{token_id}").
		WithConfig(config).
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

// Deletes one of the authenticated user's APNs tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.
func (api *ChannelsService) DeleteApnsToken(ctx context.Context, tokenId string) (*shared.ClientResponse[DiscardResult], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/channels/mobile_push/apns/tokens/{token_id}").
		WithConfig(config).
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

// Lists all Expo tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.
func (api *ChannelsService) ListExpoTokens(ctx context.Context, params ListExpoTokensRequestParams) (*shared.ClientResponse[ExpoTokenCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/mobile_push/expo/tokens").
		WithConfig(config).
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

// Saves the Expo token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.
func (api *ChannelsService) SaveExpoToken(ctx context.Context, expoTokenPayload ExpoTokenPayload) (*shared.ClientResponse[ExpoTokenPayload], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/channels/mobile_push/expo/tokens").
		WithConfig(config).
		WithBody(expoTokenPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ExpoTokenPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[ExpoTokenPayload](err)
	}

	return shared.NewClientResponse[ExpoTokenPayload](resp), nil
}

// Fetches details of a specific Expo token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.
func (api *ChannelsService) FetchExpoToken(ctx context.Context, tokenId string) (*shared.ClientResponse[ExpoToken], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/mobile_push/expo/tokens/{token_id}").
		WithConfig(config).
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

// Deletes one of the authenticated user's Expo tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.
func (api *ChannelsService) DeleteExpoToken(ctx context.Context, tokenId string) (*shared.ClientResponse[DiscardResult], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/channels/mobile_push/expo/tokens/{token_id}").
		WithConfig(config).
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

// Lists all FCM tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.
func (api *ChannelsService) ListFcmTokens(ctx context.Context, params ListFcmTokensRequestParams) (*shared.ClientResponse[FcmTokenCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/mobile_push/fcm/tokens").
		WithConfig(config).
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

// Saves the FCM token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.
func (api *ChannelsService) SaveFcmToken(ctx context.Context, fcmTokenPayload FcmTokenPayload) (*shared.ClientResponse[FcmTokenPayload], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/channels/mobile_push/fcm/tokens").
		WithConfig(config).
		WithBody(fcmTokenPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[FcmTokenPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[FcmTokenPayload](err)
	}

	return shared.NewClientResponse[FcmTokenPayload](resp), nil
}

// Fetches details of a specific FCM token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.
func (api *ChannelsService) FetchFcmToken(ctx context.Context, tokenId string) (*shared.ClientResponse[FcmToken], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/mobile_push/fcm/tokens/{token_id}").
		WithConfig(config).
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

// Deletes one of the authenticated user's FCM tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.
func (api *ChannelsService) DeleteFcmToken(ctx context.Context, tokenId string) (*shared.ClientResponse[DiscardResult], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/channels/mobile_push/fcm/tokens/{token_id}").
		WithConfig(config).
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

// Lists all Slack tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.
func (api *ChannelsService) ListSlackTokens(ctx context.Context, params ListSlackTokensRequestParams) (*shared.ClientResponse[SlackTokenCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/slack/tokens").
		WithConfig(config).
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

// Saves the Slack token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.
func (api *ChannelsService) SaveSlackToken(ctx context.Context, slackTokenPayload SlackTokenPayload) (*shared.ClientResponse[SlackTokenPayload], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/channels/slack/tokens").
		WithConfig(config).
		WithBody(slackTokenPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SlackTokenPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[SlackTokenPayload](err)
	}

	return shared.NewClientResponse[SlackTokenPayload](resp), nil
}

// Fetches details of a specific Slack token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.
func (api *ChannelsService) FetchSlackToken(ctx context.Context, tokenId string) (*shared.ClientResponse[SlackToken], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/slack/tokens/{token_id}").
		WithConfig(config).
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

// Deletes one of the authenticated user's Slack tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.
func (api *ChannelsService) DeleteSlackToken(ctx context.Context, tokenId string) (*shared.ClientResponse[DiscardResult], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/channels/slack/tokens/{token_id}").
		WithConfig(config).
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

// Lists all Teams tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.
func (api *ChannelsService) ListTeamsTokens(ctx context.Context, params ListTeamsTokensRequestParams) (*shared.ClientResponse[TeamsTokenCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/teams/tokens").
		WithConfig(config).
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

// Saves the Teams token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.
func (api *ChannelsService) SaveTeamsToken(ctx context.Context, teamsTokenPayload TeamsTokenPayload) (*shared.ClientResponse[TeamsTokenPayload], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/channels/teams/tokens").
		WithConfig(config).
		WithBody(teamsTokenPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[TeamsTokenPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[TeamsTokenPayload](err)
	}

	return shared.NewClientResponse[TeamsTokenPayload](resp), nil
}

// Fetches details of a specific Teams token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.
func (api *ChannelsService) FetchTeamsToken(ctx context.Context, tokenId string) (*shared.ClientResponse[TeamsToken], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/teams/tokens/{token_id}").
		WithConfig(config).
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

// Deletes one of the authenticated user's Teams tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.
func (api *ChannelsService) DeleteTeamsToken(ctx context.Context, tokenId string) (*shared.ClientResponse[DiscardResult], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/channels/teams/tokens/{token_id}").
		WithConfig(config).
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

// Lists all Web Push tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.
func (api *ChannelsService) ListWebPushTokens(ctx context.Context, params ListWebPushTokensRequestParams) (*shared.ClientResponse[WebPushTokenCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/web_push/tokens").
		WithConfig(config).
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

// Saves the Web Push token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.
func (api *ChannelsService) SaveWebPushToken(ctx context.Context, webPushTokenPayload shared.WebPushTokenPayload) (*shared.ClientResponse[shared.WebPushTokenPayload], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/channels/web_push/tokens").
		WithConfig(config).
		WithBody(webPushTokenPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[shared.WebPushTokenPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[shared.WebPushTokenPayload](err)
	}

	return shared.NewClientResponse[shared.WebPushTokenPayload](resp), nil
}

// Fetches details of a specific Web Push token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.
func (api *ChannelsService) FetchWebPushToken(ctx context.Context, tokenId string) (*shared.ClientResponse[WebPushToken], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/web_push/tokens/{token_id}").
		WithConfig(config).
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

// Deletes one of the authenticated user's Web Push tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.
func (api *ChannelsService) DeleteWebPushToken(ctx context.Context, tokenId string) (*shared.ClientResponse[DiscardResult], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/channels/web_push/tokens/{token_id}").
		WithConfig(config).
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

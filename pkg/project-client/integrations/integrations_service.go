package integrations

import (
	"context"
	restClient "github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest/httptransport"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/configmanager"
	"github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
	"github.com/magicbell/magicbell-go/pkg/project-client/shared"
	"time"
)

type IntegrationsService struct {
	manager *configmanager.ConfigManager
}

func NewIntegrationsService() *IntegrationsService {
	return &IntegrationsService{
		manager: configmanager.NewConfigManager(clientconfig.Config{}),
	}
}

func (api *IntegrationsService) WithConfigManager(manager *configmanager.ConfigManager) *IntegrationsService {
	api.manager = manager
	return api
}

func (api *IntegrationsService) getConfig() *clientconfig.Config {
	return api.manager.GetIntegrations()
}

func (api *IntegrationsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *IntegrationsService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *IntegrationsService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// Lists all available and configured integrations for the project. Returns a summary of each integration including its type, status, and basic configuration information.
func (api *IntegrationsService) ListIntegrations(ctx context.Context, params ListIntegrationsRequestParams) (*shared.ClientResponse[IntegrationConfigCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[IntegrationConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[IntegrationConfigCollection](err)
	}

	return shared.NewClientResponse[IntegrationConfigCollection](resp), nil
}

// Retrieves the current APNs integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) ListApnsIntegrations(ctx context.Context) (*shared.ClientResponse[ApnsConfigCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/apns").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ApnsConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[ApnsConfigCollection](err)
	}

	return shared.NewClientResponse[ApnsConfigCollection](resp), nil
}

// Updates or creates the APNs integration for the project.
func (api *IntegrationsService) SaveApnsIntegration(ctx context.Context, apnsConfigPayload ApnsConfigPayload) (*shared.ClientResponse[ApnsConfigPayload], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/apns").
		WithConfig(config).
		WithBody(apnsConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ApnsConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[ApnsConfigPayload](err)
	}

	return shared.NewClientResponse[ApnsConfigPayload](resp), nil
}

// Deletes the APNs integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteApnsIntegration(ctx context.Context, params DeleteApnsIntegrationRequestParams) (*shared.ClientResponse[any], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/apns").
		WithConfig(config).
		WithOptions(params).
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

// Retrieves the current Expo integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) ListExpoIntegrations(ctx context.Context) (*shared.ClientResponse[ExpoConfigCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/expo").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ExpoConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[ExpoConfigCollection](err)
	}

	return shared.NewClientResponse[ExpoConfigCollection](resp), nil
}

// Updates or creates the Expo integration for the project.
func (api *IntegrationsService) SaveExpoIntegration(ctx context.Context, expoConfigPayload ExpoConfigPayload) (*shared.ClientResponse[ExpoConfigPayload], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/expo").
		WithConfig(config).
		WithBody(expoConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ExpoConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[ExpoConfigPayload](err)
	}

	return shared.NewClientResponse[ExpoConfigPayload](resp), nil
}

// Deletes the Expo integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteExpoIntegration(ctx context.Context, params DeleteExpoIntegrationRequestParams) (*shared.ClientResponse[any], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/expo").
		WithConfig(config).
		WithOptions(params).
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

// Retrieves the current FCM integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) ListFcmIntegrations(ctx context.Context) (*shared.ClientResponse[FcmConfigCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/fcm").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[FcmConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[FcmConfigCollection](err)
	}

	return shared.NewClientResponse[FcmConfigCollection](resp), nil
}

// Updates or creates the FCM integration for the project.
func (api *IntegrationsService) SaveFcmIntegration(ctx context.Context, fcmConfigPayload FcmConfigPayload) (*shared.ClientResponse[FcmConfigPayload], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/fcm").
		WithConfig(config).
		WithBody(fcmConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[FcmConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[FcmConfigPayload](err)
	}

	return shared.NewClientResponse[FcmConfigPayload](resp), nil
}

// Deletes the FCM integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteFcmIntegration(ctx context.Context, params DeleteFcmIntegrationRequestParams) (*shared.ClientResponse[any], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/fcm").
		WithConfig(config).
		WithOptions(params).
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

// Retrieves the current Inbox integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) ListInboxIntegrations(ctx context.Context) (*shared.ClientResponse[InboxConfigCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/inbox").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[InboxConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[InboxConfigCollection](err)
	}

	return shared.NewClientResponse[InboxConfigCollection](resp), nil
}

// Updates or creates the Inbox integration for the project.
func (api *IntegrationsService) SaveInboxIntegration(ctx context.Context, inboxConfigPayload InboxConfigPayload) (*shared.ClientResponse[InboxConfigPayload], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/inbox").
		WithConfig(config).
		WithBody(inboxConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[InboxConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[InboxConfigPayload](err)
	}

	return shared.NewClientResponse[InboxConfigPayload](resp), nil
}

// Deletes the Inbox integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteInboxIntegration(ctx context.Context, params DeleteInboxIntegrationRequestParams) (*shared.ClientResponse[any], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/inbox").
		WithConfig(config).
		WithOptions(params).
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

// Retrieves the current Mailgun integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) ListMailgunIntegrations(ctx context.Context) (*shared.ClientResponse[MailgunConfigCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/mailgun").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[MailgunConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[MailgunConfigCollection](err)
	}

	return shared.NewClientResponse[MailgunConfigCollection](resp), nil
}

// Updates or creates the Mailgun integration for the project.
func (api *IntegrationsService) SaveMailgunIntegration(ctx context.Context, mailgunConfigPayload MailgunConfigPayload) (*shared.ClientResponse[MailgunConfigPayload], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/mailgun").
		WithConfig(config).
		WithBody(mailgunConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[MailgunConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[MailgunConfigPayload](err)
	}

	return shared.NewClientResponse[MailgunConfigPayload](resp), nil
}

// Deletes the Mailgun integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteMailgunIntegration(ctx context.Context, params DeleteMailgunIntegrationRequestParams) (*shared.ClientResponse[any], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/mailgun").
		WithConfig(config).
		WithOptions(params).
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

// Retrieves the current Ping Email integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) ListPingEmailIntegrations(ctx context.Context) (*shared.ClientResponse[PingConfigCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/ping_email").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[PingConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[PingConfigCollection](err)
	}

	return shared.NewClientResponse[PingConfigCollection](resp), nil
}

// Updates or creates the Ping Email integration for the project.
func (api *IntegrationsService) SavePingEmailIntegration(ctx context.Context, pingConfigPayload PingConfigPayload) (*shared.ClientResponse[PingConfigPayload], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/ping_email").
		WithConfig(config).
		WithBody(pingConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[PingConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[PingConfigPayload](err)
	}

	return shared.NewClientResponse[PingConfigPayload](resp), nil
}

// Deletes the Ping Email integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeletePingEmailIntegration(ctx context.Context, params DeletePingEmailIntegrationRequestParams) (*shared.ClientResponse[any], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/ping_email").
		WithConfig(config).
		WithOptions(params).
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

// Retrieves the current SendGrid integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) ListSendgridIntegrations(ctx context.Context) (*shared.ClientResponse[SendgridConfigCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/sendgrid").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SendgridConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[SendgridConfigCollection](err)
	}

	return shared.NewClientResponse[SendgridConfigCollection](resp), nil
}

// Updates or creates the SendGrid integration for the project.
func (api *IntegrationsService) SaveSendgridIntegration(ctx context.Context, sendgridConfigPayload SendgridConfigPayload) (*shared.ClientResponse[SendgridConfigPayload], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/sendgrid").
		WithConfig(config).
		WithBody(sendgridConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SendgridConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[SendgridConfigPayload](err)
	}

	return shared.NewClientResponse[SendgridConfigPayload](resp), nil
}

// Deletes the SendGrid integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteSendgridIntegration(ctx context.Context, params DeleteSendgridIntegrationRequestParams) (*shared.ClientResponse[any], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/sendgrid").
		WithConfig(config).
		WithOptions(params).
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

// Retrieves the current Amazon SES integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) ListSesIntegrations(ctx context.Context) (*shared.ClientResponse[SesConfigCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/ses").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SesConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[SesConfigCollection](err)
	}

	return shared.NewClientResponse[SesConfigCollection](resp), nil
}

// Updates or creates the Amazon SES integration for the project.
func (api *IntegrationsService) SaveSesIntegration(ctx context.Context, sesConfigPayload SesConfigPayload) (*shared.ClientResponse[SesConfigPayload], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/ses").
		WithConfig(config).
		WithBody(sesConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SesConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[SesConfigPayload](err)
	}

	return shared.NewClientResponse[SesConfigPayload](resp), nil
}

// Deletes the Amazon SES integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteSesIntegration(ctx context.Context, params DeleteSesIntegrationRequestParams) (*shared.ClientResponse[any], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/ses").
		WithConfig(config).
		WithOptions(params).
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

// Retrieves the current Slack integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) ListSlackIntegrations(ctx context.Context) (*shared.ClientResponse[SlackConfigCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/slack").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SlackConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[SlackConfigCollection](err)
	}

	return shared.NewClientResponse[SlackConfigCollection](resp), nil
}

// Updates or creates the Slack integration for the project.
func (api *IntegrationsService) SaveSlackIntegration(ctx context.Context, slackConfigPayload SlackConfigPayload) (*shared.ClientResponse[SlackConfigPayload], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/slack").
		WithConfig(config).
		WithBody(slackConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SlackConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[SlackConfigPayload](err)
	}

	return shared.NewClientResponse[SlackConfigPayload](resp), nil
}

// Deletes the Slack integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteSlackIntegration(ctx context.Context, params DeleteSlackIntegrationRequestParams) (*shared.ClientResponse[any], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/slack").
		WithConfig(config).
		WithOptions(params).
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

// Retrieves the current Twilio integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) ListTwilioIntegrations(ctx context.Context) (*shared.ClientResponse[TwilioConfigCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/twilio").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[TwilioConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[TwilioConfigCollection](err)
	}

	return shared.NewClientResponse[TwilioConfigCollection](resp), nil
}

// Updates or creates the Twilio integration for the project.
func (api *IntegrationsService) SaveTwilioIntegration(ctx context.Context, twilioConfigPayload TwilioConfigPayload) (*shared.ClientResponse[TwilioConfigPayload], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/twilio").
		WithConfig(config).
		WithBody(twilioConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[TwilioConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[TwilioConfigPayload](err)
	}

	return shared.NewClientResponse[TwilioConfigPayload](resp), nil
}

// Deletes the Twilio integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteTwilioIntegration(ctx context.Context, params DeleteTwilioIntegrationRequestParams) (*shared.ClientResponse[any], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/twilio").
		WithConfig(config).
		WithOptions(params).
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

// Retrieves the current Web Push integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) ListWebPushIntegrations(ctx context.Context) (*shared.ClientResponse[WebpushConfigCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/web_push").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[WebpushConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[WebpushConfigCollection](err)
	}

	return shared.NewClientResponse[WebpushConfigCollection](resp), nil
}

// Updates or creates the Web Push integration for the project.
func (api *IntegrationsService) SaveWebPushIntegration(ctx context.Context, webpushConfigPayload WebpushConfigPayload) (*shared.ClientResponse[WebpushConfigPayload], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/web_push").
		WithConfig(config).
		WithBody(webpushConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[WebpushConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[WebpushConfigPayload](err)
	}

	return shared.NewClientResponse[WebpushConfigPayload](resp), nil
}

// Deletes the Web Push integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteWebPushIntegration(ctx context.Context, params DeleteWebPushIntegrationRequestParams) (*shared.ClientResponse[any], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/web_push").
		WithConfig(config).
		WithOptions(params).
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

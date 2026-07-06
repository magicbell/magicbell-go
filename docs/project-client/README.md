---
title: "project-client"
---

# Setup & Configuration

## Supported Language Versions

This SDK is compatible with the following versions: `Go >= 1.19.0`

## Installation

To get started with the SDK, we recommend installing using `go get`:

```bash
go get client
```

## Authentication

### Access Token Authentication

The Client API uses an Access Token for authentication.

This token must be provided to authenticate your requests to the API.

#### Setting the Access Token

When you initialize the SDK, you can set the access token as follows:

```go
import (
    "github.com/magicbell/magicbell-go/pkg/project-client/client"
    "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  )

config := clientconfig.NewConfig()
config.SetAccessToken("YOUR-TOKEN")

sdk := client.NewClient(config)
```

If you need to set or update the access token after initializing the SDK, you can use:

```go
import (
    "github.com/magicbell/magicbell-go/pkg/project-client/client"
    "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  )

config := clientconfig.NewConfig()

sdk := client.NewClient(config)
sdk.SetAccessToken("YOUR-TOKEN")
```

## Setting a Custom Timeout

You can set a custom timeout for the SDK's HTTP requests as follows:

```go
import "time"

config := clientconfig.NewConfig()

sdk := client.NewClient(config)

sdk.SetTimeout(10 * time.Second)
```

# Sample Usage

Below is a comprehensive example demonstrating how to authenticate and call a simple endpoint:

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/broadcasts"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := broadcasts.ListBroadcastsRequestParams{
  Limit: util.ToPointer(int64(0)),
  StartingAfter: util.ToPointer("starting_after"),
  EndingBefore: util.ToPointer("ending_before"),
}

response, err := client.Broadcasts.ListBroadcasts(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)

```

## Services

The SDK provides various services to interact with the API.

<details>
<summary>Below is a list of all available services with links to their detailed documentation:</summary>

| Name                                                                  |
| :-------------------------------------------------------------------- |
| [BroadcastsService](services/broadcasts_service.md)     |
| [ChannelsService](services/channels_service.md)         |
| [EventsService](services/events_service.md)             |
| [IntegrationsService](services/integrations_service.md) |
| [UsersService](services/users_service.md)               |
| [WorkflowsService](services/workflows_service.md)       |

</details>

### Response Wrappers

All services use response wrappers to provide a consistent interface to return the responses from the API.

The response wrapper itself is a generic struct that contains the response data and metadata.

<details>
<summary>Below are the response wrappers used in the SDK:</summary>

#### `ClientResponse[T]`

This response wrapper is used to return the response data from the API. It contains the following fields:

| Name     | Type                     | Description                                 |
| :------- | :----------------------- | :------------------------------------------ |
| Data     | `T`                      | The body of the API response                |
| Metadata | `ClientResponseMetadata` | Status code and headers returned by the API |

#### `ClientError[T]`

This response wrapper is used to return an error. It contains the following fields:

| Name     | Type                  | Description                                                       |
| :------- | :-------------------- | :---------------------------------------------------------------- |
| Err      | `error`               | The error that occurred                                           |
| Data     | `*T`                  | The deserialized error response data (nil if unmarshaling failed) |
| Body     | `[]byte`              | The raw body of the API response                                  |
| Metadata | `ClientErrorMetadata` | Status code and headers returned by the API                       |

#### `ClientResponseMetadata`

This struct is shared by both response wrappers and contains the following fields:

| Name       | Type                | Description                                      |
| :--------- | :------------------ | :----------------------------------------------- |
| Headers    | `map[string]string` | A map containing the headers returned by the API |
| StatusCode | `int`               | The status code returned by the API              |

</details>

## Models

The SDK includes several models that represent the data structures used in API requests and responses. These models help in organizing and managing the data efficiently.

<details>
<summary>Below is a list of all available models with links to their detailed documentation:</summary>

| Name                                                                                    | Description |
| :-------------------------------------------------------------------------------------- | :---------- |
| [BroadcastCollection](models/broadcast_collection.md)                     |             |
| [Broadcast](models/broadcast.md)                                          |             |
| [User](models/user.md)                                                    |             |
| [Links](models/links.md)                                                  |             |
| [CategoryDeliveryConfig](models/category_delivery_config.md)              |             |
| [InboxTokenResponseCollection](models/inbox_token_response_collection.md) |             |
| [InboxTokenResponse](models/inbox_token_response.md)                      |             |
| [Links](models/links.md)                                                  |             |
| [DiscardResult](models/discard_result.md)                                 |             |
| [ApnsTokenCollection](models/apns_token_collection.md)                    |             |
| [ApnsToken](models/apns_token.md)                                         |             |
| [ExpoTokenCollection](models/expo_token_collection.md)                    |             |
| [ExpoToken](models/expo_token.md)                                         |             |
| [FcmTokenCollection](models/fcm_token_collection.md)                      |             |
| [FcmToken](models/fcm_token.md)                                           |             |
| [SlackTokenCollection](models/slack_token_collection.md)                  |             |
| [SlackToken](models/slack_token.md)                                       |             |
| [TeamsTokenCollection](models/teams_token_collection.md)                  |             |
| [TeamsToken](models/teams_token.md)                                       |             |
| [WebPushTokenCollection](models/web_push_token_collection.md)             |             |
| [WebPushToken](models/web_push_token.md)                                  |             |
| [EventCollection](models/event_collection.md)                             |             |
| [Event](models/event.md)                                                  |             |
| [Links](models/links.md)                                                  |             |
| [IntegrationConfigCollection](models/integration_config_collection.md)    |             |
| [IntegrationConfig](models/integration_config.md)                         |             |
| [Links](models/links.md)                                                  |             |
| [ApnsConfigCollection](models/apns_config_collection.md)                  |             |
| [ApnsConfig](models/apns_config.md)                                       |             |
| [ApnsConfigPayload](models/apns_config_payload.md)                        |             |
| [EventSourceConfigCollection](models/event_source_config_collection.md)   |             |
| [EventSourceConfig](models/event_source_config.md)                        |             |
| [EventSourceConfigPayload](models/event_source_config_payload.md)         |             |
| [ExpoConfigCollection](models/expo_config_collection.md)                  |             |
| [ExpoConfig](models/expo_config.md)                                       |             |
| [ExpoConfigPayload](models/expo_config_payload.md)                        |             |
| [FcmConfigCollection](models/fcm_config_collection.md)                    |             |
| [FcmConfig](models/fcm_config.md)                                         |             |
| [FcmConfigPayload](models/fcm_config_payload.md)                          |             |
| [GithubConfigCollection](models/github_config_collection.md)              |             |
| [GithubConfig](models/github_config.md)                                   |             |
| [GithubConfigPayload](models/github_config_payload.md)                    |             |
| [InboxConfigCollection](models/inbox_config_collection.md)                |             |
| [InboxConfig](models/inbox_config.md)                                     |             |
| [InboxConfigPayload](models/inbox_config_payload.md)                      |             |
| [SlackBotConfigCollection](models/slack_bot_config_collection.md)         |             |
| [SlackBotConfig](models/slack_bot_config.md)                              |             |
| [SlackBotConfigPayload](models/slack_bot_config_payload.md)               |             |
| [MailgunConfigCollection](models/mailgun_config_collection.md)            |             |
| [MailgunConfig](models/mailgun_config.md)                                 |             |
| [MailgunConfigPayload](models/mailgun_config_payload.md)                  |             |
| [PingConfigCollection](models/ping_config_collection.md)                  |             |
| [PingConfig](models/ping_config.md)                                       |             |
| [PingConfigPayload](models/ping_config_payload.md)                        |             |
| [SendgridConfigCollection](models/sendgrid_config_collection.md)          |             |
| [SendgridConfig](models/sendgrid_config.md)                               |             |
| [SendgridConfigPayload](models/sendgrid_config_payload.md)                |             |
| [SesConfigCollection](models/ses_config_collection.md)                    |             |
| [SesConfig](models/ses_config.md)                                         |             |
| [SesConfigPayload](models/ses_config_payload.md)                          |             |
| [SlackConfigCollection](models/slack_config_collection.md)                |             |
| [SlackConfig](models/slack_config.md)                                     |             |
| [SlackConfigPayload](models/slack_config_payload.md)                      |             |
| [SmtpConfigObjectCollection](models/smtp_config_object_collection.md)     |             |
| [SmtpConfigObject](models/smtp_config_object.md)                          |             |
| [SmtpConfig](models/smtp_config.md)                                       |             |
| [StripeConfigCollection](models/stripe_config_collection.md)              |             |
| [StripeConfig](models/stripe_config.md)                                   |             |
| [StripeConfigPayload](models/stripe_config_payload.md)                    |             |
| [TwilioConfigCollection](models/twilio_config_collection.md)              |             |
| [TwilioConfig](models/twilio_config.md)                                   |             |
| [TwilioConfigPayload](models/twilio_config_payload.md)                    |             |
| [WebpushConfigCollection](models/webpush_config_collection.md)            |             |
| [WebpushConfig](models/webpush_config.md)                                 |             |
| [WebpushConfigPayload](models/webpush_config_payload.md)                  |             |
| [UserCollection](models/user_collection.md)                               |             |
| [User](models/user.md)                                                    |             |
| [Links](models/links.md)                                                  |             |
| [WorkflowList](models/workflow_list.md)                                   |             |
| [WorkflowDefinition](models/workflow_definition.md)                       |             |
| [CreateRunResponse](models/create_run_response.md)                        |             |
| [ExecuteWorkflowRequest](models/execute_workflow_request.md)              |             |
| [GetRunResponse](models/get_run_response.md)                              |             |
| [WorkflowRunCollection](models/workflow_run_collection.md)                |             |
| [WorkflowRun](models/workflow_run.md)                                     |             |
| [Links](models/links.md)                                                  |             |

</details>

<!-- This file was generated by liblab | https://liblab.com/ -->

---
title: "user-client"
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
    "github.com/magicbell/magicbell-go/pkg/user-client/client"
    "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  )

config := clientconfig.NewConfig()
config.SetAccessToken("YOUR-TOKEN")

sdk := client.NewClient(config)
```

If you need to set or update the access token after initializing the SDK, you can use:

```go
import (
    "github.com/magicbell/magicbell-go/pkg/user-client/client"
    "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
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
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/channels"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := channels.ListInboxTokensRequestParams{
  Limit: util.ToPointer(int64(8)),
  StartingAfter: util.ToPointer("starting_after"),
  EndingBefore: util.ToPointer("ending_before"),
}

response, err := client.Channels.ListInboxTokens(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)

```

## Services

The SDK provides various services to interact with the API.

<details>
<summary>Below is a list of all available services with links to their detailed documentation:</summary>

| Name                                                                    |
| :---------------------------------------------------------------------- |
| [ChannelsService](services/channels_service.md)           |
| [IntegrationsService](services/integrations_service.md)   |
| [NotificationsService](services/notifications_service.md) |

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

| Name                                                                                             | Description |
| :----------------------------------------------------------------------------------------------- | :---------- |
| [InboxTokenResponseCollection](models/inbox_token_response_collection.md)          |             |
| [InboxTokenResponse](models/inbox_token_response.md)                               |             |
| [Links](models/links.md)                                                           |             |
| [InboxToken](models/inbox_token.md)                                                |             |
| [DiscardResult](models/discard_result.md)                                          |             |
| [ApnsTokenCollection](models/apns_token_collection.md)                             |             |
| [ApnsToken](models/apns_token.md)                                                  |             |
| [ApnsTokenPayload](models/apns_token_payload.md)                                   |             |
| [ExpoTokenCollection](models/expo_token_collection.md)                             |             |
| [ExpoToken](models/expo_token.md)                                                  |             |
| [ExpoTokenPayload](models/expo_token_payload.md)                                   |             |
| [FcmTokenCollection](models/fcm_token_collection.md)                               |             |
| [FcmToken](models/fcm_token.md)                                                    |             |
| [FcmTokenPayload](models/fcm_token_payload.md)                                     |             |
| [SlackTokenCollection](models/slack_token_collection.md)                           |             |
| [SlackToken](models/slack_token.md)                                                |             |
| [SlackTokenPayload](models/slack_token_payload.md)                                 |             |
| [TeamsTokenCollection](models/teams_token_collection.md)                           |             |
| [TeamsToken](models/teams_token.md)                                                |             |
| [TeamsTokenPayload](models/teams_token_payload.md)                                 |             |
| [UserPreferences](models/user_preferences.md)                                      |             |
| [WebPushTokenCollection](models/web_push_token_collection.md)                      |             |
| [WebPushToken](models/web_push_token.md)                                           |             |
| [WebPushTokenPayload](models/web_push_token_payload.md)                            |             |
| [InboxConfigPayload](models/inbox_config_payload.md)                               |             |
| [SlackInstallation](models/slack_installation.md)                                  |             |
| [SlackFinishInstallResponse](models/slack_finish_install_response.md)              |             |
| [SlackStartInstallResponseContent](models/slack_start_install_response_content.md) |             |
| [SlackStartInstall](models/slack_start_install.md)                                 |             |
| [WebPushTokenPayload](models/web_push_token_payload.md)                            |             |
| [WebPushStartInstallationResponse](models/web_push_start_installation_response.md) |             |
| [NotificationCollection](models/notification_collection.md)                        |             |
| [Notification](models/notification.md)                                             |             |
| [Links](models/links.md)                                                           |             |
| [CountResponse](models/count_response.md)                                          |             |

</details>

<!-- This file was generated by liblab | https://liblab.com/ -->

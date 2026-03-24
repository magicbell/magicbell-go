# ChannelsService

A list of all methods in the `ChannelsService` service. Click on the method name to view detailed information about that method.

| Methods                                                               | Description                                                                                                                                                                                                                                                             |
| :-------------------------------------------------------------------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [SaveChannelsConfig](#savechannelsconfig)                             | Save the channels configuration for a given key.                                                                                                                                                                                                                        |
| [FetchChannelsConfig](#fetchchannelsconfig)                           | Fetches the channels config for a given key.                                                                                                                                                                                                                            |
| [ListUserInboxTokens](#listuserinboxtokens)                           | Lists all Inbox tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.                                                                          |
| [FetchUserInboxToken](#fetchuserinboxtoken)                           | Fetches a specific Inbox token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.              |
| [DeleteUserInboxToken](#deleteuserinboxtoken)                         | Deletes a specific user's Inbox token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.                    |
| [ListUserApnsTokens](#listuserapnstokens)                             | Lists all APNs tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.                                                                           |
| [FetchUserApnsToken](#fetchuserapnstoken)                             | Fetches a specific APNs token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.               |
| [DeleteUserApnsToken](#deleteuserapnstoken)                           | Deletes a specific user's APNs token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.                     |
| [ListUserExpoTokens](#listuserexpotokens)                             | Lists all Expo tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.                                                                           |
| [FetchUserExpoToken](#fetchuserexpotoken)                             | Fetches a specific Expo token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.               |
| [DeleteUserExpoToken](#deleteuserexpotoken)                           | Deletes a specific user's Expo token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.                     |
| [ListUserFcmTokens](#listuserfcmtokens)                               | Lists all FCM tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.                                                                            |
| [FetchUserFcmToken](#fetchuserfcmtoken)                               | Fetches a specific FCM token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.                |
| [DeleteUserFcmToken](#deleteuserfcmtoken)                             | Deletes a specific user's FCM token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.                      |
| [ListUserMagicbellSlackbotTokens](#listusermagicbellslackbottokens)   | Lists all MagicBell SlackBot tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.                                                             |
| [FetchUserMagicbellSlackbotToken](#fetchusermagicbellslackbottoken)   | Fetches a specific MagicBell SlackBot token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata. |
| [DeleteUserMagicbellSlackbotToken](#deleteusermagicbellslackbottoken) | Deletes a specific user's MagicBell SlackBot token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.       |
| [ListUserSlackTokens](#listuserslacktokens)                           | Lists all Slack tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.                                                                          |
| [FetchUserSlackToken](#fetchuserslacktoken)                           | Fetches a specific Slack token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.              |
| [DeleteUserSlackToken](#deleteuserslacktoken)                         | Deletes a specific user's Slack token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.                    |
| [ListUserTeamsTokens](#listuserteamstokens)                           | Lists all Teams tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.                                                                          |
| [FetchUserTeamsToken](#fetchuserteamstoken)                           | Fetches a specific Teams token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.              |
| [DeleteUserTeamsToken](#deleteuserteamstoken)                         | Deletes a specific user's Teams token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.                    |
| [ListUserWebPushTokens](#listuserwebpushtokens)                       | Lists all Web Push tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.                                                                       |
| [FetchUserWebPushToken](#fetchuserwebpushtoken)                       | Fetches a specific Web Push token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.           |
| [DeleteUserWebPushToken](#deleteuserwebpushtoken)                     | Deletes a specific user's Web Push token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.                 |

## SaveChannelsConfig

Save the channels configuration for a given key.

- HTTP Method: `PUT`
- Endpoint: `/channels`

**Parameters**

| Name                   | Type                   | Required | Description                 |
| :--------------------- | :--------------------- | :------- | :-------------------------- |
| ctx                    | Context                | ✅       | Default go language context |
| categoryDeliveryConfig | CategoryDeliveryConfig | ✅       |                             |

**Return Type**

`CategoryDeliveryConfig`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/channels"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

channel := channels.CHANNEL_IN_APP

categoryDeliveryConfigChannels := channels.CategoryDeliveryConfigChannels{
  Channel: &channel,
  Delay: util.ToPointer(int64(8)),
  If_: util.ToPointer(util.Nullable[string]{ Value: "if" }),
}

request := channels.CategoryDeliveryConfig{
  Channels: []channels.CategoryDeliveryConfigChannels{categoryDeliveryConfigChannels},
  Disabled: util.ToPointer(true),
  Key: util.ToPointer("key"),
}

response, err := client.Channels.SaveChannelsConfig(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchChannelsConfig

Fetches the channels config for a given key.

- HTTP Method: `GET`
- Endpoint: `/channels/{key}`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |
| key  | string  | ✅       |                             |

**Return Type**

`CategoryDeliveryConfig`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.FetchChannelsConfig(context.Background(), "key")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListUserInboxTokens

Lists all Inbox tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/in_app/inbox/tokens`

**Parameters**

| Name   | Type                             | Required | Description                   |
| :----- | :------------------------------- | :------- | :---------------------------- |
| ctx    | Context                          | ✅       | Default go language context   |
| userId | string                           | ✅       |                               |
| params | ListUserInboxTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`InboxTokenResponseCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/channels"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := channels.ListUserInboxTokensRequestParams{
  Limit: util.ToPointer(int64(3)),
  StartingAfter: util.ToPointer("starting_after"),
  EndingBefore: util.ToPointer("ending_before"),
}

response, err := client.Channels.ListUserInboxTokens(context.Background(), "user_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchUserInboxToken

Fetches a specific Inbox token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/in_app/inbox/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`InboxTokenResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.FetchUserInboxToken(context.Background(), "user_id", "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteUserInboxToken

Deletes a specific user's Inbox token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.

- HTTP Method: `DELETE`
- Endpoint: `/users/{user_id}/channels/in_app/inbox/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.DeleteUserInboxToken(context.Background(), "user_id", "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListUserApnsTokens

Lists all APNs tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/mobile_push/apns/tokens`

**Parameters**

| Name   | Type                            | Required | Description                   |
| :----- | :------------------------------ | :------- | :---------------------------- |
| ctx    | Context                         | ✅       | Default go language context   |
| userId | string                          | ✅       |                               |
| params | ListUserApnsTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`ApnsTokenCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/channels"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := channels.ListUserApnsTokensRequestParams{
  Limit: util.ToPointer(int64(2)),
  StartingAfter: util.ToPointer("starting_after"),
  EndingBefore: util.ToPointer("ending_before"),
}

response, err := client.Channels.ListUserApnsTokens(context.Background(), "user_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchUserApnsToken

Fetches a specific APNs token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/mobile_push/apns/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`ApnsToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.FetchUserApnsToken(context.Background(), "user_id", "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteUserApnsToken

Deletes a specific user's APNs token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.

- HTTP Method: `DELETE`
- Endpoint: `/users/{user_id}/channels/mobile_push/apns/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.DeleteUserApnsToken(context.Background(), "user_id", "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListUserExpoTokens

Lists all Expo tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/mobile_push/expo/tokens`

**Parameters**

| Name   | Type                            | Required | Description                   |
| :----- | :------------------------------ | :------- | :---------------------------- |
| ctx    | Context                         | ✅       | Default go language context   |
| userId | string                          | ✅       |                               |
| params | ListUserExpoTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`ExpoTokenCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/channels"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := channels.ListUserExpoTokensRequestParams{
  Limit: util.ToPointer(int64(1)),
  StartingAfter: util.ToPointer("starting_after"),
  EndingBefore: util.ToPointer("ending_before"),
}

response, err := client.Channels.ListUserExpoTokens(context.Background(), "user_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchUserExpoToken

Fetches a specific Expo token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/mobile_push/expo/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`ExpoToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.FetchUserExpoToken(context.Background(), "user_id", "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteUserExpoToken

Deletes a specific user's Expo token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.

- HTTP Method: `DELETE`
- Endpoint: `/users/{user_id}/channels/mobile_push/expo/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.DeleteUserExpoToken(context.Background(), "user_id", "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListUserFcmTokens

Lists all FCM tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/mobile_push/fcm/tokens`

**Parameters**

| Name   | Type                           | Required | Description                   |
| :----- | :----------------------------- | :------- | :---------------------------- |
| ctx    | Context                        | ✅       | Default go language context   |
| userId | string                         | ✅       |                               |
| params | ListUserFcmTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`FcmTokenCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/channels"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := channels.ListUserFcmTokensRequestParams{
  Limit: util.ToPointer(int64(1)),
  StartingAfter: util.ToPointer("starting_after"),
  EndingBefore: util.ToPointer("ending_before"),
}

response, err := client.Channels.ListUserFcmTokens(context.Background(), "user_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchUserFcmToken

Fetches a specific FCM token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/mobile_push/fcm/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`FcmToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.FetchUserFcmToken(context.Background(), "user_id", "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteUserFcmToken

Deletes a specific user's FCM token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.

- HTTP Method: `DELETE`
- Endpoint: `/users/{user_id}/channels/mobile_push/fcm/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.DeleteUserFcmToken(context.Background(), "user_id", "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListUserMagicbellSlackbotTokens

Lists all MagicBell SlackBot tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/slack/magicbell_slackbot/tokens`

**Parameters**

| Name   | Type                                         | Required | Description                   |
| :----- | :------------------------------------------- | :------- | :---------------------------- |
| ctx    | Context                                      | ✅       | Default go language context   |
| userId | string                                       | ✅       |                               |
| params | ListUserMagicbellSlackbotTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`SlackTokenCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/channels"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := channels.ListUserMagicbellSlackbotTokensRequestParams{
  Limit: util.ToPointer(int64(6)),
  StartingAfter: util.ToPointer("starting_after"),
  EndingBefore: util.ToPointer("ending_before"),
}

response, err := client.Channels.ListUserMagicbellSlackbotTokens(context.Background(), "user_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchUserMagicbellSlackbotToken

Fetches a specific MagicBell SlackBot token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/slack/magicbell_slackbot/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`SlackToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.FetchUserMagicbellSlackbotToken(context.Background(), "user_id", "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteUserMagicbellSlackbotToken

Deletes a specific user's MagicBell SlackBot token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.

- HTTP Method: `DELETE`
- Endpoint: `/users/{user_id}/channels/slack/magicbell_slackbot/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.DeleteUserMagicbellSlackbotToken(context.Background(), "user_id", "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListUserSlackTokens

Lists all Slack tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/slack/tokens`

**Parameters**

| Name   | Type                             | Required | Description                   |
| :----- | :------------------------------- | :------- | :---------------------------- |
| ctx    | Context                          | ✅       | Default go language context   |
| userId | string                           | ✅       |                               |
| params | ListUserSlackTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`SlackTokenCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/channels"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := channels.ListUserSlackTokensRequestParams{
  Limit: util.ToPointer(int64(4)),
  StartingAfter: util.ToPointer("starting_after"),
  EndingBefore: util.ToPointer("ending_before"),
}

response, err := client.Channels.ListUserSlackTokens(context.Background(), "user_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchUserSlackToken

Fetches a specific Slack token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/slack/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`SlackToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.FetchUserSlackToken(context.Background(), "user_id", "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteUserSlackToken

Deletes a specific user's Slack token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.

- HTTP Method: `DELETE`
- Endpoint: `/users/{user_id}/channels/slack/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.DeleteUserSlackToken(context.Background(), "user_id", "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListUserTeamsTokens

Lists all Teams tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/teams/tokens`

**Parameters**

| Name   | Type                             | Required | Description                   |
| :----- | :------------------------------- | :------- | :---------------------------- |
| ctx    | Context                          | ✅       | Default go language context   |
| userId | string                           | ✅       |                               |
| params | ListUserTeamsTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`TeamsTokenCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/channels"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := channels.ListUserTeamsTokensRequestParams{
  Limit: util.ToPointer(int64(6)),
  StartingAfter: util.ToPointer("starting_after"),
  EndingBefore: util.ToPointer("ending_before"),
}

response, err := client.Channels.ListUserTeamsTokens(context.Background(), "user_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchUserTeamsToken

Fetches a specific Teams token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/teams/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`TeamsToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.FetchUserTeamsToken(context.Background(), "user_id", "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteUserTeamsToken

Deletes a specific user's Teams token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.

- HTTP Method: `DELETE`
- Endpoint: `/users/{user_id}/channels/teams/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.DeleteUserTeamsToken(context.Background(), "user_id", "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListUserWebPushTokens

Lists all Web Push tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/web_push/tokens`

**Parameters**

| Name   | Type                               | Required | Description                   |
| :----- | :--------------------------------- | :------- | :---------------------------- |
| ctx    | Context                            | ✅       | Default go language context   |
| userId | string                             | ✅       |                               |
| params | ListUserWebPushTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`WebPushTokenCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/channels"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := channels.ListUserWebPushTokensRequestParams{
  Limit: util.ToPointer(int64(5)),
  StartingAfter: util.ToPointer("starting_after"),
  EndingBefore: util.ToPointer("ending_before"),
}

response, err := client.Channels.ListUserWebPushTokens(context.Background(), "user_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchUserWebPushToken

Fetches a specific Web Push token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/users/{user_id}/channels/web_push/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`WebPushToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.FetchUserWebPushToken(context.Background(), "user_id", "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteUserWebPushToken

Deletes a specific user's Web Push token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.

- HTTP Method: `DELETE`
- Endpoint: `/users/{user_id}/channels/web_push/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| userId  | string  | ✅       |                             |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.DeleteUserWebPushToken(context.Background(), "user_id", "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

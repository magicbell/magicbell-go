# ChannelsService

A list of all methods in the `ChannelsService` service. Click on the method name to view detailed information about that method.

| Methods                                           | Description                                                                                                                                                                                                                                                   |
| :------------------------------------------------ | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| [SaveChannelsConfig](#savechannelsconfig)         | Save the channels configuration for a given key.                                                                                                                                                                                                              |
| [FetchChannelsConfig](#fetchchannelsconfig)       | Fetches the channels config for a given key.                                                                                                                                                                                                                  |
| [ListUserInboxTokens](#listuserinboxtokens)       | Lists all Inbox tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.                                                                |
| [FetchUserInboxToken](#fetchuserinboxtoken)       | Fetches a specific Inbox token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.    |
| [DeleteUserInboxToken](#deleteuserinboxtoken)     | Deletes a specific user's Inbox token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.          |
| [ListUserApnsTokens](#listuserapnstokens)         | Lists all APNs tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.                                                                 |
| [FetchUserApnsToken](#fetchuserapnstoken)         | Fetches a specific APNs token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.     |
| [DeleteUserApnsToken](#deleteuserapnstoken)       | Deletes a specific user's APNs token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.           |
| [ListUserExpoTokens](#listuserexpotokens)         | Lists all Expo tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.                                                                 |
| [FetchUserExpoToken](#fetchuserexpotoken)         | Fetches a specific Expo token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.     |
| [DeleteUserExpoToken](#deleteuserexpotoken)       | Deletes a specific user's Expo token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.           |
| [ListUserFcmTokens](#listuserfcmtokens)           | Lists all FCM tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.                                                                  |
| [FetchUserFcmToken](#fetchuserfcmtoken)           | Fetches a specific FCM token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.      |
| [DeleteUserFcmToken](#deleteuserfcmtoken)         | Deletes a specific user's FCM token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.            |
| [ListUserSlackTokens](#listuserslacktokens)       | Lists all Slack tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.                                                                |
| [FetchUserSlackToken](#fetchuserslacktoken)       | Fetches a specific Slack token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.    |
| [DeleteUserSlackToken](#deleteuserslacktoken)     | Deletes a specific user's Slack token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.          |
| [ListUserTeamsTokens](#listuserteamstokens)       | Lists all Teams tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.                                                                |
| [FetchUserTeamsToken](#fetchuserteamstoken)       | Fetches a specific Teams token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.    |
| [DeleteUserTeamsToken](#deleteuserteamstoken)     | Deletes a specific user's Teams token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.          |
| [ListUserWebPushTokens](#listuserwebpushtokens)   | Lists all Web Push tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.                                                             |
| [FetchUserWebPushToken](#fetchuserwebpushtoken)   | Fetches a specific Web Push token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata. |
| [DeleteUserWebPushToken](#deleteuserwebpushtoken) | Deletes a specific user's Web Push token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.       |

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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/channels"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)

channel := channels.CHANNEL_IN_APP

categoryDeliveryConfigChannels := channels.CategoryDeliveryConfigChannels{
  Channel: &channel,
  Delay: util.ToPointer(int64(123)),
  If_: util.ToPointer(util.Nullable[string]{ Value: "If_" }),
}

request := channels.CategoryDeliveryConfig{
  Channels: []channels.CategoryDeliveryConfigChannels{categoryDeliveryConfigChannels},
  Disabled: util.ToPointer(true),
  Key: util.ToPointer("Key"),
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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

  "github.com/magicbell/magicbell-go/pkg/project-client/channels"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


params := channels.ListUserInboxTokensRequestParams{

}

response, err := client.Channels.ListUserInboxTokens(context.Background(), "userId", params)
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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
client := client.NewClient(config)

response, err := client.Channels.FetchUserInboxToken(context.Background(), "userId", "tokenId")
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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
client := client.NewClient(config)

response, err := client.Channels.DeleteUserInboxToken(context.Background(), "userId", "tokenId")
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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

  "github.com/magicbell/magicbell-go/pkg/project-client/channels"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


params := channels.ListUserApnsTokensRequestParams{

}

response, err := client.Channels.ListUserApnsTokens(context.Background(), "userId", params)
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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
client := client.NewClient(config)

response, err := client.Channels.FetchUserApnsToken(context.Background(), "userId", "tokenId")
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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
client := client.NewClient(config)

response, err := client.Channels.DeleteUserApnsToken(context.Background(), "userId", "tokenId")
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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

  "github.com/magicbell/magicbell-go/pkg/project-client/channels"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


params := channels.ListUserExpoTokensRequestParams{

}

response, err := client.Channels.ListUserExpoTokens(context.Background(), "userId", params)
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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
client := client.NewClient(config)

response, err := client.Channels.FetchUserExpoToken(context.Background(), "userId", "tokenId")
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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
client := client.NewClient(config)

response, err := client.Channels.DeleteUserExpoToken(context.Background(), "userId", "tokenId")
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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

  "github.com/magicbell/magicbell-go/pkg/project-client/channels"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


params := channels.ListUserFcmTokensRequestParams{

}

response, err := client.Channels.ListUserFcmTokens(context.Background(), "userId", params)
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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
client := client.NewClient(config)

response, err := client.Channels.FetchUserFcmToken(context.Background(), "userId", "tokenId")
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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
client := client.NewClient(config)

response, err := client.Channels.DeleteUserFcmToken(context.Background(), "userId", "tokenId")
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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

  "github.com/magicbell/magicbell-go/pkg/project-client/channels"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


params := channels.ListUserSlackTokensRequestParams{

}

response, err := client.Channels.ListUserSlackTokens(context.Background(), "userId", params)
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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
client := client.NewClient(config)

response, err := client.Channels.FetchUserSlackToken(context.Background(), "userId", "tokenId")
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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
client := client.NewClient(config)

response, err := client.Channels.DeleteUserSlackToken(context.Background(), "userId", "tokenId")
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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

  "github.com/magicbell/magicbell-go/pkg/project-client/channels"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


params := channels.ListUserTeamsTokensRequestParams{

}

response, err := client.Channels.ListUserTeamsTokens(context.Background(), "userId", params)
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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
client := client.NewClient(config)

response, err := client.Channels.FetchUserTeamsToken(context.Background(), "userId", "tokenId")
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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
client := client.NewClient(config)

response, err := client.Channels.DeleteUserTeamsToken(context.Background(), "userId", "tokenId")
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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

  "github.com/magicbell/magicbell-go/pkg/project-client/channels"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


params := channels.ListUserWebPushTokensRequestParams{

}

response, err := client.Channels.ListUserWebPushTokens(context.Background(), "userId", params)
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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
client := client.NewClient(config)

response, err := client.Channels.FetchUserWebPushToken(context.Background(), "userId", "tokenId")
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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
client := client.NewClient(config)

response, err := client.Channels.DeleteUserWebPushToken(context.Background(), "userId", "tokenId")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

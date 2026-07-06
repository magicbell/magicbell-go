# ChannelsService

A list of all methods in the `ChannelsService` service. Click on the method name to view detailed information about that method.

| Methods                                                       | Description                                                                                                                                                                                                                                                               |
| :------------------------------------------------------------ | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| [ListInboxTokens](#listinboxtokens)                           | Lists all Inbox tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.                                                                                                          |
| [SaveInboxToken](#saveinboxtoken)                             | Saves the Inbox token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.              |
| [FetchInboxToken](#fetchinboxtoken)                           | Fetches details of a specific Inbox token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.                                                          |
| [DeleteInboxToken](#deleteinboxtoken)                         | Deletes one of the authenticated user's Inbox tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.                                    |
| [ListApnsTokens](#listapnstokens)                             | Lists all APNs tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.                                                                                                           |
| [SaveApnsToken](#saveapnstoken)                               | Saves the APNs token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.               |
| [FetchApnsToken](#fetchapnstoken)                             | Fetches details of a specific APNs token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.                                                           |
| [DeleteApnsToken](#deleteapnstoken)                           | Deletes one of the authenticated user's APNs tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.                                     |
| [ListExpoTokens](#listexpotokens)                             | Lists all Expo tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.                                                                                                           |
| [SaveExpoToken](#saveexpotoken)                               | Saves the Expo token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.               |
| [FetchExpoToken](#fetchexpotoken)                             | Fetches details of a specific Expo token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.                                                           |
| [DeleteExpoToken](#deleteexpotoken)                           | Deletes one of the authenticated user's Expo tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.                                     |
| [ListFcmTokens](#listfcmtokens)                               | Lists all FCM tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.                                                                                                            |
| [SaveFcmToken](#savefcmtoken)                                 | Saves the FCM token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.                |
| [FetchFcmToken](#fetchfcmtoken)                               | Fetches details of a specific FCM token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.                                                            |
| [DeleteFcmToken](#deletefcmtoken)                             | Deletes one of the authenticated user's FCM tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.                                      |
| [ListMagicbellSlackbotTokens](#listmagicbellslackbottokens)   | Lists all MagicBell SlackBot tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.                                                                                             |
| [SaveMagicbellSlackbotToken](#savemagicbellslackbottoken)     | Saves the MagicBell SlackBot token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel. |
| [FetchMagicbellSlackbotToken](#fetchmagicbellslackbottoken)   | Fetches details of a specific MagicBell SlackBot token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.                                             |
| [DeleteMagicbellSlackbotToken](#deletemagicbellslackbottoken) | Deletes one of the authenticated user's MagicBell SlackBot tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.                       |
| [ListSlackTokens](#listslacktokens)                           | Lists all Slack tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.                                                                                                          |
| [SaveSlackToken](#saveslacktoken)                             | Saves the Slack token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.              |
| [FetchSlackToken](#fetchslacktoken)                           | Fetches details of a specific Slack token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.                                                          |
| [DeleteSlackToken](#deleteslacktoken)                         | Deletes one of the authenticated user's Slack tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.                                    |
| [ListTeamsTokens](#listteamstokens)                           | Lists all Teams tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.                                                                                                          |
| [SaveTeamsToken](#saveteamstoken)                             | Saves the Teams token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.              |
| [FetchTeamsToken](#fetchteamstoken)                           | Fetches details of a specific Teams token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.                                                          |
| [DeleteTeamsToken](#deleteteamstoken)                         | Deletes one of the authenticated user's Teams tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.                                    |
| [FetchUserPreferences](#fetchuserpreferences)                 | Fetch a user's channel delivery preferences.                                                                                                                                                                                                                              |
| [SaveUserPreferences](#saveuserpreferences)                   | Save a user's channel preferences.                                                                                                                                                                                                                                        |
| [ListWebPushTokens](#listwebpushtokens)                       | Lists all Web Push tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.                                                                                                       |
| [SaveWebPushToken](#savewebpushtoken)                         | Saves the Web Push token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.           |
| [FetchWebPushToken](#fetchwebpushtoken)                       | Fetches details of a specific Web Push token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.                                                       |
| [DeleteWebPushToken](#deletewebpushtoken)                     | Deletes one of the authenticated user's Web Push tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.                                 |

## ListInboxTokens

Lists all Inbox tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/channels/in_app/inbox/tokens`

**Parameters**

| Name   | Type                         | Required | Description                   |
| :----- | :--------------------------- | :------- | :---------------------------- |
| ctx    | Context                      | ✅       | Default go language context   |
| params | ListInboxTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`InboxTokenResponseCollection`

**Example Usage Code Snippet**

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

## SaveInboxToken

Saves the Inbox token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.

- HTTP Method: `PUT`
- Endpoint: `/channels/in_app/inbox/tokens`

**Parameters**

| Name       | Type       | Required | Description                 |
| :--------- | :--------- | :------- | :-------------------------- |
| ctx        | Context    | ✅       | Default go language context |
| inboxToken | InboxToken | ✅       |                             |

**Return Type**

`InboxToken`

**Example Usage Code Snippet**

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


request := channels.InboxToken{
  ConnectionId: util.ToPointer(util.Nullable[string]{ Value: "connection_id" }),
  Token: util.ToPointer("ipsum nulla mag"),
}

response, err := client.Channels.SaveInboxToken(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchInboxToken

Fetches details of a specific Inbox token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.

- HTTP Method: `GET`
- Endpoint: `/channels/in_app/inbox/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`InboxTokenResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.FetchInboxToken(context.Background(), "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteInboxToken

Deletes one of the authenticated user's Inbox tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.

- HTTP Method: `DELETE`
- Endpoint: `/channels/in_app/inbox/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.DeleteInboxToken(context.Background(), "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListApnsTokens

Lists all APNs tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/channels/mobile_push/apns/tokens`

**Parameters**

| Name   | Type                        | Required | Description                   |
| :----- | :-------------------------- | :------- | :---------------------------- |
| ctx    | Context                     | ✅       | Default go language context   |
| params | ListApnsTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`ApnsTokenCollection`

**Example Usage Code Snippet**

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


params := channels.ListApnsTokensRequestParams{
  Limit: util.ToPointer(int64(7)),
  StartingAfter: util.ToPointer("starting_after"),
  EndingBefore: util.ToPointer("ending_before"),
}

response, err := client.Channels.ListApnsTokens(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveApnsToken

Saves the APNs token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.

- HTTP Method: `PUT`
- Endpoint: `/channels/mobile_push/apns/tokens`

**Parameters**

| Name             | Type             | Required | Description                 |
| :--------------- | :--------------- | :------- | :-------------------------- |
| ctx              | Context          | ✅       | Default go language context |
| apnsTokenPayload | ApnsTokenPayload | ✅       |                             |

**Return Type**

`ApnsTokenPayload`

**Example Usage Code Snippet**

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

apnsTokenPayloadInstallationId := channels.APNS_TOKEN_PAYLOAD_INSTALLATION_ID_DEVELOPMENT

request := channels.ApnsTokenPayload{
  AppId: util.ToPointer("app_id"),
  DeviceToken: util.ToPointer("sit culpa incid"),
  InstallationId: &apnsTokenPayloadInstallationId,
}

response, err := client.Channels.SaveApnsToken(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchApnsToken

Fetches details of a specific APNs token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.

- HTTP Method: `GET`
- Endpoint: `/channels/mobile_push/apns/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`ApnsToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.FetchApnsToken(context.Background(), "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteApnsToken

Deletes one of the authenticated user's APNs tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.

- HTTP Method: `DELETE`
- Endpoint: `/channels/mobile_push/apns/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.DeleteApnsToken(context.Background(), "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListExpoTokens

Lists all Expo tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/channels/mobile_push/expo/tokens`

**Parameters**

| Name   | Type                        | Required | Description                   |
| :----- | :-------------------------- | :------- | :---------------------------- |
| ctx    | Context                     | ✅       | Default go language context   |
| params | ListExpoTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`ExpoTokenCollection`

**Example Usage Code Snippet**

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


params := channels.ListExpoTokensRequestParams{
  Limit: util.ToPointer(int64(8)),
  StartingAfter: util.ToPointer("starting_after"),
  EndingBefore: util.ToPointer("ending_before"),
}

response, err := client.Channels.ListExpoTokens(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveExpoToken

Saves the Expo token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.

- HTTP Method: `PUT`
- Endpoint: `/channels/mobile_push/expo/tokens`

**Parameters**

| Name             | Type             | Required | Description                 |
| :--------------- | :--------------- | :------- | :-------------------------- |
| ctx              | Context          | ✅       | Default go language context |
| expoTokenPayload | ExpoTokenPayload | ✅       |                             |

**Return Type**

`ExpoTokenPayload`

**Example Usage Code Snippet**

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


request := channels.ExpoTokenPayload{
  DeviceToken: util.ToPointer("device_token"),
}

response, err := client.Channels.SaveExpoToken(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchExpoToken

Fetches details of a specific Expo token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.

- HTTP Method: `GET`
- Endpoint: `/channels/mobile_push/expo/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`ExpoToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.FetchExpoToken(context.Background(), "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteExpoToken

Deletes one of the authenticated user's Expo tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.

- HTTP Method: `DELETE`
- Endpoint: `/channels/mobile_push/expo/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.DeleteExpoToken(context.Background(), "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListFcmTokens

Lists all FCM tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/channels/mobile_push/fcm/tokens`

**Parameters**

| Name   | Type                       | Required | Description                   |
| :----- | :------------------------- | :------- | :---------------------------- |
| ctx    | Context                    | ✅       | Default go language context   |
| params | ListFcmTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`FcmTokenCollection`

**Example Usage Code Snippet**

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


params := channels.ListFcmTokensRequestParams{
  Limit: util.ToPointer(int64(6)),
  StartingAfter: util.ToPointer("starting_after"),
  EndingBefore: util.ToPointer("ending_before"),
}

response, err := client.Channels.ListFcmTokens(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveFcmToken

Saves the FCM token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.

- HTTP Method: `PUT`
- Endpoint: `/channels/mobile_push/fcm/tokens`

**Parameters**

| Name            | Type            | Required | Description                 |
| :-------------- | :-------------- | :------- | :-------------------------- |
| ctx             | Context         | ✅       | Default go language context |
| fcmTokenPayload | FcmTokenPayload | ✅       |                             |

**Return Type**

`FcmTokenPayload`

**Example Usage Code Snippet**

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


request := channels.FcmTokenPayload{
  DeviceToken: util.ToPointer("estminim adipis"),
}

response, err := client.Channels.SaveFcmToken(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchFcmToken

Fetches details of a specific FCM token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.

- HTTP Method: `GET`
- Endpoint: `/channels/mobile_push/fcm/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`FcmToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.FetchFcmToken(context.Background(), "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteFcmToken

Deletes one of the authenticated user's FCM tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.

- HTTP Method: `DELETE`
- Endpoint: `/channels/mobile_push/fcm/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.DeleteFcmToken(context.Background(), "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListMagicbellSlackbotTokens

Lists all MagicBell SlackBot tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/channels/slack/magicbell_slackbot/tokens`

**Parameters**

| Name   | Type                                     | Required | Description                   |
| :----- | :--------------------------------------- | :------- | :---------------------------- |
| ctx    | Context                                  | ✅       | Default go language context   |
| params | ListMagicbellSlackbotTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`SlackTokenCollection`

**Example Usage Code Snippet**

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


params := channels.ListMagicbellSlackbotTokensRequestParams{
  Limit: util.ToPointer(int64(7)),
  StartingAfter: util.ToPointer("starting_after"),
  EndingBefore: util.ToPointer("ending_before"),
}

response, err := client.Channels.ListMagicbellSlackbotTokens(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveMagicbellSlackbotToken

Saves the MagicBell SlackBot token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.

- HTTP Method: `PUT`
- Endpoint: `/channels/slack/magicbell_slackbot/tokens`

**Parameters**

| Name              | Type              | Required | Description                 |
| :---------------- | :---------------- | :------- | :-------------------------- |
| ctx               | Context           | ✅       | Default go language context |
| slackTokenPayload | SlackTokenPayload | ✅       |                             |

**Return Type**

`SlackTokenPayload`

**Example Usage Code Snippet**

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


slackTokenPayloadOauth := channels.SlackTokenPayloadOauth{
  ChannelId: util.ToPointer("channel_id"),
  InstallationId: util.ToPointer("installation_id"),
  Scope: util.ToPointer("scope"),
}


slackTokenPayloadWebhook := channels.SlackTokenPayloadWebhook{
  Url: util.ToPointer("url"),
}

request := channels.SlackTokenPayload{
  Oauth: &slackTokenPayloadOauth,
  Webhook: &slackTokenPayloadWebhook,
}

response, err := client.Channels.SaveMagicbellSlackbotToken(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchMagicbellSlackbotToken

Fetches details of a specific MagicBell SlackBot token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.

- HTTP Method: `GET`
- Endpoint: `/channels/slack/magicbell_slackbot/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`SlackToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.FetchMagicbellSlackbotToken(context.Background(), "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteMagicbellSlackbotToken

Deletes one of the authenticated user's MagicBell SlackBot tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.

- HTTP Method: `DELETE`
- Endpoint: `/channels/slack/magicbell_slackbot/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.DeleteMagicbellSlackbotToken(context.Background(), "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListSlackTokens

Lists all Slack tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/channels/slack/tokens`

**Parameters**

| Name   | Type                         | Required | Description                   |
| :----- | :--------------------------- | :------- | :---------------------------- |
| ctx    | Context                      | ✅       | Default go language context   |
| params | ListSlackTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`SlackTokenCollection`

**Example Usage Code Snippet**

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


params := channels.ListSlackTokensRequestParams{
  Limit: util.ToPointer(int64(2)),
  StartingAfter: util.ToPointer("starting_after"),
  EndingBefore: util.ToPointer("ending_before"),
}

response, err := client.Channels.ListSlackTokens(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveSlackToken

Saves the Slack token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.

- HTTP Method: `PUT`
- Endpoint: `/channels/slack/tokens`

**Parameters**

| Name              | Type              | Required | Description                 |
| :---------------- | :---------------- | :------- | :-------------------------- |
| ctx               | Context           | ✅       | Default go language context |
| slackTokenPayload | SlackTokenPayload | ✅       |                             |

**Return Type**

`SlackTokenPayload`

**Example Usage Code Snippet**

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


slackTokenPayloadOauth := channels.SlackTokenPayloadOauth{
  ChannelId: util.ToPointer("channel_id"),
  InstallationId: util.ToPointer("installation_id"),
  Scope: util.ToPointer("scope"),
}


slackTokenPayloadWebhook := channels.SlackTokenPayloadWebhook{
  Url: util.ToPointer("url"),
}

request := channels.SlackTokenPayload{
  Oauth: &slackTokenPayloadOauth,
  Webhook: &slackTokenPayloadWebhook,
}

response, err := client.Channels.SaveSlackToken(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchSlackToken

Fetches details of a specific Slack token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.

- HTTP Method: `GET`
- Endpoint: `/channels/slack/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`SlackToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.FetchSlackToken(context.Background(), "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteSlackToken

Deletes one of the authenticated user's Slack tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.

- HTTP Method: `DELETE`
- Endpoint: `/channels/slack/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.DeleteSlackToken(context.Background(), "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListTeamsTokens

Lists all Teams tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/channels/teams/tokens`

**Parameters**

| Name   | Type                         | Required | Description                   |
| :----- | :--------------------------- | :------- | :---------------------------- |
| ctx    | Context                      | ✅       | Default go language context   |
| params | ListTeamsTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`TeamsTokenCollection`

**Example Usage Code Snippet**

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


params := channels.ListTeamsTokensRequestParams{
  Limit: util.ToPointer(int64(4)),
  StartingAfter: util.ToPointer("starting_after"),
  EndingBefore: util.ToPointer("ending_before"),
}

response, err := client.Channels.ListTeamsTokens(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveTeamsToken

Saves the Teams token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.

- HTTP Method: `PUT`
- Endpoint: `/channels/teams/tokens`

**Parameters**

| Name              | Type              | Required | Description                 |
| :---------------- | :---------------- | :------- | :-------------------------- |
| ctx               | Context           | ✅       | Default go language context |
| teamsTokenPayload | TeamsTokenPayload | ✅       |                             |

**Return Type**

`TeamsTokenPayload`

**Example Usage Code Snippet**

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


teamsTokenPayloadWebhook := channels.TeamsTokenPayloadWebhook{
  Url: util.ToPointer("url"),
}

request := channels.TeamsTokenPayload{
  Webhook: &teamsTokenPayloadWebhook,
}

response, err := client.Channels.SaveTeamsToken(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchTeamsToken

Fetches details of a specific Teams token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.

- HTTP Method: `GET`
- Endpoint: `/channels/teams/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`TeamsToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.FetchTeamsToken(context.Background(), "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteTeamsToken

Deletes one of the authenticated user's Teams tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.

- HTTP Method: `DELETE`
- Endpoint: `/channels/teams/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.DeleteTeamsToken(context.Background(), "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchUserPreferences

Fetch a user's channel delivery preferences.

- HTTP Method: `GET`
- Endpoint: `/channels/user_preferences`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`UserPreferences`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.FetchUserPreferences(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveUserPreferences

Save a user's channel preferences.

- HTTP Method: `PUT`
- Endpoint: `/channels/user_preferences`

**Parameters**

| Name            | Type            | Required | Description                 |
| :-------------- | :-------------- | :------- | :-------------------------- |
| ctx             | Context         | ✅       | Default go language context |
| userPreferences | UserPreferences | ✅       |                             |

**Return Type**

`any`

**Example Usage Code Snippet**

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


channels := channels.Channels{
  Enabled: util.ToPointer(true),
  Name: util.ToPointer("name"),
}

categories := channels.Categories{
  Channels: []channels.Channels{channels},
  Key: util.ToPointer("key"),
  Label: util.ToPointer(util.Nullable[string]{ Value: "label" }),
}

request := channels.UserPreferences{
  Categories: []channels.Categories{categories},
}

response, err := client.Channels.SaveUserPreferences(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListWebPushTokens

Lists all Web Push tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.

- HTTP Method: `GET`
- Endpoint: `/channels/web_push/tokens`

**Parameters**

| Name   | Type                           | Required | Description                   |
| :----- | :----------------------------- | :------- | :---------------------------- |
| ctx    | Context                        | ✅       | Default go language context   |
| params | ListWebPushTokensRequestParams | ✅       | Additional request parameters |

**Return Type**

`WebPushTokenCollection`

**Example Usage Code Snippet**

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


params := channels.ListWebPushTokensRequestParams{
  Limit: util.ToPointer(int64(9)),
  StartingAfter: util.ToPointer("starting_after"),
  EndingBefore: util.ToPointer("ending_before"),
}

response, err := client.Channels.ListWebPushTokens(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveWebPushToken

Saves the Web Push token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.

- HTTP Method: `PUT`
- Endpoint: `/channels/web_push/tokens`

**Parameters**

| Name                | Type                       | Required | Description                 |
| :------------------ | :------------------------- | :------- | :-------------------------- |
| ctx                 | Context                    | ✅       | Default go language context |
| webPushTokenPayload | shared.WebPushTokenPayload | ✅       |                             |

**Return Type**

`shared.WebPushTokenPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/shared"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


webPushTokenPayloadKeys := shared.WebPushTokenPayloadKeys{
  Auth: util.ToPointer("auth"),
  P256dh: util.ToPointer("p256dh"),
}

request := shared.WebPushTokenPayload{
  Endpoint: util.ToPointer("endpoint"),
  Keys: &webPushTokenPayloadKeys,
}

response, err := client.Channels.SaveWebPushToken(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchWebPushToken

Fetches details of a specific Web Push token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.

- HTTP Method: `GET`
- Endpoint: `/channels/web_push/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`WebPushToken`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.FetchWebPushToken(context.Background(), "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteWebPushToken

Deletes one of the authenticated user's Web Push tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.

- HTTP Method: `DELETE`
- Endpoint: `/channels/web_push/tokens/{token_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| tokenId | string  | ✅       |                             |

**Return Type**

`DiscardResult`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Channels.DeleteWebPushToken(context.Background(), "token_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

<!-- This file was generated by liblab | https://liblab.com/ -->

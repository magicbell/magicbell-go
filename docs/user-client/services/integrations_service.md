# IntegrationsService

A list of all methods in the `IntegrationsService` service. Click on the method name to view detailed information about that method.

| Methods                                                                     | Description                                                                                                                                                                                                   |
| :-------------------------------------------------------------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| [SaveInboxInstallation](#saveinboxinstallation)                             | Creates a new installation of a Inbox integration for a user. This endpoint is used when an integration needs to be set up with user-specific credentials or configuration.                                   |
| [StartInboxInstallation](#startinboxinstallation)                           | Initiates the installation flow for an Inbox integration. This is the first step in a multi-step installation process where user authorization or external service configuration may be required.             |
| [SaveMagicbellSlackbotInstallation](#savemagicbellslackbotinstallation)     | Creates a new installation of a MagicBell SlackBot integration for a user. This endpoint is used when an integration needs to be set up with user-specific credentials or configuration.                      |
| [FinishMagicbellSlackbotInstallation](#finishmagicbellslackbotinstallation) | Completes the installation flow for the MagicBell SlackBot integration. This endpoint is typically called after the user has completed any required authorization steps with MagicBell SlackBot.              |
| [StartMagicbellSlackbotInstallation](#startmagicbellslackbotinstallation)   | Initiates the installation flow for a MagicBell SlackBot integration. This is the first step in a multi-step installation process where user authorization or external service configuration may be required. |
| [SaveSlackInstallation](#saveslackinstallation)                             | Creates a new installation of a Slack integration for a user. This endpoint is used when an integration needs to be set up with user-specific credentials or configuration.                                   |
| [FinishSlackInstallation](#finishslackinstallation)                         | Completes the installation flow for the Slack integration. This endpoint is typically called after the user has completed any required authorization steps with Slack.                                        |
| [StartSlackInstallation](#startslackinstallation)                           | Initiates the installation flow for a Slack integration. This is the first step in a multi-step installation process where user authorization or external service configuration may be required.              |
| [SaveWebPushInstallation](#savewebpushinstallation)                         | Creates a new installation of a Web Push integration for a user. This endpoint is used when an integration needs to be set up with user-specific credentials or configuration.                                |
| [StartWebPushInstallation](#startwebpushinstallation)                       | Initiates the installation flow for a Web Push integration. This is the first step in a multi-step installation process where user authorization or external service configuration may be required.           |

## SaveInboxInstallation

Creates a new installation of a Inbox integration for a user. This endpoint is used when an integration needs to be set up with user-specific credentials or configuration.

- HTTP Method: `PUT`
- Endpoint: `/integrations/inbox/installations`

**Parameters**

| Name               | Type               | Required | Description                 |
| :----------------- | :----------------- | :------- | :-------------------------- |
| ctx                | Context            | ✅       | Default go language context |
| inboxConfigPayload | InboxConfigPayload | ✅       |                             |

**Return Type**

`InboxConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


images := integrations.Images{
  EmptyInboxUrl: util.ToPointer("emptyInboxUrl"),
}


banner := integrations.Banner{
  BackgroundColor: util.ToPointer("backgroundColor"),
  BackgroundOpacity: util.ToPointer(float64(6.42)),
  FontSize: util.ToPointer("fontSize"),
  TextColor: util.ToPointer("textColor"),
}


dialog := integrations.Dialog{
  AccentColor: util.ToPointer("accentColor"),
  BackgroundColor: util.ToPointer("backgroundColor"),
  TextColor: util.ToPointer("textColor"),
}


footer := integrations.Footer{
  BackgroundColor: util.ToPointer("backgroundColor"),
  BorderRadius: util.ToPointer("borderRadius"),
  FontSize: util.ToPointer("fontSize"),
  TextColor: util.ToPointer("textColor"),
}


header := integrations.Header{
  BackgroundColor: util.ToPointer("backgroundColor"),
  BorderRadius: util.ToPointer("borderRadius"),
  FontFamily: util.ToPointer("fontFamily"),
  FontSize: util.ToPointer("fontSize"),
  TextColor: util.ToPointer("textColor"),
}


icon := integrations.Icon{
  BorderColor: util.ToPointer("borderColor"),
  Width: util.ToPointer("width"),
}


defaultHover := integrations.DefaultHover{
  BackgroundColor: util.ToPointer("backgroundColor"),
}


defaultState := integrations.DefaultState{
  Color: util.ToPointer("color"),
}

default_ := integrations.Default_{
  BackgroundColor: util.ToPointer("backgroundColor"),
  BorderRadius: util.ToPointer("borderRadius"),
  FontFamily: util.ToPointer("fontFamily"),
  FontSize: util.ToPointer("fontSize"),
  Hover: &defaultHover,
  Margin: util.ToPointer("margin"),
  State: &defaultState,
  TextColor: util.ToPointer("textColor"),
}


unreadHover := integrations.UnreadHover{
  BackgroundColor: util.ToPointer("backgroundColor"),
}


unreadState := integrations.UnreadState{
  Color: util.ToPointer("color"),
}

unread := integrations.Unread{
  BackgroundColor: util.ToPointer("backgroundColor"),
  Hover: &unreadHover,
  State: &unreadState,
  TextColor: util.ToPointer("textColor"),
}


unseenHover := integrations.UnseenHover{
  BackgroundColor: util.ToPointer("backgroundColor"),
}


unseenState := integrations.UnseenState{
  Color: util.ToPointer("color"),
}

unseen := integrations.Unseen{
  BackgroundColor: util.ToPointer("backgroundColor"),
  Hover: &unseenHover,
  State: &unseenState,
  TextColor: util.ToPointer("textColor"),
}

themeNotification := integrations.ThemeNotification{
  Default_: &default_,
  Unread: &unread,
  Unseen: &unseen,
}


unseenBadge := integrations.UnseenBadge{
  BackgroundColor: util.ToPointer("backgroundColor"),
}

theme := integrations.Theme{
  Banner: &banner,
  Dialog: &dialog,
  Footer: &footer,
  Header: &header,
  Icon: &icon,
  Notification: &themeNotification,
  UnseenBadge: &unseenBadge,
}

request := integrations.InboxConfigPayload{
  Images: &images,
  Locale: util.ToPointer(util.Nullable[string]{ Value: "locale" }),
  Theme: &theme,
}

response, err := client.Integrations.SaveInboxInstallation(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## StartInboxInstallation

Initiates the installation flow for an Inbox integration. This is the first step in a multi-step installation process where user authorization or external service configuration may be required.

- HTTP Method: `POST`
- Endpoint: `/integrations/inbox/installations/start`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`InboxConfigPayload`

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

response, err := client.Integrations.StartInboxInstallation(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveMagicbellSlackbotInstallation

Creates a new installation of a MagicBell SlackBot integration for a user. This endpoint is used when an integration needs to be set up with user-specific credentials or configuration.

- HTTP Method: `PUT`
- Endpoint: `/integrations/magicbell_slackbot/installations`

**Parameters**

| Name              | Type              | Required | Description                 |
| :---------------- | :---------------- | :------- | :-------------------------- |
| ctx               | Context           | ✅       | Default go language context |
| slackInstallation | SlackInstallation | ✅       |                             |

**Return Type**

`SlackInstallation`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


authedUser := integrations.AuthedUser{
  AccessToken: util.ToPointer("access_token"),
  ExpiresIn: util.ToPointer(int64(1)),
  Id: util.ToPointer("id"),
  RefreshToken: util.ToPointer("refresh_token"),
  Scope: util.ToPointer("scope"),
  TokenType: util.ToPointer("token_type"),
}


enterprise := integrations.Enterprise{
  Id: util.ToPointer("id"),
  Name: util.ToPointer("name"),
}


incomingWebhook := integrations.IncomingWebhook{
  Channel: util.ToPointer("channel"),
  ConfigurationUrl: util.ToPointer("configuration_url"),
  Url: util.ToPointer("url"),
}


team := integrations.Team{
  Id: util.ToPointer("id"),
  Name: util.ToPointer("name"),
}

request := integrations.SlackInstallation{
  AccessToken: util.ToPointer("access_token"),
  AppId: util.ToPointer("app_id"),
  AuthedUser: &authedUser,
  BotUserId: util.ToPointer("bot_user_id"),
  Enterprise: &enterprise,
  ExpiresIn: util.ToPointer(int64(2)),
  Id: util.ToPointer("318-g~J]11"),
  IncomingWebhook: &incomingWebhook,
  IsEnterpriseInstall: util.ToPointer(true),
  RefreshToken: util.ToPointer("refresh_token"),
  Scope: util.ToPointer("scope"),
  Team: &team,
  TokenType: util.ToPointer("token_type"),
}

response, err := client.Integrations.SaveMagicbellSlackbotInstallation(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FinishMagicbellSlackbotInstallation

Completes the installation flow for the MagicBell SlackBot integration. This endpoint is typically called after the user has completed any required authorization steps with MagicBell SlackBot.

- HTTP Method: `POST`
- Endpoint: `/integrations/magicbell_slackbot/installations/finish`

**Parameters**

| Name                       | Type                       | Required | Description                 |
| :------------------------- | :------------------------- | :------- | :-------------------------- |
| ctx                        | Context                    | ✅       | Default go language context |
| slackFinishInstallResponse | SlackFinishInstallResponse | ✅       |                             |

**Return Type**

`SlackInstallation`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


request := integrations.SlackFinishInstallResponse{
  AppId: util.ToPointer("app_id"),
  Code: util.ToPointer("code"),
  RedirectUrl: util.ToPointer("redirect_url"),
}

response, err := client.Integrations.FinishMagicbellSlackbotInstallation(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## StartMagicbellSlackbotInstallation

Initiates the installation flow for a MagicBell SlackBot integration. This is the first step in a multi-step installation process where user authorization or external service configuration may be required.

- HTTP Method: `POST`
- Endpoint: `/integrations/magicbell_slackbot/installations/start`

**Parameters**

| Name              | Type              | Required | Description                 |
| :---------------- | :---------------- | :------- | :-------------------------- |
| ctx               | Context           | ✅       | Default go language context |
| slackStartInstall | SlackStartInstall | ✅       |                             |

**Return Type**

`SlackStartInstallResponseContent`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


request := integrations.SlackStartInstall{
  AppId: util.ToPointer("app_id"),
  AuthUrl: util.ToPointer("auth_url"),
  ExtraScopes: []string{},
  RedirectUrl: util.ToPointer("redirect_url"),
}

response, err := client.Integrations.StartMagicbellSlackbotInstallation(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveSlackInstallation

Creates a new installation of a Slack integration for a user. This endpoint is used when an integration needs to be set up with user-specific credentials or configuration.

- HTTP Method: `PUT`
- Endpoint: `/integrations/slack/installations`

**Parameters**

| Name              | Type              | Required | Description                 |
| :---------------- | :---------------- | :------- | :-------------------------- |
| ctx               | Context           | ✅       | Default go language context |
| slackInstallation | SlackInstallation | ✅       |                             |

**Return Type**

`SlackInstallation`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


authedUser := integrations.AuthedUser{
  AccessToken: util.ToPointer("access_token"),
  ExpiresIn: util.ToPointer(int64(1)),
  Id: util.ToPointer("id"),
  RefreshToken: util.ToPointer("refresh_token"),
  Scope: util.ToPointer("scope"),
  TokenType: util.ToPointer("token_type"),
}


enterprise := integrations.Enterprise{
  Id: util.ToPointer("id"),
  Name: util.ToPointer("name"),
}


incomingWebhook := integrations.IncomingWebhook{
  Channel: util.ToPointer("channel"),
  ConfigurationUrl: util.ToPointer("configuration_url"),
  Url: util.ToPointer("url"),
}


team := integrations.Team{
  Id: util.ToPointer("id"),
  Name: util.ToPointer("name"),
}

request := integrations.SlackInstallation{
  AccessToken: util.ToPointer("access_token"),
  AppId: util.ToPointer("app_id"),
  AuthedUser: &authedUser,
  BotUserId: util.ToPointer("bot_user_id"),
  Enterprise: &enterprise,
  ExpiresIn: util.ToPointer(int64(2)),
  Id: util.ToPointer("318-g~J]11"),
  IncomingWebhook: &incomingWebhook,
  IsEnterpriseInstall: util.ToPointer(true),
  RefreshToken: util.ToPointer("refresh_token"),
  Scope: util.ToPointer("scope"),
  Team: &team,
  TokenType: util.ToPointer("token_type"),
}

response, err := client.Integrations.SaveSlackInstallation(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FinishSlackInstallation

Completes the installation flow for the Slack integration. This endpoint is typically called after the user has completed any required authorization steps with Slack.

- HTTP Method: `POST`
- Endpoint: `/integrations/slack/installations/finish`

**Parameters**

| Name                       | Type                       | Required | Description                 |
| :------------------------- | :------------------------- | :------- | :-------------------------- |
| ctx                        | Context                    | ✅       | Default go language context |
| slackFinishInstallResponse | SlackFinishInstallResponse | ✅       |                             |

**Return Type**

`SlackInstallation`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


request := integrations.SlackFinishInstallResponse{
  AppId: util.ToPointer("app_id"),
  Code: util.ToPointer("code"),
  RedirectUrl: util.ToPointer("redirect_url"),
}

response, err := client.Integrations.FinishSlackInstallation(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## StartSlackInstallation

Initiates the installation flow for a Slack integration. This is the first step in a multi-step installation process where user authorization or external service configuration may be required.

- HTTP Method: `POST`
- Endpoint: `/integrations/slack/installations/start`

**Parameters**

| Name              | Type              | Required | Description                 |
| :---------------- | :---------------- | :------- | :-------------------------- |
| ctx               | Context           | ✅       | Default go language context |
| slackStartInstall | SlackStartInstall | ✅       |                             |

**Return Type**

`SlackStartInstallResponseContent`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


request := integrations.SlackStartInstall{
  AppId: util.ToPointer("app_id"),
  AuthUrl: util.ToPointer("auth_url"),
  ExtraScopes: []string{},
  RedirectUrl: util.ToPointer("redirect_url"),
}

response, err := client.Integrations.StartSlackInstallation(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveWebPushInstallation

Creates a new installation of a Web Push integration for a user. This endpoint is used when an integration needs to be set up with user-specific credentials or configuration.

- HTTP Method: `PUT`
- Endpoint: `/integrations/web_push/installations`

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

response, err := client.Integrations.SaveWebPushInstallation(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## StartWebPushInstallation

Initiates the installation flow for a Web Push integration. This is the first step in a multi-step installation process where user authorization or external service configuration may be required.

- HTTP Method: `POST`
- Endpoint: `/integrations/web_push/installations/start`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`WebPushStartInstallationResponse`

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

response, err := client.Integrations.StartWebPushInstallation(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

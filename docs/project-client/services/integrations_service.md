# IntegrationsService

A list of all methods in the `IntegrationsService` service. Click on the method name to view detailed information about that method.

| Methods                                                   | Description                                                                                                                                                             |
| :-------------------------------------------------------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [ListIntegrations](#listintegrations)                     | Lists all available and configured integrations for the project. Returns a summary of each integration including its type, status, and basic configuration information. |
| [ListApnsIntegrations](#listapnsintegrations)             | Retrieves the current APNs integration configurations for a specific integration type in the project. Returns configuration details and status information.             |
| [SaveApnsIntegration](#saveapnsintegration)               | Updates or creates the APNs integration for the project.                                                                                                                |
| [DeleteApnsIntegration](#deleteapnsintegration)           | Deletes the APNs integration configuration from the project. This will disable the integration's functionality within the project.                                      |
| [ListExpoIntegrations](#listexpointegrations)             | Retrieves the current Expo integration configurations for a specific integration type in the project. Returns configuration details and status information.             |
| [SaveExpoIntegration](#saveexpointegration)               | Updates or creates the Expo integration for the project.                                                                                                                |
| [DeleteExpoIntegration](#deleteexpointegration)           | Deletes the Expo integration configuration from the project. This will disable the integration's functionality within the project.                                      |
| [ListFcmIntegrations](#listfcmintegrations)               | Retrieves the current FCM integration configurations for a specific integration type in the project. Returns configuration details and status information.              |
| [SaveFcmIntegration](#savefcmintegration)                 | Updates or creates the FCM integration for the project.                                                                                                                 |
| [DeleteFcmIntegration](#deletefcmintegration)             | Deletes the FCM integration configuration from the project. This will disable the integration's functionality within the project.                                       |
| [ListInboxIntegrations](#listinboxintegrations)           | Retrieves the current Inbox integration configurations for a specific integration type in the project. Returns configuration details and status information.            |
| [SaveInboxIntegration](#saveinboxintegration)             | Updates or creates the Inbox integration for the project.                                                                                                               |
| [DeleteInboxIntegration](#deleteinboxintegration)         | Deletes the Inbox integration configuration from the project. This will disable the integration's functionality within the project.                                     |
| [ListMailgunIntegrations](#listmailgunintegrations)       | Retrieves the current Mailgun integration configurations for a specific integration type in the project. Returns configuration details and status information.          |
| [SaveMailgunIntegration](#savemailgunintegration)         | Updates or creates the Mailgun integration for the project.                                                                                                             |
| [DeleteMailgunIntegration](#deletemailgunintegration)     | Deletes the Mailgun integration configuration from the project. This will disable the integration's functionality within the project.                                   |
| [ListPingEmailIntegrations](#listpingemailintegrations)   | Retrieves the current Ping Email integration configurations for a specific integration type in the project. Returns configuration details and status information.       |
| [SavePingEmailIntegration](#savepingemailintegration)     | Updates or creates the Ping Email integration for the project.                                                                                                          |
| [DeletePingEmailIntegration](#deletepingemailintegration) | Deletes the Ping Email integration configuration from the project. This will disable the integration's functionality within the project.                                |
| [ListSendgridIntegrations](#listsendgridintegrations)     | Retrieves the current SendGrid integration configurations for a specific integration type in the project. Returns configuration details and status information.         |
| [SaveSendgridIntegration](#savesendgridintegration)       | Updates or creates the SendGrid integration for the project.                                                                                                            |
| [DeleteSendgridIntegration](#deletesendgridintegration)   | Deletes the SendGrid integration configuration from the project. This will disable the integration's functionality within the project.                                  |
| [ListSesIntegrations](#listsesintegrations)               | Retrieves the current Amazon SES integration configurations for a specific integration type in the project. Returns configuration details and status information.       |
| [SaveSesIntegration](#savesesintegration)                 | Updates or creates the Amazon SES integration for the project.                                                                                                          |
| [DeleteSesIntegration](#deletesesintegration)             | Deletes the Amazon SES integration configuration from the project. This will disable the integration's functionality within the project.                                |
| [ListSlackIntegrations](#listslackintegrations)           | Retrieves the current Slack integration configurations for a specific integration type in the project. Returns configuration details and status information.            |
| [SaveSlackIntegration](#saveslackintegration)             | Updates or creates the Slack integration for the project.                                                                                                               |
| [DeleteSlackIntegration](#deleteslackintegration)         | Deletes the Slack integration configuration from the project. This will disable the integration's functionality within the project.                                     |
| [ListTwilioIntegrations](#listtwiliointegrations)         | Retrieves the current Twilio integration configurations for a specific integration type in the project. Returns configuration details and status information.           |
| [SaveTwilioIntegration](#savetwiliointegration)           | Updates or creates the Twilio integration for the project.                                                                                                              |
| [DeleteTwilioIntegration](#deletetwiliointegration)       | Deletes the Twilio integration configuration from the project. This will disable the integration's functionality within the project.                                    |
| [ListWebPushIntegrations](#listwebpushintegrations)       | Retrieves the current Web Push integration configurations for a specific integration type in the project. Returns configuration details and status information.         |
| [SaveWebPushIntegration](#savewebpushintegration)         | Updates or creates the Web Push integration for the project.                                                                                                            |
| [DeleteWebPushIntegration](#deletewebpushintegration)     | Deletes the Web Push integration configuration from the project. This will disable the integration's functionality within the project.                                  |

## ListIntegrations

Lists all available and configured integrations for the project. Returns a summary of each integration including its type, status, and basic configuration information.

- HTTP Method: `GET`
- Endpoint: `/integrations`

**Parameters**

| Name   | Type                          | Required | Description                   |
| :----- | :---------------------------- | :------- | :---------------------------- |
| ctx    | Context                       | ✅       | Default go language context   |
| params | ListIntegrationsRequestParams | ✅       | Additional request parameters |

**Return Type**

`IntegrationConfigCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


params := integrations.ListIntegrationsRequestParams{

}

response, err := client.Integrations.ListIntegrations(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListApnsIntegrations

Retrieves the current APNs integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/apns`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`ApnsConfigCollection`

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

response, err := client.Integrations.ListApnsIntegrations(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveApnsIntegration

Updates or creates the APNs integration for the project.

- HTTP Method: `PUT`
- Endpoint: `/integrations/apns`

**Parameters**

| Name              | Type              | Required | Description                 |
| :---------------- | :---------------- | :------- | :-------------------------- |
| ctx               | Context           | ✅       | Default go language context |
| apnsConfigPayload | ApnsConfigPayload | ✅       |                             |

**Return Type**

`ApnsConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)

badge := integrations.BADGE_UNREAD

payloadVersion := integrations.PAYLOAD_VERSION_1

request := integrations.ApnsConfigPayload{
  AppId: util.ToPointer("AppId"),
  Badge: &badge,
  Certificate: util.ToPointer("Certificate"),
  KeyId: util.ToPointer("KeyId"),
  PayloadVersion: &payloadVersion,
  TeamId: util.ToPointer("TeamId"),
}

response, err := client.Integrations.SaveApnsIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteApnsIntegration

Deletes the APNs integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/apns`

**Parameters**

| Name   | Type                               | Required | Description                   |
| :----- | :--------------------------------- | :------- | :---------------------------- |
| ctx    | Context                            | ✅       | Default go language context   |
| params | DeleteApnsIntegrationRequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


params := integrations.DeleteApnsIntegrationRequestParams{

}

response, err := client.Integrations.DeleteApnsIntegration(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListExpoIntegrations

Retrieves the current Expo integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/expo`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`ExpoConfigCollection`

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

response, err := client.Integrations.ListExpoIntegrations(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveExpoIntegration

Updates or creates the Expo integration for the project.

- HTTP Method: `PUT`
- Endpoint: `/integrations/expo`

**Parameters**

| Name              | Type              | Required | Description                 |
| :---------------- | :---------------- | :------- | :-------------------------- |
| ctx               | Context           | ✅       | Default go language context |
| expoConfigPayload | ExpoConfigPayload | ✅       |                             |

**Return Type**

`ExpoConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


request := integrations.ExpoConfigPayload{
  AccessToken: util.ToPointer("AccessToken"),
}

response, err := client.Integrations.SaveExpoIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteExpoIntegration

Deletes the Expo integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/expo`

**Parameters**

| Name   | Type                               | Required | Description                   |
| :----- | :--------------------------------- | :------- | :---------------------------- |
| ctx    | Context                            | ✅       | Default go language context   |
| params | DeleteExpoIntegrationRequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


params := integrations.DeleteExpoIntegrationRequestParams{

}

response, err := client.Integrations.DeleteExpoIntegration(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListFcmIntegrations

Retrieves the current FCM integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/fcm`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`FcmConfigCollection`

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

response, err := client.Integrations.ListFcmIntegrations(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveFcmIntegration

Updates or creates the FCM integration for the project.

- HTTP Method: `PUT`
- Endpoint: `/integrations/fcm`

**Parameters**

| Name             | Type             | Required | Description                 |
| :--------------- | :--------------- | :------- | :-------------------------- |
| ctx              | Context          | ✅       | Default go language context |
| fcmConfigPayload | FcmConfigPayload | ✅       |                             |

**Return Type**

`FcmConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)

type_ := integrations.TYPE_SERVICE_ACCOUNT

request := integrations.FcmConfigPayload{
  AuthProviderX509CertUrl: util.ToPointer("AuthProviderX509CertUrl"),
  AuthUri: util.ToPointer("AuthUri"),
  ClientEmail: util.ToPointer("ClientEmail"),
  ClientId: util.ToPointer("ClientId"),
  ClientX509CertUrl: util.ToPointer("ClientX509CertUrl"),
  PrivateKey: util.ToPointer("PrivateKey"),
  PrivateKeyId: util.ToPointer("PrivateKeyId"),
  ProjectId: util.ToPointer("ProjectId"),
  TokenUri: util.ToPointer("TokenUri"),
  Type_: &type_,
  UniverseDomain: util.ToPointer("UniverseDomain"),
}

response, err := client.Integrations.SaveFcmIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteFcmIntegration

Deletes the FCM integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/fcm`

**Parameters**

| Name   | Type                              | Required | Description                   |
| :----- | :-------------------------------- | :------- | :---------------------------- |
| ctx    | Context                           | ✅       | Default go language context   |
| params | DeleteFcmIntegrationRequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


params := integrations.DeleteFcmIntegrationRequestParams{

}

response, err := client.Integrations.DeleteFcmIntegration(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListInboxIntegrations

Retrieves the current Inbox integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/inbox`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`InboxConfigCollection`

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

response, err := client.Integrations.ListInboxIntegrations(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveInboxIntegration

Updates or creates the Inbox integration for the project.

- HTTP Method: `PUT`
- Endpoint: `/integrations/inbox`

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
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


images := integrations.Images{
  EmptyInboxUrl: util.ToPointer("EmptyInboxUrl"),
}


banner := integrations.Banner{
  BackgroundColor: util.ToPointer("BackgroundColor"),
  BackgroundOpacity: util.ToPointer(float64(123)),
  FontSize: util.ToPointer("FontSize"),
  TextColor: util.ToPointer("TextColor"),
}


dialog := integrations.Dialog{
  AccentColor: util.ToPointer("AccentColor"),
  BackgroundColor: util.ToPointer("BackgroundColor"),
  TextColor: util.ToPointer("TextColor"),
}


footer := integrations.Footer{
  BackgroundColor: util.ToPointer("BackgroundColor"),
  BorderRadius: util.ToPointer("BorderRadius"),
  FontSize: util.ToPointer("FontSize"),
  TextColor: util.ToPointer("TextColor"),
}


header := integrations.Header{
  BackgroundColor: util.ToPointer("BackgroundColor"),
  BorderRadius: util.ToPointer("BorderRadius"),
  FontFamily: util.ToPointer("FontFamily"),
  FontSize: util.ToPointer("FontSize"),
  TextColor: util.ToPointer("TextColor"),
}


icon := integrations.Icon{
  BorderColor: util.ToPointer("BorderColor"),
  Width: util.ToPointer("Width"),
}


defaultHover := integrations.DefaultHover{
  BackgroundColor: util.ToPointer("BackgroundColor"),
}


defaultState := integrations.DefaultState{
  Color: util.ToPointer("Color"),
}

default_ := integrations.Default_{
  BackgroundColor: util.ToPointer("BackgroundColor"),
  BorderRadius: util.ToPointer("BorderRadius"),
  FontFamily: util.ToPointer("FontFamily"),
  FontSize: util.ToPointer("FontSize"),
  Hover: &defaultHover,
  Margin: util.ToPointer("Margin"),
  State: &defaultState,
  TextColor: util.ToPointer("TextColor"),
}


unreadHover := integrations.UnreadHover{
  BackgroundColor: util.ToPointer("BackgroundColor"),
}


unreadState := integrations.UnreadState{
  Color: util.ToPointer("Color"),
}

unread := integrations.Unread{
  BackgroundColor: util.ToPointer("BackgroundColor"),
  Hover: &unreadHover,
  State: &unreadState,
  TextColor: util.ToPointer("TextColor"),
}


unseenHover := integrations.UnseenHover{
  BackgroundColor: util.ToPointer("BackgroundColor"),
}


unseenState := integrations.UnseenState{
  Color: util.ToPointer("Color"),
}

unseen := integrations.Unseen{
  BackgroundColor: util.ToPointer("BackgroundColor"),
  Hover: &unseenHover,
  State: &unseenState,
  TextColor: util.ToPointer("TextColor"),
}

notification := integrations.Notification{
  Default_: &default_,
  Unread: &unread,
  Unseen: &unseen,
}


unseenBadge := integrations.UnseenBadge{
  BackgroundColor: util.ToPointer("BackgroundColor"),
}

theme := integrations.Theme{
  Banner: &banner,
  Dialog: &dialog,
  Footer: &footer,
  Header: &header,
  Icon: &icon,
  Notification: &notification,
  UnseenBadge: &unseenBadge,
}

request := integrations.InboxConfigPayload{
  Images: &images,
  Locale: util.ToPointer(util.Nullable[string]{ Value: "Locale" }),
  Theme: &theme,
}

response, err := client.Integrations.SaveInboxIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteInboxIntegration

Deletes the Inbox integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/inbox`

**Parameters**

| Name   | Type                                | Required | Description                   |
| :----- | :---------------------------------- | :------- | :---------------------------- |
| ctx    | Context                             | ✅       | Default go language context   |
| params | DeleteInboxIntegrationRequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


params := integrations.DeleteInboxIntegrationRequestParams{

}

response, err := client.Integrations.DeleteInboxIntegration(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListMailgunIntegrations

Retrieves the current Mailgun integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/mailgun`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`MailgunConfigCollection`

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

response, err := client.Integrations.ListMailgunIntegrations(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveMailgunIntegration

Updates or creates the Mailgun integration for the project.

- HTTP Method: `PUT`
- Endpoint: `/integrations/mailgun`

**Parameters**

| Name                 | Type                 | Required | Description                 |
| :------------------- | :------------------- | :------- | :-------------------------- |
| ctx                  | Context              | ✅       | Default go language context |
| mailgunConfigPayload | MailgunConfigPayload | ✅       |                             |

**Return Type**

`MailgunConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


mailgunConfigPayloadFrom := integrations.MailgunConfigPayloadFrom{
  Email: util.ToPointer("Email"),
  Name: util.ToPointer(util.Nullable[string]{ Value: "Name" }),
}

region := integrations.REGION_US

request := integrations.MailgunConfigPayload{
  ApiKey: util.ToPointer("ApiKey"),
  Domain: util.ToPointer("Domain"),
  From: &mailgunConfigPayloadFrom,
  Region: &region,
}

response, err := client.Integrations.SaveMailgunIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteMailgunIntegration

Deletes the Mailgun integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/mailgun`

**Parameters**

| Name   | Type                                  | Required | Description                   |
| :----- | :------------------------------------ | :------- | :---------------------------- |
| ctx    | Context                               | ✅       | Default go language context   |
| params | DeleteMailgunIntegrationRequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


params := integrations.DeleteMailgunIntegrationRequestParams{

}

response, err := client.Integrations.DeleteMailgunIntegration(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListPingEmailIntegrations

Retrieves the current Ping Email integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/ping_email`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`PingConfigCollection`

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

response, err := client.Integrations.ListPingEmailIntegrations(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SavePingEmailIntegration

Updates or creates the Ping Email integration for the project.

- HTTP Method: `PUT`
- Endpoint: `/integrations/ping_email`

**Parameters**

| Name              | Type              | Required | Description                 |
| :---------------- | :---------------- | :------- | :-------------------------- |
| ctx               | Context           | ✅       | Default go language context |
| pingConfigPayload | PingConfigPayload | ✅       |                             |

**Return Type**

`PingConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


request := integrations.PingConfigPayload{
  Url: util.ToPointer("Url"),
}

response, err := client.Integrations.SavePingEmailIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeletePingEmailIntegration

Deletes the Ping Email integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/ping_email`

**Parameters**

| Name   | Type                                    | Required | Description                   |
| :----- | :-------------------------------------- | :------- | :---------------------------- |
| ctx    | Context                                 | ✅       | Default go language context   |
| params | DeletePingEmailIntegrationRequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


params := integrations.DeletePingEmailIntegrationRequestParams{

}

response, err := client.Integrations.DeletePingEmailIntegration(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListSendgridIntegrations

Retrieves the current SendGrid integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/sendgrid`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`SendgridConfigCollection`

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

response, err := client.Integrations.ListSendgridIntegrations(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveSendgridIntegration

Updates or creates the SendGrid integration for the project.

- HTTP Method: `PUT`
- Endpoint: `/integrations/sendgrid`

**Parameters**

| Name                  | Type                  | Required | Description                 |
| :-------------------- | :-------------------- | :------- | :-------------------------- |
| ctx                   | Context               | ✅       | Default go language context |
| sendgridConfigPayload | SendgridConfigPayload | ✅       |                             |

**Return Type**

`SendgridConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


sendgridConfigPayloadFrom := integrations.SendgridConfigPayloadFrom{
  Email: util.ToPointer("Email"),
  Name: util.ToPointer(util.Nullable[string]{ Value: "Name" }),
}


replyTo := integrations.ReplyTo{
  Email: util.ToPointer("Email"),
  Name: util.ToPointer(util.Nullable[string]{ Value: "Name" }),
}

request := integrations.SendgridConfigPayload{
  ApiKey: util.ToPointer("ApiKey"),
  From: &sendgridConfigPayloadFrom,
  ReplyTo: &replyTo,
}

response, err := client.Integrations.SaveSendgridIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteSendgridIntegration

Deletes the SendGrid integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/sendgrid`

**Parameters**

| Name   | Type                                   | Required | Description                   |
| :----- | :------------------------------------- | :------- | :---------------------------- |
| ctx    | Context                                | ✅       | Default go language context   |
| params | DeleteSendgridIntegrationRequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


params := integrations.DeleteSendgridIntegrationRequestParams{

}

response, err := client.Integrations.DeleteSendgridIntegration(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListSesIntegrations

Retrieves the current Amazon SES integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/ses`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`SesConfigCollection`

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

response, err := client.Integrations.ListSesIntegrations(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveSesIntegration

Updates or creates the Amazon SES integration for the project.

- HTTP Method: `PUT`
- Endpoint: `/integrations/ses`

**Parameters**

| Name             | Type             | Required | Description                 |
| :--------------- | :--------------- | :------- | :-------------------------- |
| ctx              | Context          | ✅       | Default go language context |
| sesConfigPayload | SesConfigPayload | ✅       |                             |

**Return Type**

`SesConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


sesConfigPayloadFrom := integrations.SesConfigPayloadFrom{
  Email: util.ToPointer("Email"),
  Name: util.ToPointer(util.Nullable[string]{ Value: "Name" }),
}

request := integrations.SesConfigPayload{
  From: &sesConfigPayloadFrom,
  KeyId: util.ToPointer("KeyId"),
  Region: util.ToPointer("Region"),
  SecretKey: util.ToPointer("SecretKey"),
}

response, err := client.Integrations.SaveSesIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteSesIntegration

Deletes the Amazon SES integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/ses`

**Parameters**

| Name   | Type                              | Required | Description                   |
| :----- | :-------------------------------- | :------- | :---------------------------- |
| ctx    | Context                           | ✅       | Default go language context   |
| params | DeleteSesIntegrationRequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


params := integrations.DeleteSesIntegrationRequestParams{

}

response, err := client.Integrations.DeleteSesIntegration(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListSlackIntegrations

Retrieves the current Slack integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/slack`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`SlackConfigCollection`

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

response, err := client.Integrations.ListSlackIntegrations(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveSlackIntegration

Updates or creates the Slack integration for the project.

- HTTP Method: `PUT`
- Endpoint: `/integrations/slack`

**Parameters**

| Name               | Type               | Required | Description                 |
| :----------------- | :----------------- | :------- | :-------------------------- |
| ctx                | Context            | ✅       | Default go language context |
| slackConfigPayload | SlackConfigPayload | ✅       |                             |

**Return Type**

`SlackConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


request := integrations.SlackConfigPayload{
  AppId: util.ToPointer("AppId"),
  ClientId: util.ToPointer("ClientId"),
  ClientSecret: util.ToPointer("ClientSecret"),
  SigningSecret: util.ToPointer("SigningSecret"),
}

response, err := client.Integrations.SaveSlackIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteSlackIntegration

Deletes the Slack integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/slack`

**Parameters**

| Name   | Type                                | Required | Description                   |
| :----- | :---------------------------------- | :------- | :---------------------------- |
| ctx    | Context                             | ✅       | Default go language context   |
| params | DeleteSlackIntegrationRequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


params := integrations.DeleteSlackIntegrationRequestParams{

}

response, err := client.Integrations.DeleteSlackIntegration(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListTwilioIntegrations

Retrieves the current Twilio integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/twilio`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`TwilioConfigCollection`

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

response, err := client.Integrations.ListTwilioIntegrations(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveTwilioIntegration

Updates or creates the Twilio integration for the project.

- HTTP Method: `PUT`
- Endpoint: `/integrations/twilio`

**Parameters**

| Name                | Type                | Required | Description                 |
| :------------------ | :------------------ | :------- | :-------------------------- |
| ctx                 | Context             | ✅       | Default go language context |
| twilioConfigPayload | TwilioConfigPayload | ✅       |                             |

**Return Type**

`TwilioConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


request := integrations.TwilioConfigPayload{
  AccountSid: util.ToPointer("AccountSid"),
  ApiKey: util.ToPointer("ApiKey"),
  ApiSecret: util.ToPointer("ApiSecret"),
  From: util.ToPointer("From"),
}

response, err := client.Integrations.SaveTwilioIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteTwilioIntegration

Deletes the Twilio integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/twilio`

**Parameters**

| Name   | Type                                 | Required | Description                   |
| :----- | :----------------------------------- | :------- | :---------------------------- |
| ctx    | Context                              | ✅       | Default go language context   |
| params | DeleteTwilioIntegrationRequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


params := integrations.DeleteTwilioIntegrationRequestParams{

}

response, err := client.Integrations.DeleteTwilioIntegration(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListWebPushIntegrations

Retrieves the current Web Push integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/web_push`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`WebpushConfigCollection`

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

response, err := client.Integrations.ListWebPushIntegrations(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveWebPushIntegration

Updates or creates the Web Push integration for the project.

- HTTP Method: `PUT`
- Endpoint: `/integrations/web_push`

**Parameters**

| Name                 | Type                 | Required | Description                 |
| :------------------- | :------------------- | :------- | :-------------------------- |
| ctx                  | Context              | ✅       | Default go language context |
| webpushConfigPayload | WebpushConfigPayload | ✅       |                             |

**Return Type**

`WebpushConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


request := integrations.WebpushConfigPayload{
  PrivateKey: util.ToPointer("PrivateKey"),
  PublicKey: util.ToPointer("PublicKey"),
}

response, err := client.Integrations.SaveWebPushIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteWebPushIntegration

Deletes the Web Push integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/web_push`

**Parameters**

| Name   | Type                                  | Required | Description                   |
| :----- | :------------------------------------ | :------- | :---------------------------- |
| ctx    | Context                               | ✅       | Default go language context   |
| params | DeleteWebPushIntegrationRequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
client := client.NewClient(config)


params := integrations.DeleteWebPushIntegrationRequestParams{

}

response, err := client.Integrations.DeleteWebPushIntegration(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

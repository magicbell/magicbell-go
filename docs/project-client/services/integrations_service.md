# IntegrationsService

A list of all methods in the `IntegrationsService` service. Click on the method name to view detailed information about that method.

| Methods                                                                   | Description                                                                                                                                                               |
| :------------------------------------------------------------------------ | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| [ListIntegrations](#listintegrations)                                     | Lists all available and configured integrations for the project. Returns a summary of each integration including its type, status, and basic configuration information.   |
| [ListApnsIntegrations](#listapnsintegrations)                             | Retrieves the current APNs integration configurations for a specific integration type in the project. Returns configuration details and status information.               |
| [SaveApnsIntegration](#saveapnsintegration)                               | Updates or creates the APNs integration for the project.                                                                                                                  |
| [DeleteApnsIntegration](#deleteapnsintegration)                           | Deletes the APNs integration configuration from the project. This will disable the integration's functionality within the project.                                        |
| [ListEventsourceIntegrations](#listeventsourceintegrations)               | Retrieves the current EventSource integration configurations for a specific integration type in the project. Returns configuration details and status information.        |
| [SaveEventsourceIntegration](#saveeventsourceintegration)                 | Updates or creates the EventSource integration for the project.                                                                                                           |
| [DeleteEventsourceIntegration](#deleteeventsourceintegration)             | Deletes the EventSource integration configuration from the project. This will disable the integration's functionality within the project.                                 |
| [ListExpoIntegrations](#listexpointegrations)                             | Retrieves the current Expo integration configurations for a specific integration type in the project. Returns configuration details and status information.               |
| [SaveExpoIntegration](#saveexpointegration)                               | Updates or creates the Expo integration for the project.                                                                                                                  |
| [DeleteExpoIntegration](#deleteexpointegration)                           | Deletes the Expo integration configuration from the project. This will disable the integration's functionality within the project.                                        |
| [ListFcmIntegrations](#listfcmintegrations)                               | Retrieves the current FCM integration configurations for a specific integration type in the project. Returns configuration details and status information.                |
| [SaveFcmIntegration](#savefcmintegration)                                 | Updates or creates the FCM integration for the project.                                                                                                                   |
| [DeleteFcmIntegration](#deletefcmintegration)                             | Deletes the FCM integration configuration from the project. This will disable the integration's functionality within the project.                                         |
| [ListGithubIntegrations](#listgithubintegrations)                         | Retrieves the current GitHub integration configurations for a specific integration type in the project. Returns configuration details and status information.             |
| [SaveGithubIntegration](#savegithubintegration)                           | Updates or creates the GitHub integration for the project.                                                                                                                |
| [DeleteGithubIntegration](#deletegithubintegration)                       | Deletes the GitHub integration configuration from the project. This will disable the integration's functionality within the project.                                      |
| [ListInboxIntegrations](#listinboxintegrations)                           | Retrieves the current Inbox integration configurations for a specific integration type in the project. Returns configuration details and status information.              |
| [SaveInboxIntegration](#saveinboxintegration)                             | Updates or creates the Inbox integration for the project.                                                                                                                 |
| [DeleteInboxIntegration](#deleteinboxintegration)                         | Deletes the Inbox integration configuration from the project. This will disable the integration's functionality within the project.                                       |
| [ListMagicbellSlackbotIntegrations](#listmagicbellslackbotintegrations)   | Retrieves the current MagicBell SlackBot integration configurations for a specific integration type in the project. Returns configuration details and status information. |
| [SaveMagicbellSlackbotIntegration](#savemagicbellslackbotintegration)     | Updates or creates the MagicBell SlackBot integration for the project.                                                                                                    |
| [DeleteMagicbellSlackbotIntegration](#deletemagicbellslackbotintegration) | Deletes the MagicBell SlackBot integration configuration from the project. This will disable the integration's functionality within the project.                          |
| [ListMailgunIntegrations](#listmailgunintegrations)                       | Retrieves the current Mailgun integration configurations for a specific integration type in the project. Returns configuration details and status information.            |
| [SaveMailgunIntegration](#savemailgunintegration)                         | Updates or creates the Mailgun integration for the project.                                                                                                               |
| [DeleteMailgunIntegration](#deletemailgunintegration)                     | Deletes the Mailgun integration configuration from the project. This will disable the integration's functionality within the project.                                     |
| [ListPingEmailIntegrations](#listpingemailintegrations)                   | Retrieves the current Ping Email integration configurations for a specific integration type in the project. Returns configuration details and status information.         |
| [SavePingEmailIntegration](#savepingemailintegration)                     | Updates or creates the Ping Email integration for the project.                                                                                                            |
| [DeletePingEmailIntegration](#deletepingemailintegration)                 | Deletes the Ping Email integration configuration from the project. This will disable the integration's functionality within the project.                                  |
| [ListSendgridIntegrations](#listsendgridintegrations)                     | Retrieves the current SendGrid integration configurations for a specific integration type in the project. Returns configuration details and status information.           |
| [SaveSendgridIntegration](#savesendgridintegration)                       | Updates or creates the SendGrid integration for the project.                                                                                                              |
| [DeleteSendgridIntegration](#deletesendgridintegration)                   | Deletes the SendGrid integration configuration from the project. This will disable the integration's functionality within the project.                                    |
| [ListSesIntegrations](#listsesintegrations)                               | Retrieves the current Amazon SES integration configurations for a specific integration type in the project. Returns configuration details and status information.         |
| [SaveSesIntegration](#savesesintegration)                                 | Updates or creates the Amazon SES integration for the project.                                                                                                            |
| [DeleteSesIntegration](#deletesesintegration)                             | Deletes the Amazon SES integration configuration from the project. This will disable the integration's functionality within the project.                                  |
| [ListSlackIntegrations](#listslackintegrations)                           | Retrieves the current Slack integration configurations for a specific integration type in the project. Returns configuration details and status information.              |
| [SaveSlackIntegration](#saveslackintegration)                             | Updates or creates the Slack integration for the project.                                                                                                                 |
| [DeleteSlackIntegration](#deleteslackintegration)                         | Deletes the Slack integration configuration from the project. This will disable the integration's functionality within the project.                                       |
| [ListSmtpIntegrations](#listsmtpintegrations)                             | Retrieves the current SMTP integration configurations for a specific integration type in the project. Returns configuration details and status information.               |
| [SaveSmtpIntegration](#savesmtpintegration)                               | Updates or creates the SMTP integration for the project.                                                                                                                  |
| [DeleteSmtpIntegration](#deletesmtpintegration)                           | Deletes the SMTP integration configuration from the project. This will disable the integration's functionality within the project.                                        |
| [ListStripeIntegrations](#liststripeintegrations)                         | Retrieves the current Stripe integration configurations for a specific integration type in the project. Returns configuration details and status information.             |
| [SaveStripeIntegration](#savestripeintegration)                           | Updates or creates the Stripe integration for the project.                                                                                                                |
| [DeleteStripeIntegration](#deletestripeintegration)                       | Deletes the Stripe integration configuration from the project. This will disable the integration's functionality within the project.                                      |
| [ListTwilioIntegrations](#listtwiliointegrations)                         | Retrieves the current Twilio integration configurations for a specific integration type in the project. Returns configuration details and status information.             |
| [SaveTwilioIntegration](#savetwiliointegration)                           | Updates or creates the Twilio integration for the project.                                                                                                                |
| [DeleteTwilioIntegration](#deletetwiliointegration)                       | Deletes the Twilio integration configuration from the project. This will disable the integration's functionality within the project.                                      |
| [ListWebPushIntegrations](#listwebpushintegrations)                       | Retrieves the current Web Push integration configurations for a specific integration type in the project. Returns configuration details and status information.           |
| [SaveWebPushIntegration](#savewebpushintegration)                         | Updates or creates the Web Push integration for the project.                                                                                                              |
| [DeleteWebPushIntegration](#deletewebpushintegration)                     | Deletes the Web Push integration configuration from the project. This will disable the integration's functionality within the project.                                    |

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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := integrations.ListIntegrationsRequestParams{
  Limit: util.ToPointer(int64(3)),
  StartingAfter: util.ToPointer("starting_after"),
  EndingBefore: util.ToPointer("ending_before"),
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

badge := integrations.BADGE_UNREAD

payloadVersion := integrations.PAYLOAD_VERSION_1

request := integrations.ApnsConfigPayload{
  AppId: util.ToPointer("app_id"),
  Badge: &badge,
  Certificate: util.ToPointer("BEGIN PRIVATE KEY--------
tahhm84591o=
----------END PRIVATE KEYYYYYY-------"),
  KeyId: util.ToPointer("sintdolor "),
  PayloadVersion: &payloadVersion,
  TeamId: util.ToPointer("laborumiru"),
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := integrations.DeleteApnsIntegrationRequestParams{
  Id: util.ToPointer("id"),
}

response, err := client.Integrations.DeleteApnsIntegration(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListEventsourceIntegrations

Retrieves the current EventSource integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/eventsource`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`EventSourceConfigCollection`

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

response, err := client.Integrations.ListEventsourceIntegrations(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveEventsourceIntegration

Updates or creates the EventSource integration for the project.

- HTTP Method: `PUT`
- Endpoint: `/integrations/eventsource`

**Parameters**

| Name                     | Type                     | Required | Description                 |
| :----------------------- | :----------------------- | :------- | :-------------------------- |
| ctx                      | Context                  | ✅       | Default go language context |
| eventSourceConfigPayload | EventSourceConfigPayload | ✅       |                             |

**Return Type**

`EventSourceConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


request := integrations.EventSourceConfigPayload{
  Source: util.ToPointer("source"),
}

response, err := client.Integrations.SaveEventsourceIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteEventsourceIntegration

Deletes the EventSource integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/eventsource`

**Parameters**

| Name   | Type                                      | Required | Description                   |
| :----- | :---------------------------------------- | :------- | :---------------------------- |
| ctx    | Context                                   | ✅       | Default go language context   |
| params | DeleteEventsourceIntegrationRequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := integrations.DeleteEventsourceIntegrationRequestParams{
  Id: util.ToPointer("id"),
}

response, err := client.Integrations.DeleteEventsourceIntegration(context.Background(), params)
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


request := integrations.ExpoConfigPayload{
  AccessToken: util.ToPointer("access_token"),
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := integrations.DeleteExpoIntegrationRequestParams{
  Id: util.ToPointer("id"),
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

type_ := integrations.TYPE_SERVICE_ACCOUNT

request := integrations.FcmConfigPayload{
  AuthProviderX509CertUrl: util.ToPointer("auth_provider_x509_cert_url"),
  AuthUri: util.ToPointer("auth_uri"),
  ClientEmail: util.ToPointer("client_email"),
  ClientId: util.ToPointer("client_id"),
  ClientX509CertUrl: util.ToPointer("client_x509_cert_url"),
  PrivateKey: util.ToPointer(" BEGINNNHB--------
BRMn2Y=
--- ENDUZSZ-----------"),
  PrivateKeyId: util.ToPointer("private_key_id"),
  ProjectId: util.ToPointer("project_id"),
  TokenUri: util.ToPointer("token_uri"),
  Type_: &type_,
  UniverseDomain: util.ToPointer("universe_domain"),
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := integrations.DeleteFcmIntegrationRequestParams{
  Id: util.ToPointer("id"),
}

response, err := client.Integrations.DeleteFcmIntegration(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListGithubIntegrations

Retrieves the current GitHub integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/github`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`GithubConfigCollection`

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

response, err := client.Integrations.ListGithubIntegrations(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveGithubIntegration

Updates or creates the GitHub integration for the project.

- HTTP Method: `PUT`
- Endpoint: `/integrations/github`

**Parameters**

| Name                | Type                | Required | Description                 |
| :------------------ | :------------------ | :------- | :-------------------------- |
| ctx                 | Context             | ✅       | Default go language context |
| githubConfigPayload | GithubConfigPayload | ✅       |                             |

**Return Type**

`GithubConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


request := integrations.GithubConfigPayload{
  WebhookSigningSecret: util.ToPointer("webhook_signing_secret"),
}

response, err := client.Integrations.SaveGithubIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteGithubIntegration

Deletes the GitHub integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/github`

**Parameters**

| Name   | Type                                 | Required | Description                   |
| :----- | :----------------------------------- | :------- | :---------------------------- |
| ctx    | Context                              | ✅       | Default go language context   |
| params | DeleteGithubIntegrationRequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := integrations.DeleteGithubIntegrationRequestParams{
  Id: util.ToPointer("id"),
}

response, err := client.Integrations.DeleteGithubIntegration(context.Background(), params)
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


images := integrations.Images{
  EmptyInboxUrl: util.ToPointer("emptyInboxUrl"),
}


banner := integrations.Banner{
  BackgroundColor: util.ToPointer("backgroundColor"),
  BackgroundOpacity: util.ToPointer(float64(0.65)),
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

notification := integrations.Notification{
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
  Notification: &notification,
  UnseenBadge: &unseenBadge,
}

request := integrations.InboxConfigPayload{
  Images: &images,
  Locale: util.ToPointer(util.Nullable[string]{ Value: "locale" }),
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := integrations.DeleteInboxIntegrationRequestParams{
  Id: util.ToPointer("id"),
}

response, err := client.Integrations.DeleteInboxIntegration(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListMagicbellSlackbotIntegrations

Retrieves the current MagicBell SlackBot integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/magicbell_slackbot`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`SlackBotConfigCollection`

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

response, err := client.Integrations.ListMagicbellSlackbotIntegrations(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveMagicbellSlackbotIntegration

Updates or creates the MagicBell SlackBot integration for the project.

- HTTP Method: `PUT`
- Endpoint: `/integrations/magicbell_slackbot`

**Parameters**

| Name                  | Type                  | Required | Description                 |
| :-------------------- | :-------------------- | :------- | :-------------------------- |
| ctx                   | Context               | ✅       | Default go language context |
| slackBotConfigPayload | SlackBotConfigPayload | ✅       |                             |

**Return Type**

`SlackBotConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


request := integrations.SlackBotConfigPayload{
  Enabled: util.ToPointer(true),
}

response, err := client.Integrations.SaveMagicbellSlackbotIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteMagicbellSlackbotIntegration

Deletes the MagicBell SlackBot integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/magicbell_slackbot`

**Parameters**

| Name   | Type                                            | Required | Description                   |
| :----- | :---------------------------------------------- | :------- | :---------------------------- |
| ctx    | Context                                         | ✅       | Default go language context   |
| params | DeleteMagicbellSlackbotIntegrationRequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := integrations.DeleteMagicbellSlackbotIntegrationRequestParams{
  Id: util.ToPointer("id"),
}

response, err := client.Integrations.DeleteMagicbellSlackbotIntegration(context.Background(), params)
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


mailgunConfigPayloadFrom := integrations.MailgunConfigPayloadFrom{
  Email: util.ToPointer("email"),
  Name: util.ToPointer(util.Nullable[string]{ Value: "name" }),
}

region := integrations.REGION_US

request := integrations.MailgunConfigPayload{
  ApiKey: util.ToPointer("api_key"),
  Domain: util.ToPointer("domain"),
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := integrations.DeleteMailgunIntegrationRequestParams{
  Id: util.ToPointer("id"),
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


request := integrations.PingConfigPayload{
  Url: util.ToPointer("url"),
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := integrations.DeletePingEmailIntegrationRequestParams{
  Id: util.ToPointer("id"),
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


sendgridConfigPayloadFrom := integrations.SendgridConfigPayloadFrom{
  Email: util.ToPointer("email"),
  Name: util.ToPointer(util.Nullable[string]{ Value: "name" }),
}


sendgridConfigPayloadReplyTo := integrations.SendgridConfigPayloadReplyTo{
  Email: util.ToPointer("email"),
  Name: util.ToPointer(util.Nullable[string]{ Value: "name" }),
}

request := integrations.SendgridConfigPayload{
  ApiKey: util.ToPointer("api_key"),
  From: &sendgridConfigPayloadFrom,
  ReplyTo: &sendgridConfigPayloadReplyTo,
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := integrations.DeleteSendgridIntegrationRequestParams{
  Id: util.ToPointer("id"),
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


sesConfigPayloadFrom := integrations.SesConfigPayloadFrom{
  Email: util.ToPointer("email"),
  Name: util.ToPointer(util.Nullable[string]{ Value: "name" }),
}

request := integrations.SesConfigPayload{
  From: &sesConfigPayloadFrom,
  KeyId: util.ToPointer("key_id"),
  Region: util.ToPointer("region"),
  SecretKey: util.ToPointer("secret_key"),
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := integrations.DeleteSesIntegrationRequestParams{
  Id: util.ToPointer("id"),
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


request := integrations.SlackConfigPayload{
  AppId: util.ToPointer("N"),
  ClientId: util.ToPointer("26.63"),
  ClientSecret: util.ToPointer("irure proidentincididunt exsit E"),
  SigningSecret: util.ToPointer("incididunt laborisculpa magna al"),
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := integrations.DeleteSlackIntegrationRequestParams{
  Id: util.ToPointer("id"),
}

response, err := client.Integrations.DeleteSlackIntegration(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListSmtpIntegrations

Retrieves the current SMTP integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/smtp`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`SmtpConfigObjectCollection`

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

response, err := client.Integrations.ListSmtpIntegrations(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveSmtpIntegration

Updates or creates the SMTP integration for the project.

- HTTP Method: `PUT`
- Endpoint: `/integrations/smtp`

**Parameters**

| Name       | Type       | Required | Description                 |
| :--------- | :--------- | :------- | :-------------------------- |
| ctx        | Context    | ✅       | Default go language context |
| smtpConfig | SmtpConfig | ✅       |                             |

**Return Type**

`SmtpConfig`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


smtpConfigFrom := integrations.SmtpConfigFrom{
  Email: util.ToPointer("email"),
  Name: util.ToPointer("name"),
}


smtpConfigReplyTo := integrations.SmtpConfigReplyTo{
  Email: util.ToPointer("email"),
  Name: util.ToPointer("name"),
}

security := integrations.SECURITY_NONE

request := integrations.SmtpConfig{
  From: &smtpConfigFrom,
  Host: util.ToPointer("smtp.gmail.com"),
  Password: util.ToPointer("password"),
  Port: util.ToPointer(int64(25)),
  ReplyTo: &smtpConfigReplyTo,
  Security: &security,
  Username: util.ToPointer("username"),
}

response, err := client.Integrations.SaveSmtpIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteSmtpIntegration

Deletes the SMTP integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/smtp`

**Parameters**

| Name   | Type                               | Required | Description                   |
| :----- | :--------------------------------- | :------- | :---------------------------- |
| ctx    | Context                            | ✅       | Default go language context   |
| params | DeleteSmtpIntegrationRequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := integrations.DeleteSmtpIntegrationRequestParams{
  Id: util.ToPointer("id"),
}

response, err := client.Integrations.DeleteSmtpIntegration(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListStripeIntegrations

Retrieves the current Stripe integration configurations for a specific integration type in the project. Returns configuration details and status information.

- HTTP Method: `GET`
- Endpoint: `/integrations/stripe`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`StripeConfigCollection`

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

response, err := client.Integrations.ListStripeIntegrations(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveStripeIntegration

Updates or creates the Stripe integration for the project.

- HTTP Method: `PUT`
- Endpoint: `/integrations/stripe`

**Parameters**

| Name                | Type                | Required | Description                 |
| :------------------ | :------------------ | :------- | :-------------------------- |
| ctx                 | Context             | ✅       | Default go language context |
| stripeConfigPayload | StripeConfigPayload | ✅       |                             |

**Return Type**

`StripeConfigPayload`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


request := integrations.StripeConfigPayload{
  Id: util.ToPointer("id"),
  WebhookSigningSecret: util.ToPointer("webhook_signing_secret"),
}

response, err := client.Integrations.SaveStripeIntegration(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteStripeIntegration

Deletes the Stripe integration configuration from the project. This will disable the integration's functionality within the project.

- HTTP Method: `DELETE`
- Endpoint: `/integrations/stripe`

**Parameters**

| Name   | Type                                 | Required | Description                   |
| :----- | :----------------------------------- | :------- | :---------------------------- |
| ctx    | Context                              | ✅       | Default go language context   |
| params | DeleteStripeIntegrationRequestParams | ✅       | Additional request parameters |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := integrations.DeleteStripeIntegrationRequestParams{
  Id: util.ToPointer("id"),
}

response, err := client.Integrations.DeleteStripeIntegration(context.Background(), params)
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


request := integrations.TwilioConfigPayload{
  AccountSid: util.ToPointer("account_sid"),
  ApiKey: util.ToPointer("api_key"),
  ApiSecret: util.ToPointer("api_secret"),
  From: util.ToPointer("+8547811425211"),
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := integrations.DeleteTwilioIntegrationRequestParams{
  Id: util.ToPointer("id"),
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


request := integrations.WebpushConfigPayload{
  PrivateKey: util.ToPointer("private_key"),
  PublicKey: util.ToPointer("public_key"),
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
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/integrations"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := integrations.DeleteWebPushIntegrationRequestParams{
  Id: util.ToPointer("id"),
}

response, err := client.Integrations.DeleteWebPushIntegration(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

<!-- This file was generated by liblab | https://liblab.com/ -->

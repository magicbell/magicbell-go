# BroadcastsService

A list of all methods in the `BroadcastsService` service. Click on the method name to view detailed information about that method.

| Methods                             | Description                                                                                                                                      |
| :---------------------------------- | :----------------------------------------------------------------------------------------------------------------------------------------------- |
| [ListBroadcasts](#listbroadcasts)   | Retrieves a paginated list of broadcasts for the project. Returns basic information about each broadcast including its creation time and status. |
| [CreateBroadcast](#createbroadcast) | Creates a new broadcast. When a broadcast is created, it generates individual notifications for relevant users within the project.               |
| [FetchBroadcast](#fetchbroadcast)   | Retrieves detailed information about a specific broadcast by its ID. Includes the broadcast's configuration and current status.                  |

## ListBroadcasts

Retrieves a paginated list of broadcasts for the project. Returns basic information about each broadcast including its creation time and status.

- HTTP Method: `GET`
- Endpoint: `/broadcasts`

**Parameters**

| Name   | Type                        | Required | Description                   |
| :----- | :-------------------------- | :------- | :---------------------------- |
| ctx    | Context                     | ✅       | Default go language context   |
| params | ListBroadcastsRequestParams | ✅       | Additional request parameters |

**Return Type**

`BroadcastCollection`

**Example Usage Code Snippet**

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

## CreateBroadcast

Creates a new broadcast. When a broadcast is created, it generates individual notifications for relevant users within the project.

- HTTP Method: `POST`
- Endpoint: `/broadcasts`

**Parameters**

| Name      | Type      | Required | Description                 |
| :-------- | :-------- | :------- | :-------------------------- |
| ctx       | Context   | ✅       | Default go language context |
| broadcast | Broadcast | ✅       |                             |

**Return Type**

`Broadcast`

**Example Usage Code Snippet**

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


email := broadcasts.Email{
  ActionUrl: util.ToPointer(util.Nullable[string]{ Value: "action_url" }),
  Content: util.ToPointer("content"),
  Title: util.ToPointer("title"),
}


inApp := broadcasts.InApp{
  ActionUrl: util.ToPointer(util.Nullable[string]{ Value: "action_url" }),
  Content: util.ToPointer("content"),
  Title: util.ToPointer("title"),
}


mobilePush := broadcasts.MobilePush{
  ActionUrl: util.ToPointer(util.Nullable[string]{ Value: "action_url" }),
  Content: util.ToPointer("content"),
  Title: util.ToPointer("title"),
}


sms := broadcasts.Sms{
  ActionUrl: util.ToPointer(util.Nullable[string]{ Value: "action_url" }),
  Content: util.ToPointer("content"),
  Title: util.ToPointer("title"),
}

overridesChannels := broadcasts.OverridesChannels{
  Email: &email,
  InApp: &inApp,
  MobilePush: &mobilePush,
  Sms: &sms,
}


providers := broadcasts.Providers{
  Apns: []byte{},
  Expo: []byte{},
  Fcm: []byte{},
  Mailgun: []byte{},
  Sendgrid: []byte{},
  Ses: []byte{},
  Slack: []byte{},
  Teams: []byte{},
  Twilio: []byte{},
  WebPush: []byte{},
}

overrides := broadcasts.Overrides{
  Channels: &overridesChannels,
  Providers: &providers,
}


user := shared.User{
  CreatedAt: util.ToPointer(util.Nullable[string]{ Value: "created_at" }),
  CustomAttributes: []byte{},
  Email: util.ToPointer(util.Nullable[string]{ Value: "email" }),
  ExternalId: util.ToPointer(util.Nullable[string]{ Value: "external_id" }),
  FirstName: util.ToPointer(util.Nullable[string]{ Value: "first_name" }),
  Id: util.ToPointer("id"),
  LastName: util.ToPointer(util.Nullable[string]{ Value: "last_name" }),
  LastNotifiedAt: util.ToPointer(util.Nullable[string]{ Value: "last_notified_at" }),
  LastSeenAt: util.ToPointer(util.Nullable[string]{ Value: "last_seen_at" }),
  UpdatedAt: util.ToPointer(util.Nullable[string]{ Value: "updated_at" }),
}


errors := broadcasts.Errors{
  Message: util.ToPointer("message"),
}

statusStatus := broadcasts.STATUS_STATUS_ENQUEUED


summary := broadcasts.Summary{
  Failures: util.ToPointer(int64(8)),
  Total: util.ToPointer(int64(8)),
}

broadcastStatus := broadcasts.BroadcastStatus{
  Errors: []broadcasts.Errors{errors},
  Status: &statusStatus,
  Summary: &summary,
}

request := broadcasts.Broadcast{
  ActionUrl: util.ToPointer(util.Nullable[string]{ Value: "action_url" }),
  Category: util.ToPointer(util.Nullable[string]{ Value: "category" }),
  Content: util.ToPointer(util.Nullable[string]{ Value: "content" }),
  CreatedAt: util.ToPointer("created_at"),
  CustomAttributes: []byte{},
  Id: util.ToPointer("id"),
  Overrides: &overrides,
  Recipients: []shared.User{user},
  Status: &broadcastStatus,
  Title: util.ToPointer("title"),
  Topic: util.ToPointer(util.Nullable[string]{ Value: "topic" }),
}

response, err := client.Broadcasts.CreateBroadcast(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchBroadcast

Retrieves detailed information about a specific broadcast by its ID. Includes the broadcast's configuration and current status.

- HTTP Method: `GET`
- Endpoint: `/broadcasts/{broadcast_id}`

**Parameters**

| Name        | Type    | Required | Description                 |
| :---------- | :------ | :------- | :-------------------------- |
| ctx         | Context | ✅       | Default go language context |
| broadcastId | string  | ✅       |                             |

**Return Type**

`Broadcast`

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

response, err := client.Broadcasts.FetchBroadcast(context.Background(), "broadcast_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

<!-- This file was generated by liblab | https://liblab.com/ -->

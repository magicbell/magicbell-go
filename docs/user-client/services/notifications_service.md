# NotificationsService

A list of all methods in the `NotificationsService` service. Click on the method name to view detailed information about that method.

| Methods                                                         | Description                                                                                     |
| :-------------------------------------------------------------- | :---------------------------------------------------------------------------------------------- |
| [ListNotifications](#listnotifications)                         | Lists all notifications for a user.                                                             |
| [ArchiveAllNotifications](#archiveallnotifications)             | Archive all notifications.                                                                      |
| [MarkAllNotificationsRead](#markallnotificationsread)           | Marks all notifications as read.                                                                |
| [FetchUnreadNotificationsCount](#fetchunreadnotificationscount) | Returns the count of unread notifications for a user. Supports filtering by category and topic. |
| [FetchNotification](#fetchnotification)                         | Gets a notification by ID.                                                                      |
| [ArchiveNotification](#archivenotification)                     | Archive a notification.                                                                         |
| [MarkNotificationRead](#marknotificationread)                   | Marks a notification as read.                                                                   |
| [UnarchiveNotification](#unarchivenotification)                 | Unarchives a notification.                                                                      |
| [MarkNotificationUnread](#marknotificationunread)               | Marks a notification as unread.                                                                 |

## ListNotifications

Lists all notifications for a user.

- HTTP Method: `GET`
- Endpoint: `/notifications`

**Parameters**

| Name   | Type                           | Required | Description                   |
| :----- | :----------------------------- | :------- | :---------------------------- |
| ctx    | Context                        | ✅       | Default go language context   |
| params | ListNotificationsRequestParams | ✅       | Additional request parameters |

**Return Type**

`NotificationCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/notifications"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := notifications.ListNotificationsRequestParams{
  Limit: util.ToPointer(int64(5)),
  StartingAfter: util.ToPointer("starting_after"),
  EndingBefore: util.ToPointer("ending_before"),
  Status: util.ToPointer("status"),
  Category: util.ToPointer("category"),
  Topic: util.ToPointer("topic"),
}

response, err := client.Notifications.ListNotifications(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ArchiveAllNotifications

Archive all notifications.

- HTTP Method: `POST`
- Endpoint: `/notifications/archive`

**Parameters**

| Name   | Type                                 | Required | Description                   |
| :----- | :----------------------------------- | :------- | :---------------------------- |
| ctx    | Context                              | ✅       | Default go language context   |
| params | ArchiveAllNotificationsRequestParams | ✅       | Additional request parameters |

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
  "github.com/magicbell/magicbell-go/pkg/user-client/notifications"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := notifications.ArchiveAllNotificationsRequestParams{
  Category: util.ToPointer("category"),
  Topic: util.ToPointer("topic"),
}

response, err := client.Notifications.ArchiveAllNotifications(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## MarkAllNotificationsRead

Marks all notifications as read.

- HTTP Method: `POST`
- Endpoint: `/notifications/read`

**Parameters**

| Name   | Type                                  | Required | Description                   |
| :----- | :------------------------------------ | :------- | :---------------------------- |
| ctx    | Context                               | ✅       | Default go language context   |
| params | MarkAllNotificationsReadRequestParams | ✅       | Additional request parameters |

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
  "github.com/magicbell/magicbell-go/pkg/user-client/notifications"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := notifications.MarkAllNotificationsReadRequestParams{
  Category: util.ToPointer("category"),
  Topic: util.ToPointer("topic"),
}

response, err := client.Notifications.MarkAllNotificationsRead(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchUnreadNotificationsCount

Returns the count of unread notifications for a user. Supports filtering by category and topic.

- HTTP Method: `GET`
- Endpoint: `/notifications/unread/count`

**Parameters**

| Name   | Type                                       | Required | Description                   |
| :----- | :----------------------------------------- | :------- | :---------------------------- |
| ctx    | Context                                    | ✅       | Default go language context   |
| params | FetchUnreadNotificationsCountRequestParams | ✅       | Additional request parameters |

**Return Type**

`CountResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/user-client/client"
  "github.com/magicbell/magicbell-go/pkg/user-client/util"
  "github.com/magicbell/magicbell-go/pkg/user-client/notifications"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := notifications.FetchUnreadNotificationsCountRequestParams{
  Category: util.ToPointer("category"),
  Topic: util.ToPointer("topic"),
}

response, err := client.Notifications.FetchUnreadNotificationsCount(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchNotification

Gets a notification by ID.

- HTTP Method: `GET`
- Endpoint: `/notifications/{notification_id}`

**Parameters**

| Name           | Type    | Required | Description                 |
| :------------- | :------ | :------- | :-------------------------- |
| ctx            | Context | ✅       | Default go language context |
| notificationId | string  | ✅       |                             |

**Return Type**

`Notification`

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

response, err := client.Notifications.FetchNotification(context.Background(), "notification_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ArchiveNotification

Archive a notification.

- HTTP Method: `POST`
- Endpoint: `/notifications/{notification_id}/archive`

**Parameters**

| Name           | Type    | Required | Description                 |
| :------------- | :------ | :------- | :-------------------------- |
| ctx            | Context | ✅       | Default go language context |
| notificationId | string  | ✅       |                             |

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

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Notifications.ArchiveNotification(context.Background(), "notification_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## MarkNotificationRead

Marks a notification as read.

- HTTP Method: `POST`
- Endpoint: `/notifications/{notification_id}/read`

**Parameters**

| Name           | Type    | Required | Description                 |
| :------------- | :------ | :------- | :-------------------------- |
| ctx            | Context | ✅       | Default go language context |
| notificationId | string  | ✅       |                             |

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

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Notifications.MarkNotificationRead(context.Background(), "notification_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## UnarchiveNotification

Unarchives a notification.

- HTTP Method: `POST`
- Endpoint: `/notifications/{notification_id}/unarchive`

**Parameters**

| Name           | Type    | Required | Description                 |
| :------------- | :------ | :------- | :-------------------------- |
| ctx            | Context | ✅       | Default go language context |
| notificationId | string  | ✅       |                             |

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

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Notifications.UnarchiveNotification(context.Background(), "notification_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## MarkNotificationUnread

Marks a notification as unread.

- HTTP Method: `POST`
- Endpoint: `/notifications/{notification_id}/unread`

**Parameters**

| Name           | Type    | Required | Description                 |
| :------------- | :------ | :------- | :-------------------------- |
| ctx            | Context | ✅       | Default go language context |
| notificationId | string  | ✅       |                             |

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

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Notifications.MarkNotificationUnread(context.Background(), "notification_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

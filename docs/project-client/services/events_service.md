# EventsService

A list of all methods in the `EventsService` service. Click on the method name to view detailed information about that method.

| Methods                   | Description                                           |
| :------------------------ | :---------------------------------------------------- |
| [ListEvents](#listevents) | Retrieves a paginated list of events for the project. |
| [FetchEvent](#fetchevent) | Fetches a project event by its ID.                    |

## ListEvents

Retrieves a paginated list of events for the project.

- HTTP Method: `GET`
- Endpoint: `/events`

**Parameters**

| Name   | Type                    | Required | Description                   |
| :----- | :---------------------- | :------- | :---------------------------- |
| ctx    | Context                 | ✅       | Default go language context   |
| params | ListEventsRequestParams | ✅       | Additional request parameters |

**Return Type**

`EventCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/events"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := events.ListEventsRequestParams{
  Limit: util.ToPointer(int64(5)),
  StartingAfter: util.ToPointer("starting_after"),
  EndingBefore: util.ToPointer("ending_before"),
}

response, err := client.Events.ListEvents(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchEvent

Fetches a project event by its ID.

- HTTP Method: `GET`
- Endpoint: `/events/{event_id}`

**Parameters**

| Name    | Type    | Required | Description                 |
| :------ | :------ | :------- | :-------------------------- |
| ctx     | Context | ✅       | Default go language context |
| eventId | string  | ✅       |                             |

**Return Type**

`Event`

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

response, err := client.Events.FetchEvent(context.Background(), "event_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

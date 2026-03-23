# UsersService

A list of all methods in the `UsersService` service. Click on the method name to view detailed information about that method.

| Methods                   | Description                                                                                                                         |
| :------------------------ | :---------------------------------------------------------------------------------------------------------------------------------- |
| [ListUsers](#listusers)   | Lists all users in the project.                                                                                                     |
| [SaveUser](#saveuser)     | Creates or updates a user with the provided details. The user will be associated with the project specified in the request context. |
| [DeleteUser](#deleteuser) | Removes a user and all associated data from the project.                                                                            |

## ListUsers

Lists all users in the project.

- HTTP Method: `GET`
- Endpoint: `/users`

**Parameters**

| Name   | Type                   | Required | Description                   |
| :----- | :--------------------- | :------- | :---------------------------- |
| ctx    | Context                | ✅       | Default go language context   |
| params | ListUsersRequestParams | ✅       | Additional request parameters |

**Return Type**

`UserCollection`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/users"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


params := users.ListUsersRequestParams{
  Limit: util.ToPointer(int64(2)),
  StartingAfter: util.ToPointer("starting_after"),
  EndingBefore: util.ToPointer("ending_before"),
  Query: util.ToPointer("query"),
}

response, err := client.Users.ListUsers(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveUser

Creates or updates a user with the provided details. The user will be associated with the project specified in the request context.

- HTTP Method: `PUT`
- Endpoint: `/users`

**Parameters**

| Name | Type        | Required | Description                 |
| :--- | :---------- | :------- | :-------------------------- |
| ctx  | Context     | ✅       | Default go language context |
| user | shared.User | ✅       |                             |

**Return Type**

`shared.User`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/shared"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


request := shared.User{
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

response, err := client.Users.SaveUser(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteUser

Removes a user and all associated data from the project.

- HTTP Method: `DELETE`
- Endpoint: `/users/{user_id}`

**Parameters**

| Name   | Type    | Required | Description                 |
| :----- | :------ | :------- | :-------------------------- |
| ctx    | Context | ✅       | Default go language context |
| userId | string  | ✅       |                             |

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

)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)

response, err := client.Users.DeleteUser(context.Background(), "user_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

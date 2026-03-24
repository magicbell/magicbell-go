# WorkflowsService

A list of all methods in the `WorkflowsService` service. Click on the method name to view detailed information about that method.

| Methods                                 | Description                                              |
| :-------------------------------------- | :------------------------------------------------------- |
| [FetchWorkflows](#fetchworkflows)       | Retrieves all workflow definitions for the project       |
| [SaveWorkflow](#saveworkflow)           | Creates or updates a workflow definition for the project |
| [FetchWorkflow](#fetchworkflow)         | Retrieves a workflow definition by key                   |
| [CreateWorkflowRun](#createworkflowrun) | Executes a workflow with the provided input parameters   |
| [FetchWorkflowRun](#fetchworkflowrun)   | Retrieves the status and details of a workflow run       |
| [ListWorkflowRuns](#listworkflowruns)   | Retrieves all runs for a specific workflow               |

## FetchWorkflows

Retrieves all workflow definitions for the project

- HTTP Method: `GET`
- Endpoint: `/workflows`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`WorkflowList`

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

response, err := client.Workflows.FetchWorkflows(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## SaveWorkflow

Creates or updates a workflow definition for the project

- HTTP Method: `PUT`
- Endpoint: `/workflows`

**Parameters**

| Name               | Type               | Required | Description                 |
| :----------------- | :----------------- | :------- | :-------------------------- |
| ctx                | Context            | ✅       | Default go language context |
| workflowDefinition | WorkflowDefinition | ✅       |                             |

**Return Type**

`WorkflowDefinition`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/workflows"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


workflowDefinitionSteps := workflows.WorkflowDefinitionSteps{
  Command: util.ToPointer("command"),
  If_: util.ToPointer(util.Nullable[string]{ Value: "if" }),
  Input: []byte{},
}

request := workflows.WorkflowDefinition{
  Disabled: util.ToPointer(true),
  Key: util.ToPointer("key"),
  Steps: []workflows.WorkflowDefinitionSteps{workflowDefinitionSteps},
}

response, err := client.Workflows.SaveWorkflow(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchWorkflow

Retrieves a workflow definition by key

- HTTP Method: `GET`
- Endpoint: `/workflows/*`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`WorkflowDefinition`

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

response, err := client.Workflows.FetchWorkflow(context.Background())
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## CreateWorkflowRun

Executes a workflow with the provided input parameters

- HTTP Method: `POST`
- Endpoint: `/workflows/runs`

**Parameters**

| Name                   | Type                   | Required | Description                 |
| :--------------------- | :--------------------- | :------- | :-------------------------- |
| ctx                    | Context                | ✅       | Default go language context |
| executeWorkflowRequest | ExecuteWorkflowRequest | ✅       |                             |

**Return Type**

`CreateRunResponse`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
  "github.com/magicbell/magicbell-go/pkg/project-client/client"
  "github.com/magicbell/magicbell-go/pkg/project-client/util"
  "github.com/magicbell/magicbell-go/pkg/project-client/workflows"
)

config := clientconfig.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := client.NewClient(config)


request := workflows.ExecuteWorkflowRequest{
  Input: []byte{},
  Key: util.ToPointer("key"),
}

response, err := client.Workflows.CreateWorkflowRun(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## FetchWorkflowRun

Retrieves the status and details of a workflow run

- HTTP Method: `GET`
- Endpoint: `/workflows/runs/{run_id}`

**Parameters**

| Name  | Type    | Required | Description                 |
| :---- | :------ | :------- | :-------------------------- |
| ctx   | Context | ✅       | Default go language context |
| runId | string  | ✅       |                             |

**Return Type**

`GetRunResponse`

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

response, err := client.Workflows.FetchWorkflowRun(context.Background(), "run_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## ListWorkflowRuns

Retrieves all runs for a specific workflow

- HTTP Method: `GET`
- Endpoint: `/workflows/{workflow_key}/runs`

**Parameters**

| Name        | Type    | Required | Description                 |
| :---------- | :------ | :------- | :-------------------------- |
| ctx         | Context | ✅       | Default go language context |
| workflowKey | string  | ✅       |                             |

**Return Type**

`WorkflowRunCollection`

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

response, err := client.Workflows.ListWorkflowRuns(context.Background(), "workflow_key")
if err != nil {
  panic(err)
}

fmt.Println(response)
```

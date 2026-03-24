# GetRunResponse

**Properties**

| Name        | Type                           | Required | Description |
| :---------- | :----------------------------- | :------- | :---------- |
| CreatedAt   | string                         | ❌       |             |
| Id          | string                         | ❌       |             |
| Input       | any                            | ❌       |             |
| Status      | workflows.GetRunResponseStatus | ❌       |             |
| WorkflowKey | string                         | ❌       |             |

# GetRunResponseStatus

**Properties**

| Name        | Type   | Required | Description |
| :---------- | :----- | :------- | :---------- |
| CompletedAt | string | ❌       |             |
| Error       | string | ❌       |             |
| NextStep    | int64  | ❌       |             |
| StartedAt   | string | ❌       |             |
| State       | int64  | ❌       |             |

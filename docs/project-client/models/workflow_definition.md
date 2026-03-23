# WorkflowDefinition

**Properties**

| Name     | Type                                | Required | Description                                                            |
| :------- | :---------------------------------- | :------- | :--------------------------------------------------------------------- |
| Key      | string                              | ✅       | Unique identifier for this workflow definition.                        |
| Steps    | []workflows.WorkflowDefinitionSteps | ✅       | Ordered list describing each action that will run inside the workflow. |
| Disabled | bool                                | ❌       | When true, prevents the workflow from being triggered.                 |

# WorkflowDefinitionSteps

**Properties**

| Name    | Type   | Required | Description                                                       |
| :------ | :----- | :------- | :---------------------------------------------------------------- |
| Command | string | ✅       | Command to execute (e.g., broadcast, pause, wait, abort)          |
| If\_    | string | ❌       | JMESPath condition that must evaluate truthy for the step to run. |
| Input   | any    | ❌       | Optional payload passed to the command when it executes.          |

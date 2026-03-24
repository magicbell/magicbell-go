# ExecuteWorkflowRequest

**Properties**

| Name  | Type   | Required | Description                                                                    |
| :---- | :----- | :------- | :----------------------------------------------------------------------------- |
| Key   | string | ✅       | The unique workflow key to execute (e.g. integration.stripe.charge.succeeded). |
| Input | any    | ❌       | Optional JSON payload that will be passed as the workflow input context.       |

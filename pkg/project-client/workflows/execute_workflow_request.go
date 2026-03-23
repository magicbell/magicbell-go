package workflows

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/unmarshal"
	"github.com/magicbell/magicbell-go/pkg/project-client/util"
)

type ExecuteWorkflowRequest struct {
	// Optional JSON payload that will be passed as the workflow input context.
	Input *util.Nullable[any] `json:"input,omitempty"`
	// The unique workflow key to execute (e.g. integration.stripe.charge.succeeded).
	Key *string `json:"key,omitempty" required:"true" pattern:"^[A-Za-z0-9\_\.\-\:]+$"`
}

func (e *ExecuteWorkflowRequest) GetInput() *util.Nullable[any] {
	if e == nil {
		return nil
	}
	return e.Input
}

func (e *ExecuteWorkflowRequest) SetInput(input util.Nullable[any]) {
	e.Input = &input
}

func (e *ExecuteWorkflowRequest) SetInputNull() {
	e.Input = &util.Nullable[any]{IsNull: true}
}

func (e *ExecuteWorkflowRequest) GetKey() *string {
	if e == nil {
		return nil
	}
	return e.Key
}

func (e *ExecuteWorkflowRequest) SetKey(key string) {
	e.Key = &key
}

func (e ExecuteWorkflowRequest) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ExecuteWorkflowRequest to string"
	}
	return string(jsonData)
}

func (e *ExecuteWorkflowRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, e)
}

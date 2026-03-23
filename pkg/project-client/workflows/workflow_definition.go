package workflows

import "encoding/json"

type WorkflowDefinition struct {
	// When true, prevents the workflow from being triggered.
	Disabled *bool `json:"disabled,omitempty"`
	// Unique identifier for this workflow definition.
	Key *string `json:"key,omitempty" required:"true" minLength:"3" pattern:"^[A-Za-z0-9\_\.\-\:]+$"`
	// Ordered list describing each action that will run inside the workflow.
	Steps []WorkflowDefinitionSteps `json:"steps,omitempty" required:"true"`
}

func (w *WorkflowDefinition) GetDisabled() *bool {
	if w == nil {
		return nil
	}
	return w.Disabled
}

func (w *WorkflowDefinition) SetDisabled(disabled bool) {
	w.Disabled = &disabled
}

func (w *WorkflowDefinition) GetKey() *string {
	if w == nil {
		return nil
	}
	return w.Key
}

func (w *WorkflowDefinition) SetKey(key string) {
	w.Key = &key
}

func (w *WorkflowDefinition) GetSteps() []WorkflowDefinitionSteps {
	if w == nil {
		return nil
	}
	return w.Steps
}

func (w *WorkflowDefinition) SetSteps(steps []WorkflowDefinitionSteps) {
	w.Steps = steps
}

func (w WorkflowDefinition) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowDefinition to string"
	}
	return string(jsonData)
}

package workflows

import "encoding/json"

type WorkflowRunStatus struct {
	CompletedAt *string `json:"completed_at,omitempty"`
	Error       *string `json:"error,omitempty"`
	NextStep    *int64  `json:"next_step,omitempty"`
	StartedAt   *string `json:"started_at,omitempty"`
	State       *int64  `json:"state,omitempty"`
}

func (w *WorkflowRunStatus) GetCompletedAt() *string {
	if w == nil {
		return nil
	}
	return w.CompletedAt
}

func (w *WorkflowRunStatus) SetCompletedAt(completedAt string) {
	w.CompletedAt = &completedAt
}

func (w *WorkflowRunStatus) GetError() *string {
	if w == nil {
		return nil
	}
	return w.Error
}

func (w *WorkflowRunStatus) SetError(error string) {
	w.Error = &error
}

func (w *WorkflowRunStatus) GetNextStep() *int64 {
	if w == nil {
		return nil
	}
	return w.NextStep
}

func (w *WorkflowRunStatus) SetNextStep(nextStep int64) {
	w.NextStep = &nextStep
}

func (w *WorkflowRunStatus) GetStartedAt() *string {
	if w == nil {
		return nil
	}
	return w.StartedAt
}

func (w *WorkflowRunStatus) SetStartedAt(startedAt string) {
	w.StartedAt = &startedAt
}

func (w *WorkflowRunStatus) GetState() *int64 {
	if w == nil {
		return nil
	}
	return w.State
}

func (w *WorkflowRunStatus) SetState(state int64) {
	w.State = &state
}

func (w WorkflowRunStatus) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowRunStatus to string"
	}
	return string(jsonData)
}

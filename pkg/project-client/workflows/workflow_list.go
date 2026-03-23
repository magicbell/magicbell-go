package workflows

import "encoding/json"

type WorkflowList struct {
	Items []Items `json:"items,omitempty"`
}

func (w *WorkflowList) GetItems() []Items {
	if w == nil {
		return nil
	}
	return w.Items
}

func (w *WorkflowList) SetItems(items []Items) {
	w.Items = items
}

func (w WorkflowList) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowList to string"
	}
	return string(jsonData)
}

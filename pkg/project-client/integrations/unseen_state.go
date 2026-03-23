package integrations

import "encoding/json"

// State indicator styling for unseen notifications.
type UnseenState struct {
	// Color for the unseen state indicator.
	Color *string `json:"color,omitempty" required:"true"`
}

func (u *UnseenState) GetColor() *string {
	if u == nil {
		return nil
	}
	return u.Color
}

func (u *UnseenState) SetColor(color string) {
	u.Color = &color
}

func (u UnseenState) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UnseenState to string"
	}
	return string(jsonData)
}

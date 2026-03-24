package integrations

import "encoding/json"

// Styles applied when a notification is hovered.
type DefaultHover struct {
	// Background color on hover.
	BackgroundColor *string `json:"backgroundColor,omitempty" required:"true"`
}

func (d *DefaultHover) GetBackgroundColor() *string {
	if d == nil {
		return nil
	}
	return d.BackgroundColor
}

func (d *DefaultHover) SetBackgroundColor(backgroundColor string) {
	d.BackgroundColor = &backgroundColor
}

func (d DefaultHover) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DefaultHover to string"
	}
	return string(jsonData)
}

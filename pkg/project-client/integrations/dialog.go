package integrations

import "encoding/json"

// Styling for confirmation and action dialogs.
type Dialog struct {
	// Accent color for dialog buttons and highlights.
	AccentColor *string `json:"accentColor,omitempty" required:"true"`
	// Dialog background color.
	BackgroundColor *string `json:"backgroundColor,omitempty" required:"true"`
	// Dialog text color.
	TextColor *string `json:"textColor,omitempty" required:"true"`
}

func (d *Dialog) GetAccentColor() *string {
	if d == nil {
		return nil
	}
	return d.AccentColor
}

func (d *Dialog) SetAccentColor(accentColor string) {
	d.AccentColor = &accentColor
}

func (d *Dialog) GetBackgroundColor() *string {
	if d == nil {
		return nil
	}
	return d.BackgroundColor
}

func (d *Dialog) SetBackgroundColor(backgroundColor string) {
	d.BackgroundColor = &backgroundColor
}

func (d *Dialog) GetTextColor() *string {
	if d == nil {
		return nil
	}
	return d.TextColor
}

func (d *Dialog) SetTextColor(textColor string) {
	d.TextColor = &textColor
}

func (d Dialog) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: Dialog to string"
	}
	return string(jsonData)
}

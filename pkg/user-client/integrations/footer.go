package integrations

import "encoding/json"

// Footer styling for the inbox modal.
type Footer struct {
	// Footer background color.
	BackgroundColor *string `json:"backgroundColor,omitempty" required:"true"`
	// Border radius applied to the footer container.
	BorderRadius *string `json:"borderRadius,omitempty" required:"true"`
	// Font size used in the footer.
	FontSize *string `json:"fontSize,omitempty" required:"true"`
	// Footer text color.
	TextColor *string `json:"textColor,omitempty" required:"true"`
}

func (f *Footer) GetBackgroundColor() *string {
	if f == nil {
		return nil
	}
	return f.BackgroundColor
}

func (f *Footer) SetBackgroundColor(backgroundColor string) {
	f.BackgroundColor = &backgroundColor
}

func (f *Footer) GetBorderRadius() *string {
	if f == nil {
		return nil
	}
	return f.BorderRadius
}

func (f *Footer) SetBorderRadius(borderRadius string) {
	f.BorderRadius = &borderRadius
}

func (f *Footer) GetFontSize() *string {
	if f == nil {
		return nil
	}
	return f.FontSize
}

func (f *Footer) SetFontSize(fontSize string) {
	f.FontSize = &fontSize
}

func (f *Footer) GetTextColor() *string {
	if f == nil {
		return nil
	}
	return f.TextColor
}

func (f *Footer) SetTextColor(textColor string) {
	f.TextColor = &textColor
}

func (f Footer) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: Footer to string"
	}
	return string(jsonData)
}

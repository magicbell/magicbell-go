package integrations

import "encoding/json"

// Header styling for the inbox modal.
type Header struct {
	// Header background color.
	BackgroundColor *string `json:"backgroundColor,omitempty" required:"true"`
	// Border radius applied to the header container.
	BorderRadius *string `json:"borderRadius,omitempty" required:"true"`
	// CSS font family for the header title.
	FontFamily *string `json:"fontFamily,omitempty" required:"true"`
	// Font size used in the header.
	FontSize *string `json:"fontSize,omitempty" required:"true"`
	// Header text color.
	TextColor *string `json:"textColor,omitempty" required:"true"`
}

func (h *Header) GetBackgroundColor() *string {
	if h == nil {
		return nil
	}
	return h.BackgroundColor
}

func (h *Header) SetBackgroundColor(backgroundColor string) {
	h.BackgroundColor = &backgroundColor
}

func (h *Header) GetBorderRadius() *string {
	if h == nil {
		return nil
	}
	return h.BorderRadius
}

func (h *Header) SetBorderRadius(borderRadius string) {
	h.BorderRadius = &borderRadius
}

func (h *Header) GetFontFamily() *string {
	if h == nil {
		return nil
	}
	return h.FontFamily
}

func (h *Header) SetFontFamily(fontFamily string) {
	h.FontFamily = &fontFamily
}

func (h *Header) GetFontSize() *string {
	if h == nil {
		return nil
	}
	return h.FontSize
}

func (h *Header) SetFontSize(fontSize string) {
	h.FontSize = &fontSize
}

func (h *Header) GetTextColor() *string {
	if h == nil {
		return nil
	}
	return h.TextColor
}

func (h *Header) SetTextColor(textColor string) {
	h.TextColor = &textColor
}

func (h Header) String() string {
	jsonData, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return "error converting struct: Header to string"
	}
	return string(jsonData)
}

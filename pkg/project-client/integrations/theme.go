package integrations

import "encoding/json"

// Visual customization options for the hosted inbox widget.
type Theme struct {
	// Top banner styling options.
	Banner *Banner `json:"banner,omitempty"`
	// Styling for confirmation and action dialogs.
	Dialog *Dialog `json:"dialog,omitempty"`
	// Footer styling for the inbox modal.
	Footer *Footer `json:"footer,omitempty"`
	// Header styling for the inbox modal.
	Header *Header `json:"header,omitempty"`
	// Launcher icon styling overrides.
	Icon *Icon `json:"icon,omitempty"`
	// Styling overrides for notification list items.
	Notification *Notification `json:"notification,omitempty"`
	// Badge styling for unseen notification counts.
	UnseenBadge *UnseenBadge `json:"unseenBadge,omitempty"`
}

func (t *Theme) GetBanner() *Banner {
	if t == nil {
		return nil
	}
	return t.Banner
}

func (t *Theme) SetBanner(banner Banner) {
	t.Banner = &banner
}

func (t *Theme) GetDialog() *Dialog {
	if t == nil {
		return nil
	}
	return t.Dialog
}

func (t *Theme) SetDialog(dialog Dialog) {
	t.Dialog = &dialog
}

func (t *Theme) GetFooter() *Footer {
	if t == nil {
		return nil
	}
	return t.Footer
}

func (t *Theme) SetFooter(footer Footer) {
	t.Footer = &footer
}

func (t *Theme) GetHeader() *Header {
	if t == nil {
		return nil
	}
	return t.Header
}

func (t *Theme) SetHeader(header Header) {
	t.Header = &header
}

func (t *Theme) GetIcon() *Icon {
	if t == nil {
		return nil
	}
	return t.Icon
}

func (t *Theme) SetIcon(icon Icon) {
	t.Icon = &icon
}

func (t *Theme) GetNotification() *Notification {
	if t == nil {
		return nil
	}
	return t.Notification
}

func (t *Theme) SetNotification(notification Notification) {
	t.Notification = &notification
}

func (t *Theme) GetUnseenBadge() *UnseenBadge {
	if t == nil {
		return nil
	}
	return t.UnseenBadge
}

func (t *Theme) SetUnseenBadge(unseenBadge UnseenBadge) {
	t.UnseenBadge = &unseenBadge
}

func (t Theme) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: Theme to string"
	}
	return string(jsonData)
}

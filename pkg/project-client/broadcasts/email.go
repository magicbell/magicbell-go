package broadcasts

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/unmarshal"
	"github.com/magicbell/magicbell-go/pkg/project-client/util"
)

// Overrides for email notifications.
type Email struct {
	// The link associated with the channel-specific notification.
	ActionUrl *util.Nullable[string] `json:"action_url,omitempty" maxLength:"2048"`
	// The channel-specific content.
	Content *string `json:"content,omitempty" maxLength:"1048576"`
	// The channel-specific title.
	Title *string `json:"title,omitempty" maxLength:"255" minLength:"1"`
}

func (e *Email) GetActionUrl() *util.Nullable[string] {
	if e == nil {
		return nil
	}
	return e.ActionUrl
}

func (e *Email) SetActionUrl(actionUrl util.Nullable[string]) {
	e.ActionUrl = &actionUrl
}

func (e *Email) SetActionUrlNull() {
	e.ActionUrl = &util.Nullable[string]{IsNull: true}
}

func (e *Email) GetContent() *string {
	if e == nil {
		return nil
	}
	return e.Content
}

func (e *Email) SetContent(content string) {
	e.Content = &content
}

func (e *Email) GetTitle() *string {
	if e == nil {
		return nil
	}
	return e.Title
}

func (e *Email) SetTitle(title string) {
	e.Title = &title
}

func (e Email) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: Email to string"
	}
	return string(jsonData)
}

func (e *Email) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, e)
}

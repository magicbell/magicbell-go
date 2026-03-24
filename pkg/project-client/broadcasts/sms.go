package broadcasts

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/unmarshal"
	"github.com/magicbell/magicbell-go/pkg/project-client/util"
)

// Overrides for SMS notifications.
type Sms struct {
	// The link associated with the channel-specific notification.
	ActionUrl *util.Nullable[string] `json:"action_url,omitempty" maxLength:"2048"`
	// The channel-specific content.
	Content *string `json:"content,omitempty" maxLength:"1048576"`
	// The channel-specific title.
	Title *string `json:"title,omitempty" maxLength:"255" minLength:"1"`
}

func (s *Sms) GetActionUrl() *util.Nullable[string] {
	if s == nil {
		return nil
	}
	return s.ActionUrl
}

func (s *Sms) SetActionUrl(actionUrl util.Nullable[string]) {
	s.ActionUrl = &actionUrl
}

func (s *Sms) SetActionUrlNull() {
	s.ActionUrl = &util.Nullable[string]{IsNull: true}
}

func (s *Sms) GetContent() *string {
	if s == nil {
		return nil
	}
	return s.Content
}

func (s *Sms) SetContent(content string) {
	s.Content = &content
}

func (s *Sms) GetTitle() *string {
	if s == nil {
		return nil
	}
	return s.Title
}

func (s *Sms) SetTitle(title string) {
	s.Title = &title
}

func (s Sms) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: Sms to string"
	}
	return string(jsonData)
}

func (s *Sms) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, s)
}

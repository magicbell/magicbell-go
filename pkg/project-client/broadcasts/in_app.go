package broadcasts

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/unmarshal"
	"github.com/magicbell/magicbell-go/pkg/project-client/util"
)

// Overrides for in-app notifications.
type InApp struct {
	// The link associated with the channel-specific notification.
	ActionUrl *util.Nullable[string] `json:"action_url,omitempty" maxLength:"2048"`
	// The channel-specific content.
	Content *string `json:"content,omitempty" maxLength:"1048576"`
	// The channel-specific title.
	Title *string `json:"title,omitempty" maxLength:"255" minLength:"1"`
}

func (i *InApp) GetActionUrl() *util.Nullable[string] {
	if i == nil {
		return nil
	}
	return i.ActionUrl
}

func (i *InApp) SetActionUrl(actionUrl util.Nullable[string]) {
	i.ActionUrl = &actionUrl
}

func (i *InApp) SetActionUrlNull() {
	i.ActionUrl = &util.Nullable[string]{IsNull: true}
}

func (i *InApp) GetContent() *string {
	if i == nil {
		return nil
	}
	return i.Content
}

func (i *InApp) SetContent(content string) {
	i.Content = &content
}

func (i *InApp) GetTitle() *string {
	if i == nil {
		return nil
	}
	return i.Title
}

func (i *InApp) SetTitle(title string) {
	i.Title = &title
}

func (i InApp) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InApp to string"
	}
	return string(jsonData)
}

func (i *InApp) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, i)
}

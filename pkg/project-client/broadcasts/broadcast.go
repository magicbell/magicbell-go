package broadcasts

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/unmarshal"
	"github.com/magicbell/magicbell-go/pkg/project-client/shared"
	"github.com/magicbell/magicbell-go/pkg/project-client/util"
)

type Broadcast struct {
	// The URL recipients will be directed to when interacting with the broadcast.
	ActionUrl *util.Nullable[string] `json:"action_url,omitempty" maxLength:"2048"`
	// The label used to group broadcasts.
	Category *util.Nullable[string] `json:"category,omitempty" maxLength:"255" pattern:"^[A-Za-z0-9_\.\-/:]+$"`
	// The body content delivered with the broadcast.
	Content *util.Nullable[string] `json:"content,omitempty" maxLength:"10485760"`
	// The timestamp when the broadcast was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// Arbitrary custom data associated with the broadcast.
	CustomAttributes *util.Nullable[any] `json:"custom_attributes,omitempty"`
	// The unique id for this broadcast.
	Id *string `json:"id,omitempty"`
	// Channel- or provider-specific values that override the defaults.
	Overrides *util.Nullable[Overrides] `json:"overrides,omitempty"`
	// A collection of users or filters that determine who receives the broadcast.
	Recipients *util.Nullable[[]shared.User] `json:"recipients,omitempty" required:"true" minItems:"1" maxItems:"1000"`
	// The runtime state of the broadcast execution.
	Status *BroadcastStatus `json:"status,omitempty"`
	// The subject or headline that will be shown to recipients.
	Title *string `json:"title,omitempty" required:"true" maxLength:"255" minLength:"1"`
	// The topic that further classifies the broadcast.
	Topic *util.Nullable[string] `json:"topic,omitempty" maxLength:"255" pattern:"^[A-Za-z0-9_\.\-/:]+$"`
}

func (b *Broadcast) GetActionUrl() *util.Nullable[string] {
	if b == nil {
		return nil
	}
	return b.ActionUrl
}

func (b *Broadcast) SetActionUrl(actionUrl util.Nullable[string]) {
	b.ActionUrl = &actionUrl
}

func (b *Broadcast) SetActionUrlNull() {
	b.ActionUrl = &util.Nullable[string]{IsNull: true}
}

func (b *Broadcast) GetCategory() *util.Nullable[string] {
	if b == nil {
		return nil
	}
	return b.Category
}

func (b *Broadcast) SetCategory(category util.Nullable[string]) {
	b.Category = &category
}

func (b *Broadcast) SetCategoryNull() {
	b.Category = &util.Nullable[string]{IsNull: true}
}

func (b *Broadcast) GetContent() *util.Nullable[string] {
	if b == nil {
		return nil
	}
	return b.Content
}

func (b *Broadcast) SetContent(content util.Nullable[string]) {
	b.Content = &content
}

func (b *Broadcast) SetContentNull() {
	b.Content = &util.Nullable[string]{IsNull: true}
}

func (b *Broadcast) GetCreatedAt() *string {
	if b == nil {
		return nil
	}
	return b.CreatedAt
}

func (b *Broadcast) SetCreatedAt(createdAt string) {
	b.CreatedAt = &createdAt
}

func (b *Broadcast) GetCustomAttributes() *util.Nullable[any] {
	if b == nil {
		return nil
	}
	return b.CustomAttributes
}

func (b *Broadcast) SetCustomAttributes(customAttributes util.Nullable[any]) {
	b.CustomAttributes = &customAttributes
}

func (b *Broadcast) SetCustomAttributesNull() {
	b.CustomAttributes = &util.Nullable[any]{IsNull: true}
}

func (b *Broadcast) GetId() *string {
	if b == nil {
		return nil
	}
	return b.Id
}

func (b *Broadcast) SetId(id string) {
	b.Id = &id
}

func (b *Broadcast) GetOverrides() *util.Nullable[Overrides] {
	if b == nil {
		return nil
	}
	return b.Overrides
}

func (b *Broadcast) SetOverrides(overrides util.Nullable[Overrides]) {
	b.Overrides = &overrides
}

func (b *Broadcast) SetOverridesNull() {
	b.Overrides = &util.Nullable[Overrides]{IsNull: true}
}

func (b *Broadcast) GetRecipients() *util.Nullable[[]shared.User] {
	if b == nil {
		return nil
	}
	return b.Recipients
}

func (b *Broadcast) SetRecipients(recipients util.Nullable[[]shared.User]) {
	b.Recipients = &recipients
}

func (b *Broadcast) SetRecipientsNull() {
	b.Recipients = &util.Nullable[[]shared.User]{IsNull: true}
}

func (b *Broadcast) GetStatus() *BroadcastStatus {
	if b == nil {
		return nil
	}
	return b.Status
}

func (b *Broadcast) SetStatus(status BroadcastStatus) {
	b.Status = &status
}

func (b *Broadcast) GetTitle() *string {
	if b == nil {
		return nil
	}
	return b.Title
}

func (b *Broadcast) SetTitle(title string) {
	b.Title = &title
}

func (b *Broadcast) GetTopic() *util.Nullable[string] {
	if b == nil {
		return nil
	}
	return b.Topic
}

func (b *Broadcast) SetTopic(topic util.Nullable[string]) {
	b.Topic = &topic
}

func (b *Broadcast) SetTopicNull() {
	b.Topic = &util.Nullable[string]{IsNull: true}
}

func (b Broadcast) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: Broadcast to string"
	}
	return string(jsonData)
}

func (b *Broadcast) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

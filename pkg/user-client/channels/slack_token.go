package channels

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/user-client/internal/unmarshal"
	"github.com/magicbell/magicbell-go/pkg/user-client/util"
)

type SlackToken struct {
	// The timestamp when the token was created.
	CreatedAt *string `json:"created_at,omitempty" required:"true"`
	// The timestamp when the token was discarded, if applicable.
	DiscardedAt *util.Nullable[string] `json:"discarded_at,omitempty"`
	// The unique identifier for the token.
	Id    *string          `json:"id,omitempty" required:"true"`
	Oauth *SlackTokenOauth `json:"oauth,omitempty"`
	// The timestamp when the token metadata last changed.
	UpdatedAt *util.Nullable[string] `json:"updated_at,omitempty"`
	// Obtained directly from the incoming_webhook object in the installation response from the Slack API.
	Webhook *SlackTokenWebhook `json:"webhook,omitempty"`
}

func (s *SlackToken) GetCreatedAt() *string {
	if s == nil {
		return nil
	}
	return s.CreatedAt
}

func (s *SlackToken) SetCreatedAt(createdAt string) {
	s.CreatedAt = &createdAt
}

func (s *SlackToken) GetDiscardedAt() *util.Nullable[string] {
	if s == nil {
		return nil
	}
	return s.DiscardedAt
}

func (s *SlackToken) SetDiscardedAt(discardedAt util.Nullable[string]) {
	s.DiscardedAt = &discardedAt
}

func (s *SlackToken) SetDiscardedAtNull() {
	s.DiscardedAt = &util.Nullable[string]{IsNull: true}
}

func (s *SlackToken) GetId() *string {
	if s == nil {
		return nil
	}
	return s.Id
}

func (s *SlackToken) SetId(id string) {
	s.Id = &id
}

func (s *SlackToken) GetOauth() *SlackTokenOauth {
	if s == nil {
		return nil
	}
	return s.Oauth
}

func (s *SlackToken) SetOauth(oauth SlackTokenOauth) {
	s.Oauth = &oauth
}

func (s *SlackToken) GetUpdatedAt() *util.Nullable[string] {
	if s == nil {
		return nil
	}
	return s.UpdatedAt
}

func (s *SlackToken) SetUpdatedAt(updatedAt util.Nullable[string]) {
	s.UpdatedAt = &updatedAt
}

func (s *SlackToken) SetUpdatedAtNull() {
	s.UpdatedAt = &util.Nullable[string]{IsNull: true}
}

func (s *SlackToken) GetWebhook() *SlackTokenWebhook {
	if s == nil {
		return nil
	}
	return s.Webhook
}

func (s *SlackToken) SetWebhook(webhook SlackTokenWebhook) {
	s.Webhook = &webhook
}

func (s SlackToken) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SlackToken to string"
	}
	return string(jsonData)
}

func (s *SlackToken) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, s)
}

package channels

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/unmarshal"
	"github.com/magicbell/magicbell-go/pkg/project-client/util"
)

type InboxTokenResponse struct {
	// Realtime connection ID to restrict delivery to a specific Ably connection.
	ConnectionId *util.Nullable[string] `json:"connection_id,omitempty"`
	// The timestamp when the token was created.
	CreatedAt *string `json:"created_at,omitempty" required:"true"`
	// The timestamp when the token was discarded, if applicable.
	DiscardedAt *util.Nullable[string] `json:"discarded_at,omitempty"`
	// The unique identifier for the token.
	Id *string `json:"id,omitempty" required:"true"`
	// The in-app inbox token generated for this user.
	Token *string `json:"token,omitempty" required:"true" minLength:"64"`
	// The timestamp when the token metadata last changed.
	UpdatedAt *util.Nullable[string] `json:"updated_at,omitempty"`
}

func (i *InboxTokenResponse) GetConnectionId() *util.Nullable[string] {
	if i == nil {
		return nil
	}
	return i.ConnectionId
}

func (i *InboxTokenResponse) SetConnectionId(connectionId util.Nullable[string]) {
	i.ConnectionId = &connectionId
}

func (i *InboxTokenResponse) SetConnectionIdNull() {
	i.ConnectionId = &util.Nullable[string]{IsNull: true}
}

func (i *InboxTokenResponse) GetCreatedAt() *string {
	if i == nil {
		return nil
	}
	return i.CreatedAt
}

func (i *InboxTokenResponse) SetCreatedAt(createdAt string) {
	i.CreatedAt = &createdAt
}

func (i *InboxTokenResponse) GetDiscardedAt() *util.Nullable[string] {
	if i == nil {
		return nil
	}
	return i.DiscardedAt
}

func (i *InboxTokenResponse) SetDiscardedAt(discardedAt util.Nullable[string]) {
	i.DiscardedAt = &discardedAt
}

func (i *InboxTokenResponse) SetDiscardedAtNull() {
	i.DiscardedAt = &util.Nullable[string]{IsNull: true}
}

func (i *InboxTokenResponse) GetId() *string {
	if i == nil {
		return nil
	}
	return i.Id
}

func (i *InboxTokenResponse) SetId(id string) {
	i.Id = &id
}

func (i *InboxTokenResponse) GetToken() *string {
	if i == nil {
		return nil
	}
	return i.Token
}

func (i *InboxTokenResponse) SetToken(token string) {
	i.Token = &token
}

func (i *InboxTokenResponse) GetUpdatedAt() *util.Nullable[string] {
	if i == nil {
		return nil
	}
	return i.UpdatedAt
}

func (i *InboxTokenResponse) SetUpdatedAt(updatedAt util.Nullable[string]) {
	i.UpdatedAt = &updatedAt
}

func (i *InboxTokenResponse) SetUpdatedAtNull() {
	i.UpdatedAt = &util.Nullable[string]{IsNull: true}
}

func (i InboxTokenResponse) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InboxTokenResponse to string"
	}
	return string(jsonData)
}

func (i *InboxTokenResponse) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, i)
}

package channels

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/unmarshal"
	"github.com/magicbell/magicbell-go/pkg/project-client/util"
)

type FcmToken struct {
	// The timestamp when the token was created.
	CreatedAt *string `json:"created_at,omitempty" required:"true"`
	// The Firebase Cloud Messaging device registration token to associate with the user.
	DeviceToken *string `json:"device_token,omitempty" required:"true" minLength:"64"`
	// The timestamp when the token was discarded, if applicable.
	DiscardedAt *util.Nullable[string] `json:"discarded_at,omitempty"`
	// The unique identifier for the token.
	Id *string `json:"id,omitempty" required:"true"`
	// The timestamp when the token metadata last changed.
	UpdatedAt *util.Nullable[string] `json:"updated_at,omitempty"`
}

func (f *FcmToken) GetCreatedAt() *string {
	if f == nil {
		return nil
	}
	return f.CreatedAt
}

func (f *FcmToken) SetCreatedAt(createdAt string) {
	f.CreatedAt = &createdAt
}

func (f *FcmToken) GetDeviceToken() *string {
	if f == nil {
		return nil
	}
	return f.DeviceToken
}

func (f *FcmToken) SetDeviceToken(deviceToken string) {
	f.DeviceToken = &deviceToken
}

func (f *FcmToken) GetDiscardedAt() *util.Nullable[string] {
	if f == nil {
		return nil
	}
	return f.DiscardedAt
}

func (f *FcmToken) SetDiscardedAt(discardedAt util.Nullable[string]) {
	f.DiscardedAt = &discardedAt
}

func (f *FcmToken) SetDiscardedAtNull() {
	f.DiscardedAt = &util.Nullable[string]{IsNull: true}
}

func (f *FcmToken) GetId() *string {
	if f == nil {
		return nil
	}
	return f.Id
}

func (f *FcmToken) SetId(id string) {
	f.Id = &id
}

func (f *FcmToken) GetUpdatedAt() *util.Nullable[string] {
	if f == nil {
		return nil
	}
	return f.UpdatedAt
}

func (f *FcmToken) SetUpdatedAt(updatedAt util.Nullable[string]) {
	f.UpdatedAt = &updatedAt
}

func (f *FcmToken) SetUpdatedAtNull() {
	f.UpdatedAt = &util.Nullable[string]{IsNull: true}
}

func (f FcmToken) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FcmToken to string"
	}
	return string(jsonData)
}

func (f *FcmToken) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, f)
}

package channels

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/unmarshal"
	"github.com/magicbell/magicbell-go/pkg/project-client/util"
)

type ApnsToken struct {
	// The bundle identifier of the application registering this token. Use this to override the default identifier configured on the APNs integration.
	AppId *string `json:"app_id,omitempty" pattern:"^[a-zA-Z0-9]+(.[a-zA-Z0-9]+)*$"`
	// The timestamp when the token was created.
	CreatedAt *string `json:"created_at,omitempty" required:"true"`
	// The APNs device token to register with MagicBell.
	DeviceToken *string `json:"device_token,omitempty" required:"true" minLength:"64"`
	// The timestamp when the token was discarded, if applicable.
	DiscardedAt *util.Nullable[string] `json:"discarded_at,omitempty"`
	// The unique identifier for the token.
	Id *string `json:"id,omitempty" required:"true"`
	// The APNs environment this token belongs to. If omitted we assume it targets `production`.
	InstallationId *InstallationId `json:"installation_id,omitempty"`
	// The timestamp when the token metadata last changed.
	UpdatedAt *util.Nullable[string] `json:"updated_at,omitempty"`
}

func (a *ApnsToken) GetAppId() *string {
	if a == nil {
		return nil
	}
	return a.AppId
}

func (a *ApnsToken) SetAppId(appId string) {
	a.AppId = &appId
}

func (a *ApnsToken) GetCreatedAt() *string {
	if a == nil {
		return nil
	}
	return a.CreatedAt
}

func (a *ApnsToken) SetCreatedAt(createdAt string) {
	a.CreatedAt = &createdAt
}

func (a *ApnsToken) GetDeviceToken() *string {
	if a == nil {
		return nil
	}
	return a.DeviceToken
}

func (a *ApnsToken) SetDeviceToken(deviceToken string) {
	a.DeviceToken = &deviceToken
}

func (a *ApnsToken) GetDiscardedAt() *util.Nullable[string] {
	if a == nil {
		return nil
	}
	return a.DiscardedAt
}

func (a *ApnsToken) SetDiscardedAt(discardedAt util.Nullable[string]) {
	a.DiscardedAt = &discardedAt
}

func (a *ApnsToken) SetDiscardedAtNull() {
	a.DiscardedAt = &util.Nullable[string]{IsNull: true}
}

func (a *ApnsToken) GetId() *string {
	if a == nil {
		return nil
	}
	return a.Id
}

func (a *ApnsToken) SetId(id string) {
	a.Id = &id
}

func (a *ApnsToken) GetInstallationId() *InstallationId {
	if a == nil {
		return nil
	}
	return a.InstallationId
}

func (a *ApnsToken) SetInstallationId(installationId InstallationId) {
	a.InstallationId = &installationId
}

func (a *ApnsToken) GetUpdatedAt() *util.Nullable[string] {
	if a == nil {
		return nil
	}
	return a.UpdatedAt
}

func (a *ApnsToken) SetUpdatedAt(updatedAt util.Nullable[string]) {
	a.UpdatedAt = &updatedAt
}

func (a *ApnsToken) SetUpdatedAtNull() {
	a.UpdatedAt = &util.Nullable[string]{IsNull: true}
}

func (a ApnsToken) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: ApnsToken to string"
	}
	return string(jsonData)
}

func (a *ApnsToken) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}

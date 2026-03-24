package channels

import "encoding/json"

type ApnsTokenPayload struct {
	// The bundle identifier of the application registering this token. Use this to override the default identifier configured on the APNs integration.
	AppId *string `json:"app_id,omitempty" pattern:"^[a-zA-Z0-9]+(.[a-zA-Z0-9]+)*$"`
	// The APNs device token to register with MagicBell.
	DeviceToken *string `json:"device_token,omitempty" required:"true" minLength:"64"`
	// The APNs environment this token belongs to. If omitted we assume it targets `production`.
	InstallationId *ApnsTokenPayloadInstallationId `json:"installation_id,omitempty"`
}

func (a *ApnsTokenPayload) GetAppId() *string {
	if a == nil {
		return nil
	}
	return a.AppId
}

func (a *ApnsTokenPayload) SetAppId(appId string) {
	a.AppId = &appId
}

func (a *ApnsTokenPayload) GetDeviceToken() *string {
	if a == nil {
		return nil
	}
	return a.DeviceToken
}

func (a *ApnsTokenPayload) SetDeviceToken(deviceToken string) {
	a.DeviceToken = &deviceToken
}

func (a *ApnsTokenPayload) GetInstallationId() *ApnsTokenPayloadInstallationId {
	if a == nil {
		return nil
	}
	return a.InstallationId
}

func (a *ApnsTokenPayload) SetInstallationId(installationId ApnsTokenPayloadInstallationId) {
	a.InstallationId = &installationId
}

func (a ApnsTokenPayload) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: ApnsTokenPayload to string"
	}
	return string(jsonData)
}

package channels

import "encoding/json"

type FcmTokenPayload struct {
	// The Firebase Cloud Messaging device registration token to associate with the user.
	DeviceToken *string `json:"device_token,omitempty" required:"true" minLength:"64"`
}

func (f *FcmTokenPayload) GetDeviceToken() *string {
	if f == nil {
		return nil
	}
	return f.DeviceToken
}

func (f *FcmTokenPayload) SetDeviceToken(deviceToken string) {
	f.DeviceToken = &deviceToken
}

func (f FcmTokenPayload) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FcmTokenPayload to string"
	}
	return string(jsonData)
}

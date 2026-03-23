package broadcasts

import "encoding/json"

// Overrides that are scoped to individual delivery channels.
type OverridesChannels struct {
	// Overrides for email notifications.
	Email *Email `json:"email,omitempty"`
	// Overrides for in-app notifications.
	InApp *InApp `json:"in_app,omitempty"`
	// Overrides for mobile push notifications.
	MobilePush *MobilePush `json:"mobile_push,omitempty"`
	// Overrides for SMS notifications.
	Sms *Sms `json:"sms,omitempty"`
}

func (o *OverridesChannels) GetEmail() *Email {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *OverridesChannels) SetEmail(email Email) {
	o.Email = &email
}

func (o *OverridesChannels) GetInApp() *InApp {
	if o == nil {
		return nil
	}
	return o.InApp
}

func (o *OverridesChannels) SetInApp(inApp InApp) {
	o.InApp = &inApp
}

func (o *OverridesChannels) GetMobilePush() *MobilePush {
	if o == nil {
		return nil
	}
	return o.MobilePush
}

func (o *OverridesChannels) SetMobilePush(mobilePush MobilePush) {
	o.MobilePush = &mobilePush
}

func (o *OverridesChannels) GetSms() *Sms {
	if o == nil {
		return nil
	}
	return o.Sms
}

func (o *OverridesChannels) SetSms(sms Sms) {
	o.Sms = &sms
}

func (o OverridesChannels) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OverridesChannels to string"
	}
	return string(jsonData)
}

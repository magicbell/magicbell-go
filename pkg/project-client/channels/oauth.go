package channels

import "encoding/json"

type Oauth struct {
	// The ID of the Slack channel this installation is associated with
	ChannelId *string `json:"channel_id,omitempty" required:"true"`
	// A unique identifier for this Slack workspace installation
	InstallationId *string `json:"installation_id,omitempty" required:"true"`
	// The OAuth scope granted during installation
	Scope *string `json:"scope,omitempty"`
}

func (o *Oauth) GetChannelId() *string {
	if o == nil {
		return nil
	}
	return o.ChannelId
}

func (o *Oauth) SetChannelId(channelId string) {
	o.ChannelId = &channelId
}

func (o *Oauth) GetInstallationId() *string {
	if o == nil {
		return nil
	}
	return o.InstallationId
}

func (o *Oauth) SetInstallationId(installationId string) {
	o.InstallationId = &installationId
}

func (o *Oauth) GetScope() *string {
	if o == nil {
		return nil
	}
	return o.Scope
}

func (o *Oauth) SetScope(scope string) {
	o.Scope = &scope
}

func (o Oauth) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: Oauth to string"
	}
	return string(jsonData)
}

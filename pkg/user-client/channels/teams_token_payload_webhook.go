package channels

import "encoding/json"

type TeamsTokenPayloadWebhook struct {
	Url *string `json:"url,omitempty"`
}

func (t *TeamsTokenPayloadWebhook) GetUrl() *string {
	if t == nil {
		return nil
	}
	return t.Url
}

func (t *TeamsTokenPayloadWebhook) SetUrl(url string) {
	t.Url = &url
}

func (t TeamsTokenPayloadWebhook) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TeamsTokenPayloadWebhook to string"
	}
	return string(jsonData)
}

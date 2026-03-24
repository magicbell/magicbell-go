package channels

import "encoding/json"

type TeamsTokenPayload struct {
	Webhook *TeamsTokenPayloadWebhook `json:"webhook,omitempty"`
}

func (t *TeamsTokenPayload) GetWebhook() *TeamsTokenPayloadWebhook {
	if t == nil {
		return nil
	}
	return t.Webhook
}

func (t *TeamsTokenPayload) SetWebhook(webhook TeamsTokenPayloadWebhook) {
	t.Webhook = &webhook
}

func (t TeamsTokenPayload) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TeamsTokenPayload to string"
	}
	return string(jsonData)
}

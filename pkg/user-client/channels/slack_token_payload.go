package channels

import "encoding/json"

type SlackTokenPayload struct {
	Oauth *SlackTokenPayloadOauth `json:"oauth,omitempty"`
	// Obtained directly from the incoming_webhook object in the installation response from the Slack API.
	Webhook *SlackTokenPayloadWebhook `json:"webhook,omitempty"`
}

func (s *SlackTokenPayload) GetOauth() *SlackTokenPayloadOauth {
	if s == nil {
		return nil
	}
	return s.Oauth
}

func (s *SlackTokenPayload) SetOauth(oauth SlackTokenPayloadOauth) {
	s.Oauth = &oauth
}

func (s *SlackTokenPayload) GetWebhook() *SlackTokenPayloadWebhook {
	if s == nil {
		return nil
	}
	return s.Webhook
}

func (s *SlackTokenPayload) SetWebhook(webhook SlackTokenPayloadWebhook) {
	s.Webhook = &webhook
}

func (s SlackTokenPayload) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SlackTokenPayload to string"
	}
	return string(jsonData)
}

package channels

import "encoding/json"

// Obtained directly from the incoming_webhook object in the installation response from the Slack API.
type SlackTokenWebhook struct {
	// The URL for the incoming webhook from Slack
	Url *string `json:"url,omitempty" required:"true" minLength:"1"`
}

func (s *SlackTokenWebhook) GetUrl() *string {
	if s == nil {
		return nil
	}
	return s.Url
}

func (s *SlackTokenWebhook) SetUrl(url string) {
	s.Url = &url
}

func (s SlackTokenWebhook) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SlackTokenWebhook to string"
	}
	return string(jsonData)
}

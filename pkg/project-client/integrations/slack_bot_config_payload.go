package integrations

import "encoding/json"

type SlackBotConfigPayload struct {
	Enabled *bool `json:"enabled,omitempty" required:"true"`
}

func (s *SlackBotConfigPayload) GetEnabled() *bool {
	if s == nil {
		return nil
	}
	return s.Enabled
}

func (s *SlackBotConfigPayload) SetEnabled(enabled bool) {
	s.Enabled = &enabled
}

func (s SlackBotConfigPayload) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SlackBotConfigPayload to string"
	}
	return string(jsonData)
}

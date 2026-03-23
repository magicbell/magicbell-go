package integrations

import "encoding/json"

type SendgridConfigPayload struct {
	// The API key for Sendgrid
	ApiKey  *string                       `json:"api_key,omitempty" required:"true"`
	From    *SendgridConfigPayloadFrom    `json:"from,omitempty"`
	ReplyTo *SendgridConfigPayloadReplyTo `json:"reply_to,omitempty"`
}

func (s *SendgridConfigPayload) GetApiKey() *string {
	if s == nil {
		return nil
	}
	return s.ApiKey
}

func (s *SendgridConfigPayload) SetApiKey(apiKey string) {
	s.ApiKey = &apiKey
}

func (s *SendgridConfigPayload) GetFrom() *SendgridConfigPayloadFrom {
	if s == nil {
		return nil
	}
	return s.From
}

func (s *SendgridConfigPayload) SetFrom(from SendgridConfigPayloadFrom) {
	s.From = &from
}

func (s *SendgridConfigPayload) GetReplyTo() *SendgridConfigPayloadReplyTo {
	if s == nil {
		return nil
	}
	return s.ReplyTo
}

func (s *SendgridConfigPayload) SetReplyTo(replyTo SendgridConfigPayloadReplyTo) {
	s.ReplyTo = &replyTo
}

func (s SendgridConfigPayload) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SendgridConfigPayload to string"
	}
	return string(jsonData)
}

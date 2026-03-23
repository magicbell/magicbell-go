package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/unmarshal"
	"github.com/magicbell/magicbell-go/pkg/project-client/util"
)

type SendgridConfigPayloadFrom struct {
	// The email address to send from
	Email *string `json:"email,omitempty" required:"true"`
	// The name to send from
	Name *util.Nullable[string] `json:"name,omitempty"`
}

func (s *SendgridConfigPayloadFrom) GetEmail() *string {
	if s == nil {
		return nil
	}
	return s.Email
}

func (s *SendgridConfigPayloadFrom) SetEmail(email string) {
	s.Email = &email
}

func (s *SendgridConfigPayloadFrom) GetName() *util.Nullable[string] {
	if s == nil {
		return nil
	}
	return s.Name
}

func (s *SendgridConfigPayloadFrom) SetName(name util.Nullable[string]) {
	s.Name = &name
}

func (s *SendgridConfigPayloadFrom) SetNameNull() {
	s.Name = &util.Nullable[string]{IsNull: true}
}

func (s SendgridConfigPayloadFrom) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SendgridConfigPayloadFrom to string"
	}
	return string(jsonData)
}

func (s *SendgridConfigPayloadFrom) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, s)
}

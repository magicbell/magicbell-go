package channels

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/user-client/internal/unmarshal"
	"github.com/magicbell/magicbell-go/pkg/user-client/util"
)

type Categories struct {
	Channels []Channels             `json:"channels,omitempty"`
	Key      *string                `json:"key,omitempty" maxLength:"255" pattern:"^[A-Za-z0-9_\.\-/:]+$"`
	Label    *util.Nullable[string] `json:"label,omitempty" maxLength:"255"`
}

func (c *Categories) GetChannels() []Channels {
	if c == nil {
		return nil
	}
	return c.Channels
}

func (c *Categories) SetChannels(channels []Channels) {
	c.Channels = channels
}

func (c *Categories) GetKey() *string {
	if c == nil {
		return nil
	}
	return c.Key
}

func (c *Categories) SetKey(key string) {
	c.Key = &key
}

func (c *Categories) GetLabel() *util.Nullable[string] {
	if c == nil {
		return nil
	}
	return c.Label
}

func (c *Categories) SetLabel(label util.Nullable[string]) {
	c.Label = &label
}

func (c *Categories) SetLabelNull() {
	c.Label = &util.Nullable[string]{IsNull: true}
}

func (c Categories) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: Categories to string"
	}
	return string(jsonData)
}

func (c *Categories) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}

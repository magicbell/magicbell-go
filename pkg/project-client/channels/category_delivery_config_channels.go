package channels

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/unmarshal"
	"github.com/magicbell/magicbell-go/pkg/project-client/util"
)

type CategoryDeliveryConfigChannels struct {
	// Name of the channel used for this step.
	Channel *Channel `json:"channel,omitempty" required:"true"`
	// Delay in seconds to wait after the previous step.
	Delay *int64 `json:"delay,omitempty" min:"0"`
	// Conditional expression evaluated before running the step.
	If_ *util.Nullable[string] `json:"if,omitempty"`
}

func (c *CategoryDeliveryConfigChannels) GetChannel() *Channel {
	if c == nil {
		return nil
	}
	return c.Channel
}

func (c *CategoryDeliveryConfigChannels) SetChannel(channel Channel) {
	c.Channel = &channel
}

func (c *CategoryDeliveryConfigChannels) GetDelay() *int64 {
	if c == nil {
		return nil
	}
	return c.Delay
}

func (c *CategoryDeliveryConfigChannels) SetDelay(delay int64) {
	c.Delay = &delay
}

func (c *CategoryDeliveryConfigChannels) GetIf_() *util.Nullable[string] {
	if c == nil {
		return nil
	}
	return c.If_
}

func (c *CategoryDeliveryConfigChannels) SetIf_(if_ util.Nullable[string]) {
	c.If_ = &if_
}

func (c *CategoryDeliveryConfigChannels) SetIf_Null() {
	c.If_ = &util.Nullable[string]{IsNull: true}
}

func (c CategoryDeliveryConfigChannels) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CategoryDeliveryConfigChannels to string"
	}
	return string(jsonData)
}

func (c *CategoryDeliveryConfigChannels) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}

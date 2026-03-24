package events

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/unmarshal"
	"github.com/magicbell/magicbell-go/pkg/project-client/util"
)

type Event struct {
	// The numeric code that categorizes the event.
	Code *int64 `json:"code,omitempty"`
	// Additional contextual attributes for the event.
	Context *util.Nullable[any] `json:"context,omitempty"`
	// The unique identifier for the event.
	Id *string `json:"id,omitempty" required:"true"`
	// The severity level assigned to the event.
	Level *string `json:"level,omitempty"`
	// A human-readable log message.
	Log *util.Nullable[string] `json:"log,omitempty"`
	// The raw payload delivered by the event source.
	Payload *util.Nullable[any] `json:"payload,omitempty"`
	// The time at which the event was recorded.
	Timestamp *string `json:"timestamp,omitempty" required:"true"`
	// The type of event that occurred.
	Type_ *string `json:"type,omitempty" required:"true"`
}

func (e *Event) GetCode() *int64 {
	if e == nil {
		return nil
	}
	return e.Code
}

func (e *Event) SetCode(code int64) {
	e.Code = &code
}

func (e *Event) GetContext() *util.Nullable[any] {
	if e == nil {
		return nil
	}
	return e.Context
}

func (e *Event) SetContext(context util.Nullable[any]) {
	e.Context = &context
}

func (e *Event) SetContextNull() {
	e.Context = &util.Nullable[any]{IsNull: true}
}

func (e *Event) GetId() *string {
	if e == nil {
		return nil
	}
	return e.Id
}

func (e *Event) SetId(id string) {
	e.Id = &id
}

func (e *Event) GetLevel() *string {
	if e == nil {
		return nil
	}
	return e.Level
}

func (e *Event) SetLevel(level string) {
	e.Level = &level
}

func (e *Event) GetLog() *util.Nullable[string] {
	if e == nil {
		return nil
	}
	return e.Log
}

func (e *Event) SetLog(log util.Nullable[string]) {
	e.Log = &log
}

func (e *Event) SetLogNull() {
	e.Log = &util.Nullable[string]{IsNull: true}
}

func (e *Event) GetPayload() *util.Nullable[any] {
	if e == nil {
		return nil
	}
	return e.Payload
}

func (e *Event) SetPayload(payload util.Nullable[any]) {
	e.Payload = &payload
}

func (e *Event) SetPayloadNull() {
	e.Payload = &util.Nullable[any]{IsNull: true}
}

func (e *Event) GetTimestamp() *string {
	if e == nil {
		return nil
	}
	return e.Timestamp
}

func (e *Event) SetTimestamp(timestamp string) {
	e.Timestamp = &timestamp
}

func (e *Event) GetType_() *string {
	if e == nil {
		return nil
	}
	return e.Type_
}

func (e *Event) SetType_(type_ string) {
	e.Type_ = &type_
}

func (e Event) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: Event to string"
	}
	return string(jsonData)
}

func (e *Event) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, e)
}

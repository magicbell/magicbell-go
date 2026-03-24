package broadcasts

import "encoding/json"

type Errors struct {
	// The details about the processing error.
	Message *string `json:"message,omitempty"`
}

func (e *Errors) GetMessage() *string {
	if e == nil {
		return nil
	}
	return e.Message
}

func (e *Errors) SetMessage(message string) {
	e.Message = &message
}

func (e Errors) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: Errors to string"
	}
	return string(jsonData)
}

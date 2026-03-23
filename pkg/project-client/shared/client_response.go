package shared

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest/httptransport"
	"net/http"
)

// ClientResponse is the user-facing wrapper for API responses.
// It contains the deserialized data, raw HTTP response, and metadata like headers and status code.
type ClientResponse[T any] struct {
	Data     T
	Raw      *http.Response
	Metadata ClientResponseMetadata
}

// ClientResponseMetadata contains HTTP metadata from the API response.
// Includes status code and headers for inspection and debugging.
type ClientResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

// NewClientResponse creates a new response wrapper from an internal transport response.
// Extracts data and metadata into a user-facing structure.
func NewClientResponse[T any](resp *httptransport.Response[T]) *ClientResponse[T] {
	return &ClientResponse[T]{
		Data: resp.Data,
		Raw:  resp.Raw,
		Metadata: ClientResponseMetadata{
			StatusCode: resp.StatusCode,
			Headers:    resp.Headers,
		},
	}
}

// GetData returns the deserialized response data.
func (r *ClientResponse[T]) GetData() T {
	return r.Data
}

// String returns a JSON representation of the response for debugging.
// Returns an error message if JSON marshaling fails.
func (r ClientResponse[T]) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ClientResponse to string"
	}
	return string(jsonData)
}

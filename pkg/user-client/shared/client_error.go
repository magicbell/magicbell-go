package shared

import (
	"github.com/magicbell/magicbell-go/pkg/user-client/internal/clients/rest/httptransport"
	"net/http"
)

// ClientError wraps API errors with detailed metadata including status code, headers, and raw response.
// It implements the error interface and provides structured access to error information.
type ClientError[T any] struct {
	Err      error
	Data     *T
	Body     []byte
	Raw      *http.Response
	Metadata ClientErrorMetadata
}

// ClientErrorMetadata contains HTTP metadata associated with an error response.
type ClientErrorMetadata struct {
	Headers    map[string]string
	StatusCode int
}

// NewClientError creates a new ClientError from an internal transport error.
// It extracts error details, body, status code, and headers into a user-facing error structure.
func NewClientError[T any](transportError *httptransport.ErrorResponse[T]) *ClientError[T] {
	return &ClientError[T]{
		Err:  transportError.GetError(),
		Data: transportError.Data,
		Body: transportError.GetBody(),
		Raw:  transportError.Raw,
		Metadata: ClientErrorMetadata{
			StatusCode: transportError.GetStatusCode(),
			Headers:    transportError.GetHeaders(),
		},
	}
}

// Error implements the error interface, returning the error message string.
func (e *ClientError[T]) Error() string {
	return e.Err.Error()
}

// GetData returns the deserialized error response data.
// Returns nil if unmarshaling failed or the response body was empty.
func (e *ClientError[T]) GetData() *T {
	return e.Data
}

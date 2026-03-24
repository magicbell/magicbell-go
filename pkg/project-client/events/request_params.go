package events

// ListEventsRequestParams holds the optional parameters for the API request.
// Use the Set methods to configure query parameters, headers, and path parameters.
type ListEventsRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

// SetLimit sets the Limit parameter.
func (params *ListEventsRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}

// SetStartingAfter sets the StartingAfter parameter.
func (params *ListEventsRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}

// SetEndingBefore sets the EndingBefore parameter.
func (params *ListEventsRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

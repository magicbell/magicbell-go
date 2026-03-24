package broadcasts

// ListBroadcastsRequestParams holds the optional parameters for the API request.
// Use the Set methods to configure query parameters, headers, and path parameters.
type ListBroadcastsRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

// SetLimit sets the Limit parameter.
func (params *ListBroadcastsRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}

// SetStartingAfter sets the StartingAfter parameter.
func (params *ListBroadcastsRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}

// SetEndingBefore sets the EndingBefore parameter.
func (params *ListBroadcastsRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

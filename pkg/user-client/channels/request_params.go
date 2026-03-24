package channels

// ListInboxTokensRequestParams holds the optional parameters for the API request.
// Use the Set methods to configure query parameters, headers, and path parameters.
type ListInboxTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

// SetLimit sets the Limit parameter.
func (params *ListInboxTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}

// SetStartingAfter sets the StartingAfter parameter.
func (params *ListInboxTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}

// SetEndingBefore sets the EndingBefore parameter.
func (params *ListInboxTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

// ListApnsTokensRequestParams holds the optional parameters for the API request.
// Use the Set methods to configure query parameters, headers, and path parameters.
type ListApnsTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

// SetLimit sets the Limit parameter.
func (params *ListApnsTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}

// SetStartingAfter sets the StartingAfter parameter.
func (params *ListApnsTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}

// SetEndingBefore sets the EndingBefore parameter.
func (params *ListApnsTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

// ListExpoTokensRequestParams holds the optional parameters for the API request.
// Use the Set methods to configure query parameters, headers, and path parameters.
type ListExpoTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

// SetLimit sets the Limit parameter.
func (params *ListExpoTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}

// SetStartingAfter sets the StartingAfter parameter.
func (params *ListExpoTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}

// SetEndingBefore sets the EndingBefore parameter.
func (params *ListExpoTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

// ListFcmTokensRequestParams holds the optional parameters for the API request.
// Use the Set methods to configure query parameters, headers, and path parameters.
type ListFcmTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

// SetLimit sets the Limit parameter.
func (params *ListFcmTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}

// SetStartingAfter sets the StartingAfter parameter.
func (params *ListFcmTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}

// SetEndingBefore sets the EndingBefore parameter.
func (params *ListFcmTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

// ListMagicbellSlackbotTokensRequestParams holds the optional parameters for the API request.
// Use the Set methods to configure query parameters, headers, and path parameters.
type ListMagicbellSlackbotTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

// SetLimit sets the Limit parameter.
func (params *ListMagicbellSlackbotTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}

// SetStartingAfter sets the StartingAfter parameter.
func (params *ListMagicbellSlackbotTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}

// SetEndingBefore sets the EndingBefore parameter.
func (params *ListMagicbellSlackbotTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

// ListSlackTokensRequestParams holds the optional parameters for the API request.
// Use the Set methods to configure query parameters, headers, and path parameters.
type ListSlackTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

// SetLimit sets the Limit parameter.
func (params *ListSlackTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}

// SetStartingAfter sets the StartingAfter parameter.
func (params *ListSlackTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}

// SetEndingBefore sets the EndingBefore parameter.
func (params *ListSlackTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

// ListTeamsTokensRequestParams holds the optional parameters for the API request.
// Use the Set methods to configure query parameters, headers, and path parameters.
type ListTeamsTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

// SetLimit sets the Limit parameter.
func (params *ListTeamsTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}

// SetStartingAfter sets the StartingAfter parameter.
func (params *ListTeamsTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}

// SetEndingBefore sets the EndingBefore parameter.
func (params *ListTeamsTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

// ListWebPushTokensRequestParams holds the optional parameters for the API request.
// Use the Set methods to configure query parameters, headers, and path parameters.
type ListWebPushTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

// SetLimit sets the Limit parameter.
func (params *ListWebPushTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}

// SetStartingAfter sets the StartingAfter parameter.
func (params *ListWebPushTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}

// SetEndingBefore sets the EndingBefore parameter.
func (params *ListWebPushTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

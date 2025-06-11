package channels

type ListInboxTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *ListInboxTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *ListInboxTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *ListInboxTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type ListApnsTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *ListApnsTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *ListApnsTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *ListApnsTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type ListExpoTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *ListExpoTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *ListExpoTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *ListExpoTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type ListFcmTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *ListFcmTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *ListFcmTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *ListFcmTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type ListSlackTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *ListSlackTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *ListSlackTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *ListSlackTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type ListTeamsTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *ListTeamsTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *ListTeamsTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *ListTeamsTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type ListWebPushTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *ListWebPushTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *ListWebPushTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *ListWebPushTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

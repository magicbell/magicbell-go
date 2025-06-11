package channels

type ListUserInboxTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *ListUserInboxTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *ListUserInboxTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *ListUserInboxTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type ListUserApnsTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *ListUserApnsTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *ListUserApnsTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *ListUserApnsTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type ListUserExpoTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *ListUserExpoTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *ListUserExpoTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *ListUserExpoTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type ListUserFcmTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *ListUserFcmTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *ListUserFcmTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *ListUserFcmTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type ListUserSlackTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *ListUserSlackTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *ListUserSlackTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *ListUserSlackTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type ListUserTeamsTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *ListUserTeamsTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *ListUserTeamsTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *ListUserTeamsTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type ListUserWebPushTokensRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *ListUserWebPushTokensRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *ListUserWebPushTokensRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *ListUserWebPushTokensRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

package integrations

type ListIntegrationsRequestParams struct {
	Limit         *int64  `explode:"true" serializationStyle:"form" queryParam:"limit"`
	StartingAfter *string `explode:"true" serializationStyle:"form" queryParam:"starting_after"`
	EndingBefore  *string `explode:"true" serializationStyle:"form" queryParam:"ending_before"`
}

func (params *ListIntegrationsRequestParams) SetLimit(limit int64) {
	params.Limit = &limit
}
func (params *ListIntegrationsRequestParams) SetStartingAfter(startingAfter string) {
	params.StartingAfter = &startingAfter
}
func (params *ListIntegrationsRequestParams) SetEndingBefore(endingBefore string) {
	params.EndingBefore = &endingBefore
}

type DeleteApnsIntegrationRequestParams struct {
	Id *string `explode:"true" serializationStyle:"form" queryParam:"id"`
}

func (params *DeleteApnsIntegrationRequestParams) SetId(id string) {
	params.Id = &id
}

type DeleteExpoIntegrationRequestParams struct {
	Id *string `explode:"true" serializationStyle:"form" queryParam:"id"`
}

func (params *DeleteExpoIntegrationRequestParams) SetId(id string) {
	params.Id = &id
}

type DeleteFcmIntegrationRequestParams struct {
	Id *string `explode:"true" serializationStyle:"form" queryParam:"id"`
}

func (params *DeleteFcmIntegrationRequestParams) SetId(id string) {
	params.Id = &id
}

type DeleteGithubIntegrationRequestParams struct {
	Id *string `explode:"true" serializationStyle:"form" queryParam:"id"`
}

func (params *DeleteGithubIntegrationRequestParams) SetId(id string) {
	params.Id = &id
}

type DeleteInboxIntegrationRequestParams struct {
	Id *string `explode:"true" serializationStyle:"form" queryParam:"id"`
}

func (params *DeleteInboxIntegrationRequestParams) SetId(id string) {
	params.Id = &id
}

type DeleteMailgunIntegrationRequestParams struct {
	Id *string `explode:"true" serializationStyle:"form" queryParam:"id"`
}

func (params *DeleteMailgunIntegrationRequestParams) SetId(id string) {
	params.Id = &id
}

type DeletePingEmailIntegrationRequestParams struct {
	Id *string `explode:"true" serializationStyle:"form" queryParam:"id"`
}

func (params *DeletePingEmailIntegrationRequestParams) SetId(id string) {
	params.Id = &id
}

type DeleteSendgridIntegrationRequestParams struct {
	Id *string `explode:"true" serializationStyle:"form" queryParam:"id"`
}

func (params *DeleteSendgridIntegrationRequestParams) SetId(id string) {
	params.Id = &id
}

type DeleteSesIntegrationRequestParams struct {
	Id *string `explode:"true" serializationStyle:"form" queryParam:"id"`
}

func (params *DeleteSesIntegrationRequestParams) SetId(id string) {
	params.Id = &id
}

type DeleteSlackIntegrationRequestParams struct {
	Id *string `explode:"true" serializationStyle:"form" queryParam:"id"`
}

func (params *DeleteSlackIntegrationRequestParams) SetId(id string) {
	params.Id = &id
}

type DeleteStripeIntegrationRequestParams struct {
	Id *string `explode:"true" serializationStyle:"form" queryParam:"id"`
}

func (params *DeleteStripeIntegrationRequestParams) SetId(id string) {
	params.Id = &id
}

type DeleteTwilioIntegrationRequestParams struct {
	Id *string `explode:"true" serializationStyle:"form" queryParam:"id"`
}

func (params *DeleteTwilioIntegrationRequestParams) SetId(id string) {
	params.Id = &id
}

type DeleteWebPushIntegrationRequestParams struct {
	Id *string `explode:"true" serializationStyle:"form" queryParam:"id"`
}

func (params *DeleteWebPushIntegrationRequestParams) SetId(id string) {
	params.Id = &id
}

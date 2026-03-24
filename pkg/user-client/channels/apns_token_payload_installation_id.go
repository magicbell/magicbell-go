package channels

// The APNs environment this token belongs to. If omitted we assume it targets `production`.
type ApnsTokenPayloadInstallationId string

const (
	APNS_TOKEN_PAYLOAD_INSTALLATION_ID_DEVELOPMENT ApnsTokenPayloadInstallationId = "development"
	APNS_TOKEN_PAYLOAD_INSTALLATION_ID_PRODUCTION  ApnsTokenPayloadInstallationId = "production"
)

package channels

// The APNs environment this token belongs to. If omitted we assume it targets `production`.
type ApnsTokenInstallationId string

const (
	APNS_TOKEN_INSTALLATION_ID_DEVELOPMENT ApnsTokenInstallationId = "development"
	APNS_TOKEN_INSTALLATION_ID_PRODUCTION  ApnsTokenInstallationId = "production"
)

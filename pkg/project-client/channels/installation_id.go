package channels

// The APNs environment this token belongs to. If omitted we assume it targets `production`.
type InstallationId string

const (
	INSTALLATION_ID_DEVELOPMENT InstallationId = "development"
	INSTALLATION_ID_PRODUCTION  InstallationId = "production"
)

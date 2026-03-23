package integrations

// SMTP security/encryption method
type Security string

const (
	SECURITY_NONE     Security = "none"
	SECURITY_SSL      Security = "ssl"
	SECURITY_STARTTLS Security = "starttls"
)

package integrations

// Controls whether the app icon badge counts unread or unseen notifications.
type Badge string

const (
	BADGE_UNREAD Badge = "unread"
	BADGE_UNSEEN Badge = "unseen"
)

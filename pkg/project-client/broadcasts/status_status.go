package broadcasts

// The overall processing status of the broadcast.
type StatusStatus string

const (
	STATUS_STATUS_ENQUEUED   StatusStatus = "enqueued"
	STATUS_STATUS_PROCESSING StatusStatus = "processing"
	STATUS_STATUS_PROCESSED  StatusStatus = "processed"
)

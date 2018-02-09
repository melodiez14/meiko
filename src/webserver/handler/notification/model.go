package notification

type subscribeParams struct {
	playerID string
}

type subscribeArgs struct {
	playerID string
}

type getNotificationParam struct {
	page string
}

type getNotificationArgs struct {
	page uint16
}

type Notification struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	ReadAt      string `json:"read_at,omitempty"`
	CreatedAt   int64  `json:"created_at"`
}

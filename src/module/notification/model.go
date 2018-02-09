package notification

type Notification struct {
	ID          int64  `db:"id"`
	UserID      int64  `db:"users_id"`
	OneSignalID string `db:"onesignal_id"`
}

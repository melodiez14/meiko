package assignment

type summaryResponse struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Status int8   `json:"status,omitempty"`
}

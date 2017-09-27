package information

type informationResponse struct {
	Title string `json:"title"`
	Date  string `json:"date"`
}

type getSummaryResponse struct {
	Last   []informationResponse `json:"last"`
	Recent []informationResponse `json:"recent"`
}

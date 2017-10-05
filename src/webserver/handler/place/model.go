package place

type searchParams struct {
	Query string
}

type searchArgs struct {
	Query string
}

type searchResponse struct {
	ID []string `json:"places"`
}

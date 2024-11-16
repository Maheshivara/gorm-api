package response

type Pagination struct {
	PerPage int
	Page    int
}

type SearchResult struct {
	PerPage int   `json:"perPage"`
	Page    int   `json:"page"`
	Total   int64 `json:"total"`
	Data    any   `json:"data"`
}

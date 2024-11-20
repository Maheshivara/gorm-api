package response

import "gormCompose/src/models"

type Pagination struct {
	PerPage int `example:"20"`
	Page    int `example:"1"`
}

type SearchResult struct {
	PerPage int            `json:"perPage" example:"20"`
	Page    int            `json:"page" example:"1"`
	Total   int64          `json:"total" example:"15"`
	Data    []*models.Food `json:"data"`
}

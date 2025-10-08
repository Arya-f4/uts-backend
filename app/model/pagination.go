
package model

type PaginationParams struct {
	Page   int
	Limit  int
	Sort   string
	Search string
}

type PaginationResult[T any] struct {
	Data     []T   `json:"data"`
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	Limit    int   `json:"limit"`
	LastPage int   `json:"last_page"`
}


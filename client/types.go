package client

type listResponse[T any] struct {
	Pagination struct {
		CurrentPage int `json:"current_page"`
		TotalPage   int `json:"total_page"`
		TotalCount  int `json:"total_count"`
	} `json:"pagination"`
	Data []T `json:"data"`
}

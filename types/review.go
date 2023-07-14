package types

type Review struct {
	Id          string `json:"id"`
	User        string `json:"user"`
	Product     string `json:"product"`
	Stars       int    `json:"stars"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

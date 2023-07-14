package types

type ProductVisibility string

const (
	ProductVisibilityVisible   ProductVisibility = "visible"
	ProductVisibilityHidden    ProductVisibility = "hidden"
	ProductVisibilityArchived  ProductVisibility = "archived"
	ProductVisibilityQuickLink ProductVisibility = "quick_link"
)

type Product struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Visibility  ProductVisibility `json:"visibility"`
	CreatedAt   int               `json:"created_at"`
	Experiences []string          `json:"experiences"`
	Plans       []string          `json:"plans"`
}

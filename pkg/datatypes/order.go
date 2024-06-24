package datatypes

type Pair struct {
	OrderBy string   `form:"order_by"`
	SortBy  SortType `form:"sort_by"`
}

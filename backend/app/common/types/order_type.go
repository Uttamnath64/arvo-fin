package commonType

// "asc", "desc"
type OrderType string

const (
	OrderAsc  OrderType = "asc"
	OrderDesc OrderType = "desc"
)

func (o OrderType) IsValid() bool {
	return o == OrderAsc || o == OrderDesc
}

func (o OrderType) String() string {
	if !o.IsValid() {
		return "asc" // default fallback
	}
	return string(o)
}

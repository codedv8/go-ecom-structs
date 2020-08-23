package API

type CartOrder struct {
	State      int8
	OrderItems []CartOrderItem
}

type CartOrderItem struct {
	SKU string
	Qty float64
}

package request

type Food struct {
	Name  string  `json:"name"  binding:"required"`
	Price float32 `json:"price"  binding:"required"`
}

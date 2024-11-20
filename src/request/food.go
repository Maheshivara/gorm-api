package request

type Food struct {
	Name  string  `json:"name"  binding:"required" example:"Arroz com Passas"`
	Price float32 `json:"price"  binding:"required" example:"14.75"`
}

package request

type Food struct {
	Name  string  `json:"name"  binding:"required,min=3" example:"Arroz com Passas"`
	Price float32 `json:"price"  binding:"required,gt=0" example:"14.75"`
}

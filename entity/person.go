package entity

type Person struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Age       int8   `json:"age" binding:"gte=15,lte=130"`
	Email     string `json:"email" validate:"required, email"`
}

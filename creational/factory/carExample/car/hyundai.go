package car

type i30 struct {
	car
}

func NewI30() icar {
	return &i30{car: car{
		NumberOfWheels: 4,
		Color:          "black",
		Make:           "hyundai",
		Model:          "i30",
	}}
}

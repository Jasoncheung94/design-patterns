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

type i40 struct {
	car
	Edition string // additional field for i40 that this make only has
}

func NewI40() icar {
	return &i40{
		car: car{
			NumberOfWheels: 4,
			Color:          "black",
			Make:           "hyundai",
			Model:          "i40",
		},
		Edition: "premium",
	}
}

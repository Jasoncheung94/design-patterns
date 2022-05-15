package car

func Get(carType string) icar {
	switch carType {
	case "i30":
		return NewI30()
	default:
		return nil
	}
}

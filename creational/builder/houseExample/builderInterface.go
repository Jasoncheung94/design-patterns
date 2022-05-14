package main

type Builder interface {
	setWindowType()
	setDoorType()
	setNumFloors()
	getHouse() house
}

func getBuilder(builderType string) Builder {
	switch builderType {
	case "house":
		return &houseBuilder{}
	case "igloo":
		return &iglooBuilder{}
	default:
		return nil
	}
}

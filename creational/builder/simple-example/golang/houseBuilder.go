package main

type houseBuilder struct {
	windowType string
	doorType   string
	numFloors  int
}

func newHouseBuilder() *houseBuilder {
	return &houseBuilder{}
}

func (hb *houseBuilder) setWindowType() {
	hb.windowType = "fixed"
}

func (hb *houseBuilder) setDoorType() {
	hb.doorType = "wooden"
}

func (hb *houseBuilder) setNumFloors() {
	hb.numFloors = 2
}

func (hb *houseBuilder) getHouse() house {
	return house{
		windowType: hb.windowType,
		doorType:   hb.doorType,
		numFloors:  hb.numFloors,
	}
}

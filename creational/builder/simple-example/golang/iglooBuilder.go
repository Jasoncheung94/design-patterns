package main

type iglooBuilder struct {
	windowType string
	doorType   string
	numFloors  int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (ib *iglooBuilder) setWindowType() {
	ib.windowType = "ice"
}

func (ib *iglooBuilder) setDoorType() {
	ib.doorType = "ice"
}

func (ib *iglooBuilder) setNumFloors() {
	ib.numFloors = 1
}

func (ib *iglooBuilder) getHouse() house {
	return house{
		windowType: ib.windowType,
		doorType:   ib.doorType,
		numFloors:  ib.numFloors,
	}
}

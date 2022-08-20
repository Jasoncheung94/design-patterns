package main

import "fmt"

type house struct {
	windowType string
	doorType   string
	numFloors  int
}

func (h house) printDetails() {
	fmt.Printf(
		"Window type: %s\n Door type: %s\n Number of floors: %d\n\n",
		h.windowType,
		h.windowType,
		h.numFloors,
	)
}

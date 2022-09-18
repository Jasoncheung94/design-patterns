package main

import (
	"fmt"
	"time"
)

type Plane interface {
	arrive()
	depart()
	clearForLanding()
}

type Mediator interface {
	canArrive(Plane) bool
	communicateDeparture()
}

type CommercialPlane struct {
	mediator Mediator
}

func (c *CommercialPlane) arrive() {
	if !c.mediator.canArrive(c) {
		fmt.Println("CommercialPlane: Cannot arrive")
		return
	}
	fmt.Println("CommercialPlane: Arrived")
}

func (c *CommercialPlane) depart() {
	fmt.Println("CommercialPlane: Departing")
	c.mediator.communicateDeparture()
}

func (c *CommercialPlane) clearForLanding() {
	fmt.Println("CommercialPlane: Cleared for landing")
	c.arrive()
}

type AirTrafficControl struct {
	runwayClear bool
	planes      []Plane
}

func newAirTrafficControl() *AirTrafficControl {
	return &AirTrafficControl{
		runwayClear: true,
		planes:      []Plane{},
	}
}

func (a *AirTrafficControl) canArrive(plane Plane) bool {
	if a.runwayClear {
		a.runwayClear = false
		return true
	}
	a.planes = append(a.planes, plane)
	return false
}

func (a *AirTrafficControl) communicateDeparture() {
	if !a.runwayClear {
		a.runwayClear = true
	}
	if len(a.planes) > 0 {
		a.planes[0].clearForLanding()
		a.planes = a.planes[1:]
	} else {
		fmt.Println("No planes waiting")
	}
}

func main() {
	atc := newAirTrafficControl()
	plane1 := &CommercialPlane{mediator: atc}
	plane2 := &CommercialPlane{mediator: atc}
	plane3 := &CommercialPlane{mediator: atc}

	plane1.arrive()
	plane2.arrive()
	plane3.arrive()
	time.Sleep(2 * time.Second)
	plane1.depart()
	time.Sleep(2 * time.Second)
	plane2.depart()
	time.Sleep(2 * time.Second)
	plane3.depart()
}

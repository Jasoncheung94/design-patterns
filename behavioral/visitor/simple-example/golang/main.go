package main

import (
	"fmt"
	"strconv"
)

// Element interface
type Element interface {
	Accept(v CarVisitor)
}

// Engine is ConcreteElement
type Engine struct{}

// Accept operation that takes a visitor as an argument
func (e Engine) Accept(v CarVisitor) {
	v.visitEngine(e)
}

// Wheel is ConcreteElement
type Wheel struct {
	Number int
}

// Accept operation
func (w Wheel) Accept(v CarVisitor) {
	v.visitWheel(w)
}

// Car is ConcreteElement
type Car struct {
	items []Element
}

// Accept operation
func (c Car) Accept(v CarVisitor) {
	for _, e := range c.items {
		e.Accept(v)
	}
	v.visitCar(c)
}

// CarVisitor is Visitor
type CarVisitor interface {
	visitEngine(engine Engine)
	visitWheel(wheel Wheel)
	visitCar(car Car)
}

// TestCarVisitor is ConcreteVisitor
type TestCarVisitor struct{}

func (v TestCarVisitor) visitEngine(engine Engine) {
	fmt.Println("test engine")
}

func (v TestCarVisitor) visitWheel(wheel Wheel) {
	fmt.Println("test wheel #" +
		strconv.Itoa(wheel.Number))
}

func (v TestCarVisitor) visitCar(car Car) {
	fmt.Println("test car")
}

// recoveryCarVisitor is ConcreteVisitor
type RecoveryCarVisitor struct{}

func (r RecoveryCarVisitor) visitEngine(engine Engine) {
	fmt.Println("recovery engine")
}

func (r RecoveryCarVisitor) visitWheel(wheel Wheel) {
	fmt.Println("recovery wheel #" +
		strconv.Itoa(wheel.Number))
}

func (r RecoveryCarVisitor) visitCar(car Car) {
	fmt.Println("recovery car")
}

func main() {
	car := Car{[]Element{
		Engine{},
		Wheel{1},
		Wheel{2},
		Wheel{3},
		Wheel{4},
	}}
	testCar := TestCarVisitor{}
	recoveryCar := RecoveryCarVisitor{}

	car.Accept(testCar)
	car.Accept(recoveryCar)
}

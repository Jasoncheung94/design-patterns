package main

import "fmt"

type CarInterface interface {
	setModel(name string)
	getModel() string
	setColor(color string)
	getColor() string
}

type Car struct {
	model string
	color string
}

func (c *Car) setModel(name string) {
	c.model = name
}

func (c *Car) getModel() string {
	return c.model
}

func (c *Car) setColor(color string) {
	c.color = color
}

func (c *Car) getColor() string {
	return c.color
}

type Hyundai struct {
	Car
}

func newHyundaiI30() *Hyundai {
	return &Hyundai{Car{
		model: "i30",
		color: "black",
	}}
}

type Toyota struct {
	Car
}

func newToyotaSupra() *Toyota {
	return &Toyota{Car{
		model: "supra",
		color: "red",
	}}
}

func getCar(model string) CarInterface {
	switch model {
	case "hyundai":
		return newHyundaiI30()
	case "toyota":
		return newToyotaSupra()
	default:
		return nil
	}
}

// GO111MODULE="off" go run main.go
func main() {
	hyundaiCar := getCar("hyundai")
	toyotaCar := getCar("toyota")
	fmt.Println("Model:", hyundaiCar.getModel(), "Color:", hyundaiCar.getColor())
	fmt.Println("Model:", toyotaCar.getModel(), "Color:", toyotaCar.getColor())
}

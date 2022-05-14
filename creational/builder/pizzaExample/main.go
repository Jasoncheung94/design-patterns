package main

import "github.com/jasoncheung94/design-patterns/creational/builder/pizzaExample/pizza"

func main() {
	// GO111MODULE="off" go run main.go
	meaty := pizza.Meaty{}
	director := pizza.NewDirector(&meaty)
	director.Construct()
	meaty.Details()

	veggie := pizza.Veggie{}
	director.SetBuilder(&veggie)
	director.Construct().Details()
}

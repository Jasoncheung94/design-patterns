package main

import (
	"fmt"

	"github.com/jasoncheung94/design-patterns/creational/factory_method/carExample/golang/car"
)

// GO111MODULE="off" go run main.go
func main() {
	i30 := car.Get("i30")
	fmt.Println(i30)

	i40 := car.Get("i40")
	fmt.Println(i40)
}

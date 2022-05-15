package main

import (
	"fmt"

	"github.com/jasoncheung94/design-patterns/creational/factory/carExample/car"
)

// GO111MODULE="off" go run main.go
func main() {
	i30 := car.Get("i30")
	fmt.Println(i30)
}

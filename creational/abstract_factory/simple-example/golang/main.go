package main

import "fmt"

// `go run *.go“ to run the program
func main() {
	adidasFactory, err := GetSportsFactory("Adidas")
	if err != nil {
		panic(err)
	}

	nikeFactory, err := GetSportsFactory("Nike")
	if err != nil {
		panic(err)
	}

	adidasShoe := adidasFactory.makeShoe()
	fmt.Println("Shoe -> Brand:", adidasShoe.getBrand(), "Size:", adidasShoe.getSize())

	nikeShoe := nikeFactory.makeShoe()
	fmt.Println("Shoe -> Brand:", nikeShoe.getBrand(), "Size:", nikeShoe.getSize())

	adidasJersey := adidasFactory.makeJersey()
	fmt.Println("Jersey -> Brand:", adidasJersey.getBrand(), "Size:", adidasJersey.getSize())

	nikeJersey := nikeFactory.makeJersey()
	nikeJersey.setTeam("Liverpool FC")
	fmt.Println("Jersey -> Brand:", nikeJersey.getBrand(), "Size:", nikeJersey.getSize(), "Team:", nikeJersey.getTeam())
}

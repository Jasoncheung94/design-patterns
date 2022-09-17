package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
	SetColour(colour colour)
	Colour()
}

type Circle struct {
	radius float64
	colour colour
}

func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func (c *Circle) SetColour(colour colour) {
	c.colour = colour
}

func (c *Circle) Colour() {
	c.colour.Display()
}

type Square struct {
	side   float64
	colour colour
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

func (s *Square) Perimeter() float64 {
	return 4 * s.side
}

func (s *Square) SetColour(colour colour) {
	s.colour = colour
}

func (s *Square) Colour() {
	s.colour.Display()
}

type colour interface {
	Display()
}

type Red struct{}

func (r *Red) Display() {
	println("Red")
}

type Blue struct{}

func (b *Blue) Display() {
	println("Blue")
}

// Steps 1. Create a Shape interface with two methods: Area() and Perimeter().
// 2. Create two concrete classes: Circle and Square.
// 3. Create a colour interface with a single method: Display().
// 4. Create two concrete classes: Red and Blue.
// 5. Update the Shape interface to include a colour + Setcolour() method.
// 6. Update the Circle and Square classes to implement the colour interface and it's fields
//
// GO111MODULE="off" go run main.go
func main() {
	red := &Red{}
	blue := &Blue{}

	circle := &Circle{radius: 10}
	square := &Square{side: 10}

	circle.SetColour(red)
	square.SetColour(blue)

	fmt.Println("Circle area:", circle.Area(), "perimeter:", circle.Perimeter(), "colour:")
	circle.Colour()

	fmt.Println("Square area:", square.Area(), "perimeter:", square.Perimeter(), "colour:")
	square.Colour()
}

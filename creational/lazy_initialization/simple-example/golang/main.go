package main

import "fmt"

type Fruit struct {
	name string
}

type Fruits struct {
	fruits []Fruit
}

func (f *Fruits) GetFruit(name string) Fruit {
	for _, fruit := range f.fruits {
		if fruit.name == name {
			return fruit
		}
	}
	fruit := Fruit{name: name}
	f.fruits = append(f.fruits, fruit)
	return fruit
}

func main() {
	fruits := Fruits{}
	fmt.Println(fruits)
	apple := fruits.GetFruit("apple")
	fmt.Println("Adding:", apple.name, "to fruits")
	fmt.Println(fruits)
	orange := fruits.GetFruit("orange")
	fmt.Println("Adding:", orange.name, "to fruits")
	fmt.Println(fruits)
}

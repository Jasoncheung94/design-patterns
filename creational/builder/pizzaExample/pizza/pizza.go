package pizza

import "fmt"

type Pizza struct {
	Size     string
	Sauce    string
	Toppings []string
}

func (p Pizza) Details() {
	fmt.Printf("Pizza is %s, %s and has %s\n", p.Size, p.Sauce, p.Toppings)
}

type PizzaBuilder interface {
	setSauce()
	setToppings()
	setSize()
	getPizza() Pizza
}

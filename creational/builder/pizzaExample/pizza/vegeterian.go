package pizza

type Veggie struct {
	Pizza
}

func (v *Veggie) setSauce() {
	v.Pizza.Sauce = "Tomato"
}

func (v *Veggie) setToppings() {
	v.Pizza.Toppings = []string{"Tomato", "Onion", "Mushroom"}
}

func (v *Veggie) setSize() {
	v.Pizza.Size = "Medium"
}

func (v *Veggie) getPizza() Pizza {
	return v.Pizza
}

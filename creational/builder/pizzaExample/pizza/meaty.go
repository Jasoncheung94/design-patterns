package pizza

type Meaty struct {
	Pizza
}

func (m *Meaty) setSauce() {
	m.Pizza.Sauce = "BBQ"
}

func (m *Meaty) setToppings() {
	m.Pizza.Toppings = []string{"Pepperoni", "Sausage", "Ham"}
}

func (m *Meaty) setSize() {
	m.Pizza.Size = "Large"
}

func (m *Meaty) getPizza() Pizza {
	return m.Pizza
}

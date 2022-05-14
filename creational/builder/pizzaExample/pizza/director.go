package pizza

type director struct {
	builder PizzaBuilder
}

func NewDirector(builder PizzaBuilder) *director {
	return &director{
		builder: builder,
	}
}

func (d *director) SetBuilder(builder PizzaBuilder) {
	d.builder = builder
}

func (d *director) Construct() Pizza {
	d.builder.setSauce()
	d.builder.setToppings()
	d.builder.setSize()
	return d.builder.getPizza()
}

package main

type AdidasFactory struct {
}

type AdidasShoe struct {
	shoe
}

func (a *AdidasFactory) makeShoe() ShoeInterface {
	return &AdidasShoe{
		shoe: shoe{
			size:  10,
			brand: "Adidas",
		},
	}
}

type AdidasJersey struct {
	jersey
}

func (a *AdidasFactory) makeJersey() JerseyInterface {
	return &AdidasJersey{
		jersey: jersey{
			size:  10,
			brand: "Adidas",
		},
	}
}

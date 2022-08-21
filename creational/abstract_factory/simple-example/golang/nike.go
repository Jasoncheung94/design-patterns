package main

type NikeFactory struct {
}

type NikeShoe struct {
	shoe
}

func (n *NikeFactory) makeShoe() ShoeInterface {
	return &NikeShoe{
		shoe: shoe{
			size:  10,
			brand: "Nike",
		},
	}
}

type NikeJersey struct {
	jersey
}

func (n *NikeFactory) makeJersey() JerseyInterface {
	return &NikeJersey{
		jersey: jersey{
			size:  10,
			brand: "Nike",
		},
	}
}

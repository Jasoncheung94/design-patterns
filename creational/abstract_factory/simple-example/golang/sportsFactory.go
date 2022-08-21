package main

import "fmt"

type SportsFactory interface {
	makeShoe() ShoeInterface
	makeJersey() JerseyInterface
}

// GetSportsFactory returns the appropriate sports factory for the given brand.
func GetSportsFactory(brand string) (SportsFactory, error) {
	switch brand {
	case "Adidas":
		return &AdidasFactory{}, nil
	case "Nike":
		return &NikeFactory{}, nil
	default:
		return nil, fmt.Errorf("Unknown brand %s", brand)
	}
}

package main

type ShoeInterface interface {
	setSize(size int)
	getSize() int
	setBrand(brand string)
	getBrand() string
}

type shoe struct {
	size  int
	brand string
}

func (s *shoe) setSize(size int) {
	s.size = size
}

func (s *shoe) getSize() int {
	return s.size
}

func (s *shoe) setBrand(brand string) {
	s.brand = brand
}

func (s *shoe) getBrand() string {
	return s.brand
}

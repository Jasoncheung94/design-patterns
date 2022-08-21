package main

type JerseyInterface interface {
	setSize(size int)
	getSize() int
	setBrand(brand string)
	getBrand() string
	setTeam(team string)
	getTeam() string
}

type jersey struct {
	size  int
	brand string
	team  string
}

func (j *jersey) setSize(size int) {
	j.size = size
}

func (j *jersey) getSize() int {
	return j.size
}

func (j *jersey) setBrand(brand string) {
	j.brand = brand
}

func (j *jersey) getBrand() string {
	return j.brand
}

func (j *jersey) setTeam(team string) {
	j.team = team
}

func (j *jersey) getTeam() string {
	return j.team
}

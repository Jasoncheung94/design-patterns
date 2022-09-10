package main

import "fmt"

type Collection interface {
	CreateIterator() Iterator
}

type Iterator interface {
	HasNext() bool
	Next() any
}

type ProjectCollection struct {
	projects []Project
}

func (p *ProjectCollection) CreateIterator() Iterator {
	return &ProjectIterator{
		projects: p.projects,
	}
}

type Project struct {
	id   int
	name string
}

type ProjectIterator struct {
	index    int
	projects []Project
}

func (p *ProjectIterator) HasNext() bool {
	return p.index < len(p.projects)
}

func (p *ProjectIterator) Next() any {
	if p.HasNext() {
		project := p.projects[p.index]
		p.index++
		return project
	}
	return nil
}

func main() {
	collection := &ProjectCollection{
		projects: []Project{
			{
				id:   1,
				name: "project1",
			},
			{
				id:   2,
				name: "project2",
			},
		},
	}

	iterator := collection.CreateIterator()
	for iterator.HasNext() {
		project := iterator.Next()
		fmt.Printf("Name: %s, ID: %d\n", project.(Project).name, project.(Project).id)
	}
}

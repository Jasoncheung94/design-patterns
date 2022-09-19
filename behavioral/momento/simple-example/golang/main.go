package main

import "fmt"

// Memento
type Memento struct {
	state string
}

func (m *Memento) getSavedState() string {
	return m.state
}

// Originator
type Originator struct {
	state string
}

func (e *Originator) createMemento() *Memento {
	return &Memento{state: e.state}
}

func (e *Originator) restoreMemento(m *Memento) {
	e.state = m.getSavedState()
}

func (e *Originator) setState(state string) {
	e.state = state
}

func (e *Originator) getState() string {
	return e.state
}

// Caretaker
type Caretaker struct {
	mementos []*Memento
}

func (c *Caretaker) addMemento(m *Memento) {
	c.mementos = append(c.mementos, m)
}

func (c *Caretaker) mementoAt(index int) *Memento {
	return c.mementos[index]
}

func main() {
	caretaker := &Caretaker{
		mementos: make([]*Memento, 0),
	}

	originator := &Originator{
		state: "State #1",
	}

	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("State #2")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("Corrupted State")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.restoreMemento(caretaker.mementoAt(1))
	fmt.Printf("Restored to State: %s\n", originator.getState())

	originator.restoreMemento(caretaker.mementoAt(0))
	fmt.Printf("Restored to State: %s\n", originator.getState())
	fmt.Println("Number of mementos", len(caretaker.mementos))
}

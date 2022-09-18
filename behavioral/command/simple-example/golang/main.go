package main

import "fmt"

type Command interface {
	execute()
}

type Factory struct {
	VehiclesCreated int
	Cars            int
	MotorBikes      int
}

func NewFactory() *Factory {
	return &Factory{}
}

type MakeCarCommand struct {
	numberToCreate int
	factory        *Factory
}

func (m *MakeCarCommand) execute() {
	m.factory.Cars += m.numberToCreate
	m.factory.VehiclesCreated += m.numberToCreate
	fmt.Println("Cars created: ", m.numberToCreate)
}

type MakeMotorBikeCommand struct {
	numberToCreate int
	factory        *Factory
}

func (m *MakeMotorBikeCommand) execute() {
	m.factory.MotorBikes += m.numberToCreate
	m.factory.VehiclesCreated += m.numberToCreate
	fmt.Println("MotorBikes created: ", m.numberToCreate)
}

func (f *Factory) MakeCar(numberToCreate int) Command {
	return &MakeCarCommand{
		numberToCreate: numberToCreate,
		factory:        f,
	}
}

func (f *Factory) MakeMotorBike(numberToCreate int) Command {
	return &MakeMotorBikeCommand{
		numberToCreate: numberToCreate,
		factory:        f,
	}
}

type Worker struct {
	commands []Command
}

func (w *Worker) AddCommand(c Command) {
	w.commands = append(w.commands, c)
}

func (w *Worker) executeCommands() {
	for _, command := range w.commands {
		command.execute()
	}
}

func main() {
	factory := NewFactory()
	worker := &Worker{}
	worker2 := &Worker{}

	worker.AddCommand(factory.MakeCar(2))
	worker2.AddCommand(factory.MakeCar(5))
	worker.AddCommand(factory.MakeMotorBike(3))
	worker.executeCommands()
	worker2.executeCommands()

	fmt.Println("Total vehicles created: ", factory.VehiclesCreated)
	fmt.Println("Total cars created: ", factory.Cars)
	fmt.Println("Total motorbikes created: ", factory.MotorBikes)
}

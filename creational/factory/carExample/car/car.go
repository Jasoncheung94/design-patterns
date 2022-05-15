package car

// car interface
type icar interface {
	setNumberWheels(num int64)
	setColor(color string)
	setMake(make string)
	setModel(model string)
}

// car - concrete item.
type car struct {
	NumberOfWheels int64
	Color          string
	Make           string
	Model          string
}

func (c *car) setNumberWheels(num int64) {
	c.NumberOfWheels = int64(num)
}

func (c *car) setColor(color string) {
	c.Color = color
}

func (c *car) setMake(make string) {
	c.Make = make
}

func (c *car) setModel(model string) {
	c.Model = model
}

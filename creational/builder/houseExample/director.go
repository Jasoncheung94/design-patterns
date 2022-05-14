package main

type director struct {
	builder Builder
}

func newDirector(builder Builder) *director {
	return &director{
		builder: builder,
	}
}

func (d *director) setBuilder(builder Builder) {
	d.builder = builder
}

func (d *director) construct() house {
	d.builder.setWindowType()
	d.builder.setDoorType()
	d.builder.setNumFloors()
	return d.builder.getHouse()
}

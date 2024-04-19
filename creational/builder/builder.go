package builder

import (
	"fmt"
)

type Speed float64

const (
	MPH Speed = 1
	KPH Speed = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor Color = "green"
	RedColor   Color = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels  Wheels = "steel"
)

type Builder interface {
	Paint(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() ICar
}

type ICar interface {
	Drive()
	Stop()
}

type Car struct {
	color  Color
	wheels Wheels
	speed  Speed
}

func (c *Car) Drive() {
	fmt.Println("Car is driving")
}

func (c *Car) Stop() {
	fmt.Println("Car has stopped")
}

type CarBuilder struct {
	car *Car
}

func NewBuilder() *CarBuilder {
	return &CarBuilder{car: &Car{}}
}

func (cb *CarBuilder) Paint(color Color) Builder {
	cb.car.color = color
	return cb
}

func (cb *CarBuilder) Wheels(wheels Wheels) Builder {
	cb.car.wheels = wheels
	return cb
}

func (cb *CarBuilder) TopSpeed(speed Speed) Builder {
	cb.car.speed = speed
	return cb
}

func (cb *CarBuilder) Build() ICar {
	return cb.car
}

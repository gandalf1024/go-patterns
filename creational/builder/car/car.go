package car

import "fmt"

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Interface
}

type Interface interface {
	Drive() error
	Stop() error
}

type MyInterface struct {
}

func (my *MyInterface) Drive() error {
	return nil
}

func (my *MyInterface) Stop() error {
	return nil
}

type CarBuilder struct {
}

func (cc *CarBuilder) Color(ss Color) Builder {
	return cc
}

func (cc *CarBuilder) Wheels(xx Wheels) Builder {
	return cc
}

func (cc *CarBuilder) TopSpeed(xx Speed) Builder {
	return cc
}

func (cc *CarBuilder) Build() Interface {
	mm := MyInterface{}
	return &mm
}

func (cc *CarBuilder) Paint(s string) Builder {
	fmt.Println(s)
	return cc
}

func NewCarBuilder() *CarBuilder {
	car := CarBuilder{}
	return &car
}

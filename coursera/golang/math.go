package test

import (
	"fmt"
	"math"
)

type Circle struct {
	X, Y, R float64
}

type Rectangle struct {
	X1, Y1, X2, Y2 float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.R
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.R
}

func (r *Rectangle) area() float64 {
	l := math.Abs(r.X2 - r.X1)
	w := math.Abs(r.Y2 - r.Y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	var res float64 = 2 * (math.Abs(r.X2-r.X1) + math.Abs(r.Y2-r.Y1))
	return res
}

type Shape interface {
	area() float64
	perimeter() float64
}

func ShowTotalAreas(shapes ...Shape) {
	for i, s := range shapes {
		fmt.Println(i, ". area = ", s.area())
	}
}

type Multishape struct {
	Shapes []Shape
}

func (m *Multishape) ShowPerimeters() {
	for i := 0; i < len(m.Shapes); i++ {
		fmt.Println(i, ". perimeter=", m.Shapes[i].perimeter())
	}
}

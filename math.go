
package test

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) area() float64 {
	l := math.Abs(r.x2 - r.x1)
	w := math.Abs(r.y2 - r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	var res float64 = 2 * (math.Abs(r.x2-r.x1) + math.Abs(r.y2-r.y1))
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
	shapes []Shape
}

func (m *Multishape) showPerimeters() {
	for i := 0; i < len(m.shapes); i++ {
		fmt.Println(i, ". perimeter=", m.shapes[i].perimeter())
	}
}

package main

import (
	"fmt"
	"math"
)

type Polygon interface {
	Area() float64
	Perimeter() float64
}

type MultPolygons struct {
	polygons []Polygon
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	x, y, r float64
}

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

func main() {
	a := new(Android)
	a.Name = "José"
	a.Talk()
	person := Person{"Rafael"}
	person.Talk()
	circle := Circle{0, 0, 5}
	rectangle := Rectangle{15, 10}
	multPolygons := MultPolygons{polygons: []Polygon{&circle, &rectangle}}
	fmt.Println(multPolygons.totalArea())
	fmt.Println(multPolygons.totalPerimeter())
	fmt.Println(circle.Area())
	fmt.Println(rectangle.Area())
}

func (m *MultPolygons) totalArea() float64 {
	var area float64
	for _, p := range m.polygons {
		area += p.Area()
	}
	return area
}

func (m *MultPolygons) totalPerimeter() float64 {
	var perimeter float64
	for _, p := range m.polygons {
		perimeter += p.Perimeter()
	}
	return perimeter
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (c *Circle) Area() float64 {
	return math.Pi * (c.r * c.r)
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (p *Person) Talk() {
	fmt.Println("Olá meu nome é", p.Name)
}

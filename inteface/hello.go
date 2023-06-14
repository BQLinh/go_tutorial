package main

import (
	"fmt"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	length float64
	height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.length * r.height
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (r rect) perim() float64 {
	return (r.height + r.height) * 2
}

func (c circle) perim() float64 {
	return c.radius * 2 * 3.14
}

func show(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{length: 2, height: 3}

	c := circle{radius: 3}

	show(r)
	show(c)
}

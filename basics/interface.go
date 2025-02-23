package main

import (
	"fmt"
	"math"
)

// unlike struct which is a collection of fields/data, interface is a collection of mmethods
type calculation interface {
	area() float32
	perim() float32
}

type rect struct {
	width, height float32
}

type circle struct {
	radius float32
}

func (r rect) area() float32 {
	return r.width * r.height
}
func (r rect) perim() float32 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float32 {
	return 2 * math.Pi * c.radius
}

func calculateRectArea(c calculation) {
	g, ok := c.(rect)
	if ok {
		fmt.Println(g.area())
	}
}

func main() {
	r := rect{width: 3, height: 4}
	calculateRectArea(r)
}

package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

func main() {
	t := triangle{}
	s := square{}

	// sample
	t.base = 10
	t.height = 5
	s.side = 4

	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println("Area:", s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.side * s.side
}

package main

import "fmt"

type square struct {
	sidelength float64
}

type triangle struct {
	height float64
	base   float64
}

func (s square) getArea() float64 {
	return s.sidelength * s.sidelength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	t := triangle{base: 10, height: 10}
	s := square{sidelength: 10}
	printArea(s)
	printArea(t)
}

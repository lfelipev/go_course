package main

import "fmt"

type triangle struct {
	b float64
	h float64
}
type square struct {
	l float64
}

type shape interface {
	getArea() float64
}

func main() {
	t := triangle{
		b: 15,
		h: 15,
	}
	s := square{
		l: 5,
	}

	printAreas(t)
	printAreas(s)
}

func (t triangle) getArea() float64 {
	return (t.b * t.h) / 2
}

func (s square) getArea() float64 {
	return s.l * s.l
}

func printAreas(s shape) {
	fmt.Println(s.getArea())
}

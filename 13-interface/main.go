package main

import (
	"fmt"
	"math"
)

type plane interface {
	area() float64
	circumfance() float64
}

type quadrangle struct {
	side float64
}

type parallelograms struct {
	base float64
	side float64
}

func (q quadrangle) area() float64 {
	return math.Pow(q.side, 2)
}

func (q quadrangle) circumfance() float64 {
	return q.side * 4
}

func (p parallelograms) area() float64 {
	return p.base * p.side
}

func (p parallelograms) circumfance() float64 {
	return (p.base + p.side) * 2
}

func main() {
	var pl plane = quadrangle{side: 5}

	fmt.Printf("Quadrangle area\t\t\t : %f\n", pl.area())
	fmt.Printf("Quadrangle circumfance\t\t : %f\n", pl.circumfance())

	pl = parallelograms{base: 5, side: 3}

	fmt.Printf("Parallelograms area\t\t : %f\n", pl.area())
	fmt.Printf("Parallelograms circumfance\t : %f\n", pl.circumfance())
}

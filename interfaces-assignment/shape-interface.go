package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	side   float64
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (rc triangle) getArea() float64 {
	return 0.5 * rc.side * rc.height
}

func printArea(sh shape) {
	area := sh.getArea()
	fmt.Println("The Area is: ", area)
}

package main

import "fmt"

type shape interface {
	getArea() float64
}
type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {

	tr := triangle{
		height: 5,
		base:   10,
	}
	sq := square{
		sideLength: 10,
	}

	fmt.Println(tr.getArea())
	fmt.Println(sq.getArea())

	printArea(tr)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (triangle) getArea() float64 {
	return 0.5 * base * height
}

func (square) getArea() float64 {
	return sideLength * sideLength
}

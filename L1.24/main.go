package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) Distance(other Point) float64 {
	return math.Sqrt(math.Pow(other.x-p.x, 2) + math.Pow(other.y-p.y, 2))
}

func main() {

	p1 := NewPoint(3, 7)
	p2 := NewPoint(5, 2)

	fmt.Println(p1.Distance(*p2))

}

/*
Output:
5.385164807134504
*/

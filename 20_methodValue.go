package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x - p.x, q.y - p.y)
}

func main() {

	p, q := Point{1.2,3.4}, Point{2.2, 1.1}

	fmt.Println(p.Distance(q))

	distanceFromP := p.Distance // Method value, here the p is copied, 
	// even if you copied 

	fmt.Println(distanceFromP(q))

	fmt.Printf("%T \n", distanceFromP)

	fmt.Printf("%T \n", Point.Distance)

	p= Point{10,20}

	fmt.Println(distanceFromP(q)) // same value
		



}
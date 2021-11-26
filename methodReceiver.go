package main

import (
	"fmt"
)

type Point struct {
	x, y float64
}


// THis is avalue receiver as we are gettign a copy of the actual object.
func (p Point) createNew() Point{
	return Point{p.x,p.y}
}

func (p *Point) add(x float64) {
	p.x += x
	p.y += x
}

func (p Point) String() string {
	return fmt.Sprintf("[ X=%f  Y = %f]", p.x, p.y)
}

func main(){

	 p := Point{x: 10.0, y:20.0}
	fmt.Println("1. Point = ", p)
	p.add(2.0)
	fmt.Println("2. After adding Point = ", p)

	fmt.Println(" Creating new ")
	p2 := p.createNew()
	p2.add(4.5)

	fmt.Println("3. OLD Point = ", p)
	fmt.Println("4. Creation New = ", p2)


}



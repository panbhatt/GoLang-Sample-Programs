// GO class 18. 

package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Line struct {
	Begin, End Point
}

func (l Line) Distance() float64{
	return math.Hypot(l.End.X - l.Begin.X , l.End.Y - l.Begin.Y)
}

func (l *Line) ScaleBy(f float64) {
	l.End.X=  (f-1) * (l.End.X - l.Begin.X )
	l.End.Y=  (f-1) * (l.End.Y - l.Begin.Y )
 }

// Path
type Path []Point

// This function will return float64 sum because by default its value can be zero
func (p Path) Distance() (sum float64)  {

	for i:=1 ;i < len(p); i++ {
		sum  += Line{ p[i-1], p[i]}.Distance()
	}
	return sum
}

// Creating formal interface DIstanceer like Stringer
type Distancer interface {
	Distance() float64
}

func PrintDistance(d Distancer) {
	fmt.Println(d.Distance())
}


func main() {

	var l Line = Line{ Point{1,2}, Point{10,20}, 
	}

	fmt.Println("Line Length = " , l.Distance())

	l.ScaleBy(2.5)

	fmt.Println("Line Length After Scaling = " , l.Distance())

	var p Path = Path{ {1,2}, {3,4}, {5,6}, {7,8}}

	fmt.Println("Path Perimiter = " , p.Distance())

	// via using Distancer
	PrintDistance(l)
	PrintDistance(p)


}
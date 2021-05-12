// WAP to implement a square struct with center of type point and length.

package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func (p *Point) Move(dx int, dy int) {
	p.x = p.x + dx
	p.y = p.y + dy
}

func NewPoint(x int, y int) *Point{
	p:= &Point{x,y}
	return p
	}

type Square struct {
	centre Point
	length int
}

func (s *Square) Area() int {
	return s.length * s.length
}

func NewSquare(point *Point, l int) (*Square, error) {
	if l < 0 {
		return nil, fmt.Errorf("Negative length not acceptable")
	}
	s := &Square{
		centre: Point{point.x, point.y},
		length: l,
	}
	return s,nil
}

func main() {

	point := Point{10, 12}
	square,_ := NewSquare(&point, 12)
	fmt.Printf("Area of original square == %d\n square centre = %+v", square.Area(), square.centre)
	point.Move(2, 3)
	moved_square,_ := NewSquare(&point, 14)
	fmt.Printf("Area of displaced square == %d\n square centre = %+v", moved_square.Area(), moved_square.centre)
	fmt.Printf("%+v\n", point)

}

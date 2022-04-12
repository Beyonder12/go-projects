package main

import (
	"fmt"
	"math"
)

//interface is a collection of method definition
// which have not definition which is wrapped with certain name
// interface is data type where the zero value is nil

func main() {
	var twoDimensionalFigure calculate

	twoDimensionalFigure = square{10.0}
	fmt.Println("==== square")
	fmt.Println("area   : ", twoDimensionalFigure.area())
	fmt.Println("perimeter : ", twoDimensionalFigure.perimeter())

	twoDimensionalFigure = circle{14.0}
	fmt.Println("==== circle")
	fmt.Println("area   : ", twoDimensionalFigure.area())
	fmt.Println("perimeter : ", twoDimensionalFigure.perimeter())
	fmt.Println("radius : ", twoDimensionalFigure.(circle).radius)

}

type calculate interface {
	area() float64
	perimeter() float64
}

type circle struct {
	diameter float64
}

func (c circle) radius() float64 {
	return c.diameter / 2
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius(), 2)
}

func (c circle) perimeter() float64 {
	return math.Pi * c.diameter
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return math.Pow(s.side, 2)
}

func (s square) perimeter() float64 {
	return s.side * 4
}

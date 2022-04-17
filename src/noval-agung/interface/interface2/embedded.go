package main

import (
	"fmt"
	"math"
)

func main() {
	var threeDFigure calculate = &cube{4}

	fmt.Println("==== cube")
	fmt.Println("area  :", threeDFigure.area())
	fmt.Println("perimeter :", threeDFigure.perimeter())
	fmt.Println("volume ", threeDFigure.volume())
}

type calculate2d interface {
	area() float64
	perimeter() float64
}

type calculate3d interface {
	volume() float64
}

type calculate interface {
	calculate2d
	calculate3d
}

type cube struct {
	side float64
}

func (c *cube) volume() float64 {
	return math.Pow(c.side, 3)
}

func (c *cube) area() float64 {
	return math.Pow(c.side, 2) * 6
}

func (c *cube) perimeter() float64 {
	return c.side * 12
}

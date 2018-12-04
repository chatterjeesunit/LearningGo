package main

import (
	"fmt"
	"math"
)

func main() {

	c1 := Circle{5, Point{5, 4}, "red"}
	c2 := Circle{5, Point{-3, 4}, "yellow"}

	fmt.Printf("Red Color = %s, Area = %f, Circumference = %f \n", c1.color, c1.Area(), c1.Circumference())
	fmt.Printf("Yellow Color = %s, Area = %f, Circumference = %f \n", c2.color, c2.Area(), c2.Circumference())


	fmt.Printf("Red Circle Overlaps Yellow Circle : %v", c1.Overlaps(c2))
	
}

type Point struct {
	x int
	y int
}

type Circle struct {
	radius int
	center Point
	color string
}


func (p *Point) distanceTo(o Point) float64 {
	return math.Sqrt(
			math.Pow(float64(o.x - p.x), 2) +
			math.Pow(float64(o.y - p.y), 2))
}

func (c *Circle) Area() float64 {
	return math.Pow(float64(c.radius), 2) * math.Pi
}

func (c *Circle) Circumference() float64 {
	return 2 * math.Pi * float64(c.radius)
}

func (this *Circle) Overlaps(other Circle) bool {

	distanceBetweenCenters := other.center.distanceTo(this.center)

	return distanceBetweenCenters < float64(this.radius + other.radius)
}

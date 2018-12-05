package main

import (
	"fmt"
	"math"
)

func main() {

	c1 := Circle{5, Point{5, 4}, "Red"}
	c2 := Circle{6, Point{-3, -5}, "Yellow"}

	fmt.Printf("%s Circle: Area = %f, Circumference = %f \n\n", c1.color, c1.Area(), c1.Circumference())
	fmt.Printf("%s Circle: Area = %f, Circumference = %f \n\n", c2.color, c2.Area(), c2.Circumference())

	fmt.Printf("Red Circle Overlaps Yellow Circle : %v\n\n", c1.Overlaps(c2))

	fmt.Println("Now Increasing radius of Red circle")
	c1.ChangeRadius(10)

	fmt.Printf("Updated %s Circle: Area = %f, Circumference = %f \n\n", c1.color, c1.Area(), c1.Circumference())
	fmt.Printf("Red Circle Overlaps Yellow Circle : %v\n", c1.Overlaps(c2))
}

type Point struct {
	x int
	y int
}

type Circle struct {
	radius int
	center Point
	color  string
}

func (p Point) distanceTo(o Point) float64 {
	return math.Sqrt(
		math.Pow(float64(o.x-p.x), 2) +
			math.Pow(float64(o.y-p.y), 2))
}

func (c Circle) Area() float64 {
	return math.Pow(float64(c.radius), 2) * math.Pi
}

func (c Circle) Circumference() float64 {
	return 2 * math.Pi * float64(c.radius)
}

func (this Circle) Overlaps(other Circle) bool {

	distanceBetweenCenters := other.center.distanceTo(this.center)

	return distanceBetweenCenters < float64(this.radius+other.radius)
}

// Circle is passed to reciever as a pointer
// Ideally this should be the way we define recievers, otherwise entire objects are passed as copy to reciever functions.
func (this *Circle) ChangeRadius(radius int) {
	this.radius = radius
}

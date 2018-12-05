package main

import (
	"fmt"
	"math"
)


type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length float64
	breadth float64
}

type Circle struct {
	radius float64
}

type Square struct {
	length float64
}

// Won't work if we make it pointer reciever
func (r Rectangle) Area() float64 {
	return r.breadth*r.length
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

// Won't work if we make it pointer reciever
func (s Square) Area() float64 {
	return s.length* s.length
}

func (s Square) Perimeter() float64 {
	return 4 * s.length
}


// Won't work if we make it pointer reciever
func (c Circle) Area() float64 {
	return math.Pi* c.radius* c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}


func main() {

	shapes := make([]Shape, 0, 3)

	shapes = append(shapes,
		Circle{5} ,
		Rectangle{4, 8},
		Square{3})

	for _,shape := range shapes {
		fmt.Printf("Area = %f, Perimeter = %f\n", shape.Area(), shape.Perimeter())
	}

	fmt.Println()

	for _,shape := range shapes {
		//With Interface assersions
		switch shape.(type) {
		case Rectangle: fmt.Print("Rectangle : ")
		case Circle: fmt.Print("Circle : ")
		case Square: fmt.Print("Square : ")

		}
		fmt.Printf("Area = %f, Perimeter = %f\n", shape.Area(), shape.Perimeter())
	}

}
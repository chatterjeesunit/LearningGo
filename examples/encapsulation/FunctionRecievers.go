package main

import (
	"fmt"
	"math"
)

func main() {


	fmt.Println("Calling Functions the normal way")
	fmt.Printf("4^2 = %d, and 1/4 = %f\n\n", Square(4), Reciprocal(4))


	fmt.Println("Calling Functions using function receivers")
	num := MyNumber(4)
	fmt.Printf("4^2 = %d, and 1/4 = %f", num.Square(), num.Reciprocal())
}



func Square(n int) int {
	return int(math.Pow(float64(n), 2.0))
}


func Reciprocal(n int) float64 {
	return 1.0/float64(n)
}


type MyNumber int

func (n MyNumber) Square() int {
	return int(math.Pow(float64(n), 2.0))
}


func (n MyNumber) Reciprocal() float64 {
	return 1.0/float64(n)
}

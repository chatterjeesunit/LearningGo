package main

import (
	"fmt"
	"math"
)

func main() {

	helloWord()

	helloX("Sunit Chatterjee !!")

	add(10, 20)

	printValues(10, 20, "Testing", true)

	area := calculateAreaOfRectangle(5, 7)
	fmt.Printf("Area of rectangle is %d\n", area)

	sum1, avg1, min1, max1 := calculate(5, 7, 2)
	fmt.Printf("Sum = %d, Avg = %f, Min = %d, Max = %d \n", sum1, avg1, min1, max1)


	sum2, avg2, min2, max2 := calculateAgain(5, 7, 2)
	fmt.Printf("Sum = %d, Avg = %f, Min = %d, Max = %d \n", sum2, avg2, min2, max2)
}

//function without any arguments or return types
func helloWord() {
	fmt.Println("Hello World")
}


//function with a single argument
func helloX(name string) {
	fmt.Println("Hello " + name)
}

//More than 1 arguments
func add(x, y int) {
	fmt.Printf("Sum = %d\n", x + y )
}

func printValues(a, b int, c string, d bool) {
	fmt.Printf("A = %d, B = %d, c = %s, d = %v\n", a, b, c, d)
}

//with Return values
func calculateAreaOfRectangle(length int, breadth int) int {
	return length * breadth
}


//With more than 1 return values (option 1)
func calculate(x, y, z int) (int, float64, int, int) {
	sum := x + y + z
	avg := float64(sum) / 3
	min := int(math.Min(math.Min(float64(x), float64(y)), float64(z)))
	max := int(math.Max(math.Max(float64(x), float64(y)), float64(z)))

	return sum, avg, min, max
}

//With more than 1 return values (option 2 : Named Return value)
func calculateAgain(x, y, z int) (sum int, avg float64, min, max int) {
	sum = x + y + z
	avg = float64(sum) / 3
	min = int(math.Min(math.Min(float64(x), float64(y)), float64(z)))
	max = int(math.Max(math.Max(float64(x), float64(y)), float64(z)))

	return
}
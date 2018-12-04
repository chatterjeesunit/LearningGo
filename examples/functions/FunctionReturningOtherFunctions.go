package main

import (
	"fmt"
)

func main() {

	sumIntegers := SumN(idFunc)
	sumSquares := SumN(squareFunc)
	sumCubes := SumN(cubeFunc)

	fmt.Printf("Sum of Numbers till 5 = %d \n", sumIntegers(5))

	fmt.Printf("Sum of Numbers till 5 = %d \n", sumSquares(5))

	fmt.Printf("Sum of Numbers till 5 = %d \n", sumCubes(5))

}


type myFunc func(int)int

//Used to calculate sum of numbers or squares of numbers, etc
//Numbers from 1 to N, where N is taken as input
func SumN(f myFunc) func(int) int {

	return func(number int) int {
		sum := 0
		for i := 1; i <= number; i++ {
			sum += f(i)
		}

		return sum
	}
}

func idFunc(x int) int { return x }

func squareFunc(x int) int { return x * x }

func cubeFunc(x int) int { return x * x * x }

package main

import (
	"fmt"
)

func main() {

	sum5 := calculateSumN(5, identity)
	fmt.Printf("Sum of Numbers till 5 = %d \n", sum5)

	sumSquare5 := calculateSumN(5, square)
	fmt.Printf("Sum of Numbers till 5 = %d \n", sumSquare5)

	sumCube5 := calculateSumN(5, cube)
	fmt.Printf("Sum of Numbers till 5 = %d \n", sumCube5)
	
}


type multiplyFunc func(int)int


func identity(x int) int { return x }

func square(x int) int { return x * x }

func cube(x int) int { return x * x * x }

//Used to calculate sum of numbers or squares of numbers, etc
//Numbers from 1 to N, where N is taken as input
func calculateSumN(number int,  f multiplyFunc) int   {
	sum := 0
	for i:=1; i<=number; i++ {
		sum += f(i)
	}

	return sum
}
package main

import (
	"fmt"
)

func main() {

	num := 5
	env := "zyx"

	if num%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}

	if env == "uat" {

		fmt.Println("Apply UAT environment config")

	} else if env == "prod" {

		fmt.Println("Apply Production config")

	} else if env == "dev" {

		fmt.Println("Apply Dev config")

	} else {
		fmt.Println("Invalid Environment")
	}

	if remainder := num % 5; remainder == 0 {
		fmt.Printf("%d is divisible by 5 ", num)
	} else {
		fmt.Printf("%d is not divisible by 5. Remainder = %d", num, remainder)
	}

}

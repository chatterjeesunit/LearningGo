package main

import (
	"fmt"
)

func main() {


	//continue statement
	for i:=0; i < 10; i++ {

		fmt.Print(" * ")
		if i % 2 != 0 {
			// If number is odd, continue to start of loop
			continue
		}

		square := i * i
		fmt.Print(square)
	}
}

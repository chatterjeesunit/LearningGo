package main

import (
	"fmt"
	"strconv"
)

func main() {

	for {
		var input string

		fmt.Print("Please enter a number : ")
		fmt.Scan(&input)

		num, err := strconv.Atoi(input)

		if err == nil {
			fmt.Println(num)
		}else {
			fmt.Println("Invalid number : " + err.Error())
		}
	}
}

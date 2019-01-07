package main

import "fmt"

// SHOWING THAT GO routines exit when main function exits

func main() {

	fmt.Println("Start of Program")

	go printMessage("Hello World..")

	go printMessage("Welcome to Go!")

	fmt.Println("End of Program")
}


func printMessage(message string) {
	fmt.Println(message)
}

package main

import "fmt"

// The main go routines does not wait for the goroutine to finish.

func main() {

	fmt.Println("Start of Program")
	go dummyPrint()
	fmt.Println("End of Program")

	
}

func dummyPrint() {
	fmt.Println("Hello World")
}
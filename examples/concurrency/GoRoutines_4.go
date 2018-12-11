package main

import (
	"fmt"
)

// another example of synchronization using channels

func main() {

	c := make(chan int)

	fmt.Println("Start of Program")

	go dummyPrintAgain(c)

	 <- c

	fmt.Println("End of Program")

	
}

func dummyPrintAgain(c chan int) {
	fmt.Println("Hello World")
	c <- 0
}
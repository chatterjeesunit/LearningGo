package main

import (
	"fmt"
	"time"
)

// The main go routines does not wait for the goroutine to finish.

func main() {

	fmt.Println("Start of Program")
	go dummyPrint()
	time.Sleep(300 * time.Millisecond)
	fmt.Println("End of Program")

	
}

func dummyPrint() {
	for i:=0; i <10; i++{
		fmt.Println("Hello World")
		time.Sleep(100 * time.Millisecond)
	}

}
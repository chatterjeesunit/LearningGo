package main

import (
	"fmt"
	"time"
)

// hack to force main function to wait few seconds for go routine to finish

func main() {

	fmt.Println("Start of Program")

	go printMessage("Hello World..")

	go printMessage("Welcome to Go!")

	//force the main function to wait
	time.Sleep(time.Millisecond * 100)

	fmt.Println("End of Program")

}

func printMessage(message string) {
	fmt.Println(message)
}

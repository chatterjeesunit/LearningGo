package main

import (
	"fmt"
	"time"
)

// hack to force main function to wait few seconds for go routine to finish

func main() {

	fmt.Println("Start of Program")

	go PrintAMessage("Hello World..")

	go PrintAMessage("Welcome to Go!")

	//force the main function to wait
	time.Sleep(time.Second * 1)

	fmt.Println("End of Program")

}


func PrintAMessage(message string) {
	fmt.Println(message)
}

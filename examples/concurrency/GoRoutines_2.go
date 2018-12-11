package main

import (
	"fmt"
	"time"
)

// hack to force main function to wait few seconds for go routine to finish

func main() {

	fmt.Println("Start of Program")

	go dummyPrint()

	time.Sleep(time.Second * 1)

	fmt.Println("End of Program")

	
}

func dummyPrint() {
	fmt.Println("Hello World")
}
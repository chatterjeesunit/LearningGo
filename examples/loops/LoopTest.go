package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	// Simple Loop
	for i := 1; i <= 20; i = i + 2 {
		fmt.Print(i, " ")
	}

	fmt.Println("")


	// Equivalent to a while loop
	i := 1
	for i <= 20 {
		fmt.Print(i, " ")
		rand.Seed(time.Now().UnixNano())
		i = i + rand.Intn(5)
	}
	fmt.Println("")

	for x:=0; x < 3 ; x++ {
		for y:=0; y < 2; y++ {
			fmt.Printf("[%d:%d], ", x, y)
		}
	}

	fmt.Println("")
	// Infinite loop
	for {

		_, err := http.Get("http://www.google.com")
		if err == nil {
			fmt.Println("Working")
		} else {
			fmt.Println("Not working!!")
		}
		time.Sleep(time.Second * 5)
	}
}

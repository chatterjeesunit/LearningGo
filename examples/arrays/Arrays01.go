package main

import "fmt"

func main() {

	arr := [2]string {"Hello", "World"}
	fmt.Println("1. Main : ", arr)

	transformArrayViaPointers(&arr)
	fmt.Println("2. Main : ", arr)
}

func transformArray(arr [2]string) {
	arr[1] = "Sunit"
	fmt.Println("\tInside Transform Array : ", arr)
}


func transformArrayViaPointers(arr *[2]string) {
	arr[1] = "Sunit"
	fmt.Println("\tInside Transform Array : ", arr)
}
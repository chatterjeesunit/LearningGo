package main

import (
	"fmt"
)

func main() {

	//Created an array of 5 elements
	array := [5]int{1, 2, 3, 4, 5}

	fmt.Print("Initial Array State : \t\t\t\t\t")
	fmt.Println(array)

	tryToModifyArrayDirectly(array)
	fmt.Print("After passing array as copy to a function :\t\t")
	fmt.Println(array)

	modifyArrayUsingPointer(&array)
	fmt.Print("After passing array as a pointer to a function :\t")
	fmt.Println(array)

	modifyArrayUsingSlice(array[:])
	fmt.Print("After passing array as slice to a function :\t\t")
	fmt.Println(array)

}

func tryToModifyArrayDirectly(arr [5]int) {

	for i := range arr {
		arr[i] = arr[i] * 2
	}
}

func modifyArrayUsingPointer(arr *[5]int) {

	for i := range arr {
		arr[i] = arr[i] * 3
	}
}

func modifyArrayUsingSlice(arr []int) {

	for i := range arr {
		arr[i] = arr[i] * 5
	}
}

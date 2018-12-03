package main

import (
    "fmt"
)


func main() {

    valA := 10
    addressA := &valA

    fmt.Printf("Address of A = %v\n", addressA)
    fmt.Printf("Value of A = %d\n", valA)

    fmt.Printf("Value at address of A (fetched by pointers) = %d \n\n", *addressA)
    

    valB := 20
    addressB := &valB

    fmt.Printf("Address of B = %v\n", addressB)
    fmt.Printf("Value of B = %d\n", valB)
    fmt.Printf("Value at address of B (fetched by pointers) = %d \n\n", *addressB)
    
    //Now lets change value of A using Pointers

    *addressA = 30
    fmt.Printf("Address of A is still unchanged = %v, ", &valA)
    fmt.Printf("but value A is now changed to : %d\n", valA)


    //Now change address pointers
    addressA = addressB

    fmt.Printf("A = %d, and address of A = %v \n\n", valA, &valA)


    fmt.Printf("Address A = %v, and value at addressA pointer = %d \n\n", addressA, *addressA)

}
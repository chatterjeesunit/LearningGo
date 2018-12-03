package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	acceleration, initialVelocity, initialDisplacement := readInputs()

	displacementFunc := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	fmt.Printf("Displacement after 3 seconds : %f \n", displacementFunc(3))

	fmt.Printf("Displacement after 5 seconds : %f \n", displacementFunc(5))

	fmt.Printf("Displacement after 10 seconds : %f \n", displacementFunc(10))

}

func GenDisplaceFn(acceleration float64, initialVelocity float64, initialDisplacement float64) func(float64) float64 {

	return func(time float64) float64 {
		return 0.5*acceleration*math.Pow(time, 2.0) + initialVelocity*time + initialDisplacement
	}
}

func readInputs() (acceleration float64, initialVelocity float64, initialDisplacement float64) {

	var accelerationInput, velocityInput, displacementInput string

	fmt.Printf("Enter a value for Acceleration : ")
	fmt.Scan(&accelerationInput)

	fmt.Printf("Enter a value for Initial Velocity : ")
	fmt.Scan(&velocityInput)

	fmt.Printf("Enter a value for Displacement : ")
	fmt.Scan(&displacementInput)

	acceleration, err := strconv.ParseFloat(accelerationInput, 64)

	if err == nil {
		initialVelocity, err = strconv.ParseFloat(velocityInput, 64)
	}

	if err == nil {
		initialDisplacement, err = strconv.ParseFloat(displacementInput, 64)
	}

	if err != nil {
		fmt.Println("Invalid inputs. Error : " + err.Error())
		os.Exit(1)
	}

	return
}

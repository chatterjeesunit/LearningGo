package main

import (
	"fmt"
	"strings"
)

func main() {

	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	for {

		fmt.Print("> ")
		var animalInput, actionInput string
		fmt.Scan(&animalInput)
		fmt.Scan(&actionInput)

		var animal Animal
		switch strings.ToLower(animalInput) {
		case "cow":
			animal = cow
		case "bird":
			animal = bird
		case "snake":
			animal = snake
		default:
			fmt.Println("Invalid Animal Name. Please try again")
			continue
		}

		switch strings.ToLower(actionInput) {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Invalid Animal Action. Please try again")
			continue
		}
	}
}

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

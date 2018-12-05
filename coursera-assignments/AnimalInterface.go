package main

import (
	"fmt"
	"strings"
)


type Animal interface {
	Eat()
	Speak()
	Move()
}

type Cow struct {
	name string
}

type Snake struct {
	name string
}

type Bird struct {
	name string
}


func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}



func (c Snake) Speak() {
	fmt.Println("hsss ")
}

func (c Snake) Eat() {
	fmt.Println("mice")
}

func (c Snake) Move() {
	fmt.Println("slither ")
}



func (c Bird) Eat() {
	fmt.Println("worms")
}

func (c Bird) Move() {
	fmt.Println("fly ")
}

func (c Bird) Speak() {
	fmt.Println("fly ")
}


const (
	NEW_ANIMAL_REQUEST = "newanimal"
	NEW_QUERY_REQUEST = "query"
)



func main() {

	animalMap := make(map[string]Animal)

	for {

		fmt.Print("> ")

		choice, name, arg3 := readUserInput()

		choice = strings.ToLower(choice)

		switch choice {

		case NEW_ANIMAL_REQUEST:
			animal := createNewAnimal(arg3, name)
			if animal != nil {
				animalMap[name] = animal
				fmt.Println("Created it!")
			}

		case NEW_QUERY_REQUEST:
			animal, ok := animalMap[name]
			if ok {
				queryAnimalAction(arg3, animal)
			}else {
				fmt.Println("Animal with this name does not exists. Please create a new animal first")
			}
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

	}
}

func queryAnimalAction(action string, animal Animal) {
	switch strings.ToLower(action) {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Invalid choice for action. Please try again")
	}
}

func createNewAnimal(animalType string, name string) (animal Animal) {
	switch strings.ToLower(animalType) {
	case "cow":
		animal = Cow{name}
	case "bird":
		animal = Bird{name}
	case "snake":
		animal = Snake{name}
	default:
		fmt.Println("Invalid animal type. Please try again.")
	}
	return animal
}


func readUserInput() (input1, input2, input3 string) {

	fmt.Scan(&input1, &input2, &input3)
	return
}
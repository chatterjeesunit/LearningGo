package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type person struct {
	fname string
	lname string
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter a file name: ")

	var fileName string
	if scanner.Scan() {
		fileName = scanner.Text()
	}

	persons := readFile(fileName)

	for _, p := range persons {
		fmt.Printf("%s %s\n", p.fname, p.lname)
	}

}

func readFile(fileName string) []person {

	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		fmt.Println("Error Opening file : ", err.Error())
		os.Exit(1)
	}

	reader := bufio.NewReader(file)

	persons := make([]person, 0, 10)

	for {
		line, _, err := reader.ReadLine()

		if err != nil {
			break
		}

		fields := strings.Fields(string(line))

		persons = append(persons, person{fname: fields[0], lname: fields[1]})
	}

	return persons
}

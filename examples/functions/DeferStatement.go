package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	test()

	readFile()

	deferTest()
}


func test() {
	fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")

	fmt.Println("Fourth")
}


func readFile() {

	file, _ := os.Open("/Users/in-sunit.chatterjee/go/src/LearningGo/examples/fileio/demo.txt")

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()

		if err != nil {
			break
		}

		fmt.Println(string(line))
	}
}


func deferTest() {
	i, j := 5, 10

	defer fmt.Println("First", i, j)


	i = i * 2
	j = j * 2

	fmt.Println("Last", i, j)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const inputdelimiter = '\n'

func main() {

	// create the reader to read from Standard input
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter an input string")

	// read entire line of string till a new line character
	input, err := reader.ReadString(inputdelimiter)

	// exit if any error
	if err != nil {
		fmt.Printf("Error occured : %s. Exiting the program\n", err.Error)
		os.Exit(1)
	}

	// Trim trailing space and also convert it to lower case for case insensitive match
	input = strings.TrimSpace(strings.ToLower(input))

	numChars := len(input)

	isFirstCharacterI := input[0] == 'i'
	isLastCharacterN := input[numChars-1] == 'n'
	containsCharA := strings.Contains(input, "a")

	matchFound := isFirstCharacterI && isLastCharacterN && containsCharA

	if matchFound {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}

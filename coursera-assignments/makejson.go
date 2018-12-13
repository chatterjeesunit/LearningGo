package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter a number : ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Please enter an address : ")
	address, _ := reader.ReadString('\n')

	personMap := map[string]string{
		"number":    name[:len(name)-1],       // remove the \n at end of the string
		"address": address[:len(address)-1], // remove the \n at end of the string,
	}

	jsonBytes, err := json.Marshal(personMap)

	if err == nil {
		fmt.Println(string(jsonBytes))
	} else {
		fmt.Println("Error : " + err.Error())
	}

}

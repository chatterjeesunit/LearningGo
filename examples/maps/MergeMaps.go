package main

import "fmt"

func main() {

	bigmap := map[string]string{"a":"a", "b":"b", "c":"c"}
	smallmap := map[string]string{"d":"d"}

	thirdmap := addmap(bigmap, smallmap)

	fmt.Println(bigmap)
	fmt.Println(smallmap)
	fmt.Println(thirdmap)
}


func addmap(a map[string]string, b map[string]string) map[string]string{
	for k,v := range a {
		b[k] = v
	}

	return b
}
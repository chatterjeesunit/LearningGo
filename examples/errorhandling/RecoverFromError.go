package main

import (
	"fmt"
	"time"
)

func main() {

	go demo()


	time.Sleep(5 * time.Second)

	fmt.Println("done")

}

func demo() {
	//defer RecoverFromPanic("Object is nil")
	println(&Obj{num: 1, data:[]string{"a", "b"}})

	println(nil)

	println(&Obj{num: 2, data:[]string{"c", "d"}})
}


func println(o *Obj) {

	defer RecoverFromPanic("Object is nil")

	for i := range o.data {
		fmt.Println(o.data[i])
	}

	fmt.Println(o.num)

}





type Obj struct {
	data []string
	num int
}


func RecoverFromPanic(msg string) {
	if err := recover(); err != nil {
		fmt.Printf("ERROR : err = %v, msg = %v\n", err, msg )
	}
}
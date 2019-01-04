package main

import (
	"LearningGo/examples/packages/p1"
)

func main() {

	p1.Sum(5, 10)

	p1.add(5, 10)  // Will throw compile time error - Unresolve reference.
	
}

package p1

//Name begins with lowercase. Cannot access outside package
func add(x, y int) int {
	return x + y
}

//Name begins with uppercase. Can access outside package
func Sum(x, y int) int {
	return x + y
}

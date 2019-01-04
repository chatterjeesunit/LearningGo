package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"sync"
)

var (
	wg = sync.WaitGroup{}
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	//fmt.Println(t.Value)
	ch <- t.Value
	Walk(t.Left, ch)
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	ch1 := make(chan int, 20)
	ch2 := make(chan int, 20)

	wg.Add(2)

	go func() {
		Walk(t1, ch1)
		wg.Done()
	}()

	go func() {
		Walk(t2, ch2)
		wg.Done()
	}()

	wg.Wait()

	close(ch1)
	close(ch2)

	for {
		x, xOk := <- ch1
		y, yOK := <- ch2

		if xOk == yOK && x == y {

			if xOk {
				continue
			} else {
				break
			}

		}
		return false
	}
	return true
}

func main() {

	t1 := tree.New(1)
	t2 := tree.New(3)

	fmt.Printf("T1 == T1 : %t\n", Same(t1, t1))
	fmt.Printf("T2 == T1 : %t\n", Same(t2, t1))


}

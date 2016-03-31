package main

import "golang.org/x/tour/tree"
import "fmt"

func walkTree(t *tree.Tree, ch chan int) {

	if t.Left != nil {
		walkTree(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walkTree(t.Right, ch)
	}
}

func walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	walkTree(t, ch)
}

func same(t1 *tree.Tree, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go walk(t1, ch1)
	go walk(t1, ch2)

	for i := range ch1 {
		j, ok := <-ch2
		if i != j || !ok {
			return false
		}
	}
	return true
}

func main() {
	t1 := tree.New(10)
	t2 := tree.New(10)
	result := same(t1, t2)
	fmt.Println(result)
}

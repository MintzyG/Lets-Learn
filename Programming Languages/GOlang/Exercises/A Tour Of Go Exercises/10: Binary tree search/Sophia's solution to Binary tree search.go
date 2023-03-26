package _0__Binary_tree_search

/*

I tried but found it hard, didn't manage to do it

Code doesn't run in an IDE

import "golang.org/x/tour/tree"

func Walk(t *tree.Tree, c chan int) {

	if t != nil {
		Walk(t.Left, c)
		c <- t.Value
		Walk(t.Right, c)
	}
}

func Same(t1, t2 *tree.Tree) bool {

	c1 := make(chan int)
	c2 := make(chan int)

	go Walk(t1, c1)1
	close(c1)
	go Walk(t2, c2)
	close(c2)

	for {
	v1, ok1 := <- c1
	v2, ok2 := <- c2

		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 {
			break
		}
	}

	return true

}

func main() {
	Same(tree.New(1), tree.New(1))
}
*/

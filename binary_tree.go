package gotour

import (
	"golang.org/x/tour/tree"
)

// type tree.Tree struct {
//     Left  *Tree
//     Value int
//     Right *Tree
// }

// Walk -- walks the tree t sending all values from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value

	if t.Left != nil {
		go Walk(t.Left, ch)
	}

	if t.Right != nil {
		go Walk(t.Right, ch)
	}
}

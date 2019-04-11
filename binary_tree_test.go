package gotour

import (
	"fmt"
	"testing"

	"golang.org/x/tour/tree"
)

func TestWalk(t *testing.T) {
	ch := make(chan int, 10)

	go Walk(tree.New(1), ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

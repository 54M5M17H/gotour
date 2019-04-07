package gotour

import "fmt"

func fibonacciGenerator() func() int {
	secondPrevious, previous := -1, 0

	return func() int {
		var new int
		if secondPrevious == -1 {
			new = 0
		} else if previous == 0 && secondPrevious == 0 {
			new = 1
		} else {
			new = previous + secondPrevious
		}

		secondPrevious, previous = previous, new

		return new
	}
}

// FibonacciPrinter print fib sequence
func FibonacciPrinter() {
	fibonacciCounter := fibonacciGenerator()
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacciCounter())
	}
}

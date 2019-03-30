package	gotour 

import (
	"testing"
	"fmt"
)

func TestSquareRoot(t *testing.T) {
	testCases := []struct { input, expected float64 } {
		{ 9, 3 },
		{ 100, 10 },
		{ 144, 12 },
	}

	for _, testCase := range testCases {
		result := SquareRoot(testCase.input)
		if result != testCase.expected {
			t.Errorf("Expected %f but got %f", testCase.expected, result);
		} else {
			fmt.Printf("Square of %f === %f \n", testCase.input, result)
		}
	}
}


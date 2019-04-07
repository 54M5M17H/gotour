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
		result, err := SquareRoot(testCase.input)
		if err != nil {
			t.Errorf("Unexpected error ")
		}

		if result != testCase.expected {
			t.Errorf("Expected %f but got %f", testCase.expected, result);
		} else {
			fmt.Printf("Square of %f === %f \n", testCase.input, result)
		}
	}

	_, err := SquareRoot(-2)
	if err.Error() != "Cannot determine the Square Root of a negative number: -2" {
		t.Errorf("Error not returned correctly")
	}
}

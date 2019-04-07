package gotour

import (
	"strings"
	"testing"
)

func TestFibonacci(t *testing.T) {
	output := StdoutSpy(FibonacciPrinter)
	sequence := strings.Fields(output)
	if sequence[0] != "0" {
		t.Errorf("Invalid first element of sequence")
	}
	if sequence[1] != "1" {
		t.Errorf("Invalid second element of sequence")
	}
	if sequence[2] != "1" {
		t.Errorf("Invalid third element of sequence")
	}
	if sequence[3] != "2" {
		t.Errorf("Invalid fourth element of sequence")
	}
	if sequence[4] != "3" {
		t.Errorf("Invalid fifth element of sequence")
	}
	if sequence[5] != "5" {
		t.Errorf("Invalid sixth element of sequence")
	}
}

package gotour

import "strconv"

// ErrNegativeSqrt -- error type for SquareRoot
type ErrNegativeSqrt float64

// SquareRoot --
func SquareRoot(x float64) (guess float64, err error) {
	if x < 0 {
		err = ErrNegativeSqrt(x)
		return
	}
	guess = 1;
	for i := 0; i < 10; i++ {
		diff := (guess * guess) - x;
		if diff == 0 {
			break
		}
		guess -= diff / (2 * guess)
	}
	return
}

func (e ErrNegativeSqrt) Error() string {
	return "Cannot determine the Square Root of a negative number: " + strconv.Itoa(int(e))
}

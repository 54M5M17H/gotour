package gotour 

func SquareRoot(x float64) (guess float64) {
	guess = 1;
	for i := 0; i < 10; i++ {
		diff := (guess * guess) - x;
		if diff == 0 {
			break
		}
		guess -= diff / (2 * guess)
	}
	return guess;
}


package gotour

import (
	"strings"
)

func WordCount(sentence string) map[string]int {
	wordCounter := make(map[string]int)
	words := strings.Fields(sentence)

	for _, word := range words {
		count, exists := wordCounter[word]
		if !exists {
			wordCounter[word] = 1
		} else {
			wordCounter[word] = count + 1
		}
	}

	return wordCounter
}

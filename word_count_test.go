package gotour

import (
	"testing"
)

func TestWordCount(t *testing.T) {
	sentence1 := "Hello this is a sentence where this is used as a confusing pronoun"
	sentence1WordCount := WordCount(sentence1)
	if sentence1WordCount["this"] != 2 {
		t.Errorf("Expected 2 but got %d \n", sentence1WordCount["this"])
	}

	if sentence1WordCount["where"] != 1 {
		t.Errorf("Expected 2 but got %d \n", sentence1WordCount["where"])
	}
}

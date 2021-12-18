package parts

import (
	"testing"
)

func TestSolutionTwo(t *testing.T) {
	fileName := "./test-input.txt"

	expectedAnswer := 5
	increases := SolutionTwo(fileName)

	if increases != expectedAnswer {
		t.Errorf("Got: %d; wanted %d", increases, expectedAnswer)
	} else {
		t.Logf("Got the correct increase count (using sliding windows)!")
	}
}

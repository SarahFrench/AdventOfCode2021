package parts

import (
	"testing"
)

func TestSolutionOne(t *testing.T) {
	fileName := "./test-input.txt"

	expectedAnswer := 7
	increases := SolutionOne(fileName)

	if increases != expectedAnswer {
		t.Errorf("Got: %d; wanted %d", increases, expectedAnswer)
	} else {
		t.Logf("Got the correct increase count!")
	}
}

package parts

import (
	"testing"
)

func TestSolutionTwo(t *testing.T) {
	fileName := "./test-input.txt"

	expectedScore := 1924
	finalScore := SolutionTwo(fileName)

	if finalScore != expectedScore {
		t.Errorf("Incorrect score. Got: %d; wanted %d", finalScore, expectedScore)
	} else {
		t.Logf("Got the correct score!")
	}
}

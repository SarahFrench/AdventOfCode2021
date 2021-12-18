package parts

import (
	"testing"
)

func TestSolutionOne(t *testing.T) {
	fileName := "./test-input.txt"
	
	expectedScore := 4512
	finalScore := SolutionOne(fileName)

	if finalScore != expectedScore {
		t.Errorf("Incorrect score. Got: %d; wanted %d", finalScore, expectedScore)
	} else {
		t.Logf("Got the correct score!")
	}
}

package parts

import (
	"testing"
)

func TestSolutionOne(t *testing.T) {
	fileName := "./test-1.txt"
	expectedAnswer := 150
	output := SolutionOne(fileName)
    if output != expectedAnswer {
        t.Errorf("Got: %d; wanted %d", output, expectedAnswer)
    } else {
		t.Logf("Got the correct answer for the test case!")
	}
}
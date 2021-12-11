package parts

import (
	"testing"
)

func TestSolutionTwo(t *testing.T) {
	fileName := "./test-2.txt"
	expectedAnswer := 900
	output := SolutionTwo(fileName)
    if output != expectedAnswer {
        t.Errorf("Got: %d; wanted %d", output, expectedAnswer)
    } else {
		t.Logf("Got the correct answer for the test case!")
	}
}
package parts

import (
	"testing"
)

func TestSolutionOne(t *testing.T) {
	fileName := "./test-input.txt"
	gammaDecimal := 22
	epsilonDecimal := 9
	expectedAnswer := 198

	gamma, epsilon, powerConsumption := SolutionOne(fileName)

	if gamma != gammaDecimal {
		t.Errorf("Got: %d; wanted %d", gamma, gammaDecimal)
	} else {
		t.Logf("Got the correct gamma value!")
	}
	if epsilon != epsilonDecimal {
		t.Errorf("Got: %d; wanted %d", epsilon, epsilonDecimal)
	} else {
		t.Logf("Got the correct epsilon value!")
	}
	if powerConsumption != expectedAnswer {
		t.Errorf("Got: %d; wanted %d", powerConsumption, expectedAnswer)
	} else {
		t.Logf("Got the correct power consumption value!")
	}
}

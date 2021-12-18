package parts

import (
	"testing"
)

func TestSolutionTwo(t *testing.T) {
	fileName := "./test-input.txt"

	expectedOxygen := 23
	expectedCO2 := 10
	expectedLifeSupport := 230

	oxygenGeneratorRating, co2ScrubberRating := SolutionTwo(fileName)

	if oxygenGeneratorRating != expectedOxygen {
		t.Errorf("Got: %d; wanted %d", oxygenGeneratorRating, expectedOxygen)
	} else {
		t.Logf("Got the correct oxygen generator rating value!")
	}
	if co2ScrubberRating != expectedCO2 {
		t.Errorf("Got: %d; wanted %d", co2ScrubberRating, expectedCO2)
	} else {
		t.Logf("Got the correct CO2 scrubber rating value!")
	}
	if oxygenGeneratorRating*co2ScrubberRating != expectedLifeSupport {
		t.Errorf("Got: %d; wanted %d", oxygenGeneratorRating*co2ScrubberRating, expectedLifeSupport)
	} else {
		t.Logf("Got the correct life support rating value!")
	}

}

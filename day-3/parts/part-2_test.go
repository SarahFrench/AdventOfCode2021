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
		t.Errorf("Incorrect oxygen generator rating\nGot: %d; wanted %d", oxygenGeneratorRating, expectedOxygen)
	} else {
		t.Logf("Got the correct oxygen generator rating value!")
	}
	if co2ScrubberRating != expectedCO2 {
		t.Errorf("Incorrect CO2 scrubber rating\nGot: %d; wanted %d", co2ScrubberRating, expectedCO2)
	} else {
		t.Logf("Got the correct CO2 scrubber rating value!")
	}
	if oxygenGeneratorRating*co2ScrubberRating != expectedLifeSupport {
		t.Errorf("Incorrect life support rating\nGot: %d; wanted %d", oxygenGeneratorRating*co2ScrubberRating, expectedLifeSupport)
	} else {
		t.Logf("Got the correct life support rating value!")
	}

}

func TestMakeEliminatorFunction(t *testing.T) {

	testCases := map[string]struct {
		tie                   bool
		tieDecider            int
		eliminateOnMostCommon bool
		mostCommonBit         int
		value                 int
		expectedFuncOutput    bool
	}{
		"In a tie, when a position's value matches the tie decider the input is kept": {
			tie:                   true,
			tieDecider:            1,
			eliminateOnMostCommon: true,
			mostCommonBit:         1,
			value:                 1,
			expectedFuncOutput:    false,
		},
		"In a tie, when a position's value DOESN'T match the tie decider the input is eliminated": {
			tie:                   true,
			tieDecider:            1,
			eliminateOnMostCommon: true,
			mostCommonBit:         1,
			value:                 0,
			expectedFuncOutput:    true,
		},
		"When the value differs from the mostCommonBit and eliminateOnMostCommon is TRUE, the input is kept": {
			tie:                   false,
			tieDecider:            1,
			eliminateOnMostCommon: true,
			mostCommonBit:         1,
			value:                 0,
			expectedFuncOutput:    false,
		},
		"When the value matches the mostCommonBit and eliminateOnMostCommon is TRUE, the input is eliminated ": {
			tie:                   false,
			tieDecider:            1,
			eliminateOnMostCommon: true,
			mostCommonBit:         1,
			value:                 1,
			expectedFuncOutput:    true,
		},
		"When the value matches the mostCommonBit and eliminateOnMostCommon is FALSE, the input is kept": {
			tie:                   false,
			tieDecider:            1,
			eliminateOnMostCommon: false,
			mostCommonBit:         1,
			value:                 1,
			expectedFuncOutput:    false,
		},
		"When the value differs from the mostCommonBit and eliminateOnMostCommon is FALSE, the input is eliminated": {
			tie:                   false,
			tieDecider:            1,
			eliminateOnMostCommon: false,
			mostCommonBit:         1,
			value:                 0,
			expectedFuncOutput:    true,
		},
	}
	for testName, tc := range testCases {
		outputFunc := makeEliminatorFunction(tc.eliminateOnMostCommon, tc.tieDecider)
		output := outputFunc(tc.tie, tc.mostCommonBit, tc.value)

		if output != tc.expectedFuncOutput {
			t.Errorf("FAIL: %s\nGot %t, expected %t", testName, output, tc.expectedFuncOutput)
		}
	}
}

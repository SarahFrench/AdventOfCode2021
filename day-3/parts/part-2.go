package parts

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	eliminated bool
	bits       []int
	original   string
}

func getInputsCopy(originalMap, newMap map[int]Input) map[int]Input {
	for k,v := range originalMap {
		newMap[k] = v
	  }
	return newMap
}

func makeEliminatorFunction(eliminateOnMostCommon bool, tieDecider int) func(bool, int, int) bool {
	return func(tie bool, mostCommonBit, value int) bool {
		// Handle tie situations
		if tie && value != tieDecider {
			return true
		}
		// Handle decisions based on most common bit at position
		if !tie && eliminateOnMostCommon && (value == mostCommonBit) {
			return true
		}
		if !tie && !eliminateOnMostCommon && (value != mostCommonBit) {
			return true
		}
		return false
	}
}

func processInputs(inputs map[int]Input, shouldEliminate func(bool, int, int) bool) (rating int) {
	for position := 0; position < len(inputs[0].bits); position++ {

		// Create tally of bits, identify tie or winning bit
		tally := map[int]int{
			0: 0,
			1: 0,
		}

		for k, v := range inputs {
			if inputs[k].eliminated {
				// Skip already-eliminated inputs
				continue
			}

			tally[v.bits[position]] += 1
		}

		mostCommonBit := -1
		tie := false
		if tally[0] > tally[1] {
			mostCommonBit = 0
		}
		if tally[0] < tally[1] {
			mostCommonBit = 1
		}
		tie = tally[0] == tally[1]

		if tally[0]+tally[1] == 1 {
			// Only one considered input left; abort
			break
		}

		// Eliminate inputs that fail conditions for this position
		for k, v := range inputs {
			if inputs[k].eliminated {
				// Skip already-eliminated inputs
				continue
			}

			eliminated := shouldEliminate(tie, mostCommonBit, v.bits[position])
			v.eliminated = eliminated
			inputs[k] = v
		}
	}

	// Identify non-eliminated values
	for _, v := range inputs {
		if !v.eliminated {
			decimal, err := strconv.ParseInt(v.original, 2, 64)
			if err != nil {
				log.Fatalf("cannot convert %s to an int", v.original)
			}
			rating = int(decimal)
		}
	}
	return
}

func SolutionTwo(fileName string) (oxygenGeneratorRating int, co2ScrubberRating int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputIndex := 0
	inputs := map[int]Input{}

	// Get a slice with each input prepared for use later
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, "")
		var bits []int

		for _, number := range split {
			integer, err := strconv.Atoi(number)
			if err != nil {
				log.Fatalf("cannot convert %s from line %s to an int", number, line)
			}
			bits = append(bits, integer)
		}
		if err != nil {
			log.Fatal(err)
		}
		// Add new entry for this line of input
		inputs[inputIndex] = Input{
			bits:     bits,
			original: line,
		}

		inputIndex++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Get oxygen generator rating
	// "...determine the most common value (0 or 1) in the current bit position, and keep only numbers with that bit in that position."
	// "0 and 1 are equally common, keep values with a 1 in the position being considered"
	oxygenEliminator := makeEliminatorFunction(false, 1)
	tmp := map[int]Input{} // avoid mutating original `input` map
	tmp = getInputsCopy(inputs, tmp)
	oxygenGeneratorRating = processInputs(tmp, oxygenEliminator)

	// Get co2 scrubber rating
	// "...determine the least common value (0 or 1) in the current bit position, and keep only numbers with that bit in that position."
	// "If 0 and 1 are equally common, keep values with a 0 in the position being considered"	
	co2Eliminator := makeEliminatorFunction(true, 0)

	tmp = getInputsCopy(inputs, tmp)
	co2ScrubberRating = processInputs(tmp, co2Eliminator)

	return
}

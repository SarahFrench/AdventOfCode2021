package parts

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	eliminatedO2  bool
	eliminatedCO2 bool
	bits          []int
	original      string
}

// func processInputs(inputs map[int]Input)

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

	// Find number with most popular bit at each position
	for position := 0; position < len(inputs[0].bits); position++ {

		// Create tally of bits, identify tie or winning bit
		tally := map[int]int{
			0: 0,
			1: 0,
		}

		for k, v := range inputs {
			if inputs[k].eliminatedO2 {
				// Skip already-eliminated inputs
				continue
			}

			tally[v.bits[position]] += 1
		}

		var winner int
		tie := false
		if tally[0] > tally[1] {
			winner = 0
		}
		if tally[0] < tally[1] {
			winner = 1
		}
		tie = tally[0] == tally[1]

		if tally[0]+tally[1] == 1 {
			// Only one contender left; abort
			break
		}

		// Eliminate inputs that fail conditions for this position
		for k, v := range inputs {
			if inputs[k].eliminatedO2 {
				// Skip already-eliminated inputs
				continue
			}

			if tie && v.bits[position] == 0 {
				v.eliminatedO2 = true
			}

			if !tie && v.bits[position] != winner {
				v.eliminatedO2 = true
			}
			inputs[k] = v
		}
	}
	for position := 0; position < len(inputs[0].bits); position++ {

		// Create tally of bits, identify tie or winning bit
		tally := map[int]int{
			0: 0,
			1: 0,
		}

		for k, v := range inputs {
			if inputs[k].eliminatedCO2 {
				// Skip already-eliminated inputs
				continue
			}

			tally[v.bits[position]] += 1
		}

		var winner int
		tie := false
		if tally[0] > tally[1] {
			winner = 0
		}
		if tally[0] < tally[1] {
			winner = 1
		}
		tie = tally[0] == tally[1]

		if tally[0]+tally[1] == 1 {
			// Only one contender left; abort
			break
		}

		// Eliminate inputs that fail conditions for this position
		for k, v := range inputs {
			if inputs[k].eliminatedCO2 {
				// Skip already-eliminated inputs
				continue
			}

			if tie && v.bits[position] == 1 {
				v.eliminatedCO2 = true
			}

			if !tie && v.bits[position] == winner {
				v.eliminatedCO2 = true
			}
			inputs[k] = v
		}
	}

	// Identify non-eliminated values
	for _, v := range inputs {
		if !v.eliminatedO2 {
			decimal, err := strconv.ParseInt(v.original, 2, 64)
			if err != nil {
				log.Fatalf("cannot convert %s to an int", v.original)
			}
			oxygenGeneratorRating = int(decimal)
		}
		if !v.eliminatedCO2 {
			decimal, err := strconv.ParseInt(v.original, 2, 64)
			if err != nil {
				log.Fatalf("cannot convert %s to an int", v.original)
			}
			co2ScrubberRating = int(decimal)
		}
		if oxygenGeneratorRating > 0 && co2ScrubberRating > 0 {
			// Both found
			break
		}
	}

	return
}

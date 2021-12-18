package parts

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func SolutionOne(fileName string) (int, int, int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputIndex := 0
	counts := []int{}

	// Get a slice with the sum of bits per position
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, "")

		if len(counts) == 0 {
			for i := 0; i < len(split); i++ {
				counts = append(counts, 0)
			}
		}

		for index, number := range split {
			integer, err := strconv.Atoi(number)
			if err != nil {
				log.Fatalf("cannot convert %s from line %s to an int", number, line)
			}
			counts[index] += integer
		}

		if err != nil {
			log.Fatal(err)
		}

		inputIndex++
	}

	// Create binary number string
	gammaString := ""
	epsilonString := ""
	for _, count := range counts {
		bit := count > (inputIndex / 2)
		if bit {
			gammaString = gammaString + "1"
			epsilonString = epsilonString + "0"
		} else {
			gammaString = gammaString + "0"
			epsilonString = epsilonString + "1"
		}
	}

	// Parse binary strings as ints
	gamma, err := strconv.ParseInt(gammaString, 2, 64)
	if err != nil {
		log.Fatalf("cannot convert %s to an int", gammaString)
	}
	epsilon, err := strconv.ParseInt(epsilonString, 2, 64)
	if err != nil {
		log.Fatalf("cannot convert %s to an int", epsilonString)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return int(gamma), int(epsilon), int(gamma) * int(epsilon)
}

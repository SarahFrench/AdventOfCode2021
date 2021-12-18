package parts

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func SolutionOne(fileName string) (int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var depthMeasurements []int
	index := 0
	increaseCount := 0
	for scanner.Scan() {
		// Pull next measurement from input file
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		// Compare to previous measurement if past 1st mesaurement
		if index > 0 {
			if number > depthMeasurements[index-1] {
				increaseCount++
			}
		}

		// Add measurement to record
		depthMeasurements = append(depthMeasurements, number)
		index++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return increaseCount
}
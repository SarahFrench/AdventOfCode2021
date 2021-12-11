package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
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

	log.Printf("Read %d measurements", index)
	log.Printf("Found %d increases on previous measurements", increaseCount)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

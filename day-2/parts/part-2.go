package parts

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)


func SolutionTwo(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	index := 0
	horizontal := 0
	depth := 0
	aim := 0
	for scanner.Scan() {
		// Pull next navigation step from input file
		line := scanner.Text()

		split := strings.Split(line, " ")
		direction := split[0]
		magnitude, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case "forward":
			horizontal += magnitude
			depth += (aim * magnitude)
		case "up":
			aim -= magnitude
		case "down":
			aim += magnitude
		default:
			log.Printf("unexpected direction of movement: %s", direction)
		}

		index++
	}

	log.Printf("Read %d navigation instructions", index)
	log.Printf("Final position, horizontal: %d , depth: %d", horizontal, depth)
	log.Printf("Product of horizontal & depth positions: %d", horizontal*depth)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return horizontal * depth
}

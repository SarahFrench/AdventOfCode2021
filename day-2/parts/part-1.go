package parts

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func SolutionOne(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	index := 0
	horizontal := 0
	depth := 0
	for scanner.Scan() {
		// Pull next navigation step from input file
		line := scanner.Text()

		split := strings.Split(line, " ")
		direction := split[0]
		distance, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case "forward":
			horizontal += distance
		case "up":
			depth -= distance
		case "down":
			depth += distance
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

package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const NUM_WINDOWS = 3
const WINDOW_SIZE = 3

type WindowGroup struct {
	Windows []Window
}

func NewWindowGroup(numWindows, windowSize int) WindowGroup {
	wg := WindowGroup{}
	for i := 0; i < numWindows; i++ {
		wg.Windows = append(wg.Windows, Window{
			number: i + 1,
			offset: i,
			size:   windowSize,
		})
		log.Printf("[WINDOW GROUP] Added the #%d Window to the group, with size %d and offset %d", i+1, windowSize, i)
	}
	return wg
}

type Window struct {
	number  int
	offset  int
	size    int
	Numbers []int
}

func (w *Window) Add(position, number int) (reportSum bool, sum int) {
	// Return early if window hasn't reached starting point
	if position < w.offset {
		log.Printf("[WINDOW #%d] Not ready to begin calculating using number at position %d because my offset is %d", w.number, position, w.offset)
		return reportSum, sum
	}

	// Add number
	w.Numbers = append(w.Numbers, number)

	// Reset window if reach size limit
	if len(w.Numbers) == w.size {
		// Get sum to return
		reportSum = true
		sum = w.getSum()
		// Reset window
		w.Numbers = []int{}

	}
	return reportSum, sum
}

func (w *Window) getSum() int {
	sum := 0
	for _, num := range w.Numbers {
		sum = sum + num
	}
	return sum
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	index := 0
	var increaseCount int
	var windowSums []int
	wg := NewWindowGroup(NUM_WINDOWS, WINDOW_SIZE)

	for scanner.Scan() {
		// Pull next measurement from input file
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		// Pass measurement to each window
		for i := 0; i < len(wg.Windows); i++ {
			addSum, sum := wg.Windows[i].Add(index, number)
			if addSum {
				// Check if new sum is larger than the last
				windowSumCount := len(windowSums)
				if windowSumCount > 0 && sum > windowSums[windowSumCount-1] {
					increaseCount++
				}
				// Record window sum
				windowSums = append(windowSums, sum)
			}
		}

		index++
	}

	log.Printf("Read %d measurements", index)
	log.Printf("Found %d increases on previous measurements", increaseCount)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

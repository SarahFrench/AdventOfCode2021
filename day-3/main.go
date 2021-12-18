package main

import (
	"log"

	parts "github.com/SarahFrench/AdventOfCode2021/day-3/parts"
)

func main() {
	// Part 1
	log.Printf("Part 1")
	gamma, epsilon, powerConsumption := parts.SolutionOne("./input.txt")
	log.Printf("\t answer: gamma %d, epsilon %d, power consumption %d\n", gamma, epsilon, powerConsumption)
	// Part 2
	log.Printf("Part 2")
	oxygenGeneratorRating, co2ScrubberRating := parts.SolutionTwo("./input.txt")
	lifeSupportRating := oxygenGeneratorRating * co2ScrubberRating
	log.Printf("\t answer: oxygen generator rating %d, CO2 scrubber rating %d , life support rating %d\n", oxygenGeneratorRating, co2ScrubberRating, lifeSupportRating)
}

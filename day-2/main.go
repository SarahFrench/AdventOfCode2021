package main

import (
	"log"

	parts "github.com/SarahFrench/AdventOfCode2021/day-2/parts"
)

func main() {
	// Part 1
	log.Printf("Part 1")
	output := parts.SolutionOne("./input.txt")
	log.Printf("\tanswer: %d", output)

	// Part 2
	log.Printf("Part 2")
	output = parts.SolutionTwo("./input.txt")
	log.Printf("\tanswer: %d", output)
}

package main

import (
	"log"

	"github.com/SarahFrench/AdventOfCode2021/day-1/parts"
)

func main() {
	// Part 1
	log.Printf("Part 1")
	increaseCount := parts.SolutionOne("./input.txt")
	log.Printf("Found %d increases on previous measurements", increaseCount)
	// Part 2
	log.Printf("Part 2")
	increaseCount = parts.SolutionTwo("./input.txt")
	log.Printf("Found %d increases on previous measurements", increaseCount)
}

package main

import (
	"log"

	parts "github.com/SarahFrench/AdventOfCode2021/day-4/parts"
)

func main() {
	// Part 1
	log.Printf("Part 1")
	finalScore := parts.SolutionOne("./input.txt")
	log.Printf("\t answer: score %d\n", finalScore)
	// Part 2
	log.Printf("Part 2")
	finalScore = parts.SolutionTwo("./input.txt")
	log.Printf("\t answer: score %d\n", finalScore)
}

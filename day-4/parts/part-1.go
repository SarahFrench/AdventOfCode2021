package parts

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type BingoCard struct {
	rows   [][]Number
	hasWon bool
}

type Number struct {
	value int
	mark  bool
}

func (b *BingoCard) markNumber(number int) {
	// For each row
	for i := 0; i < len(b.rows); i++ {
		// For each column in the row
		for j := 0; j < len(b.rows[i]); j++ {
			if b.rows[i][j].value == number {
				b.rows[i][j].mark = true
			}
		}
	}
}

func (b *BingoCard) isRowCompletelyMarked(rowNumber int) bool {
	for i := 0; i < len(b.rows[rowNumber]); i++ {
		if b.rows[rowNumber][i].mark == false {
			// An unmarked number
			return false
		}
	}
	return true
}

func (b *BingoCard) isColumnCompletelyMarked(column int) bool {
	for i := 0; i < len(b.rows); i++ {
		if b.rows[i][column].mark == false {
			// An unmarked number
			return false
		}
	}
	return true
}

func (b *BingoCard) hasWinningLine() bool {
	// Each row
	for i := 0; i < len(b.rows); i++ {
		verdit := b.isRowCompletelyMarked(i)
		if verdit {
			return true
		}
	}
	// Each column
	for i := 0; i < len(b.rows[0]); i++ {
		verdit := b.isColumnCompletelyMarked(i)
		if verdit {
			return true
		}
	}
	return false
}

func getIntegers(line string) []int {
	r := regexp.MustCompile(`\b\d+\b`) // regex for numbers
	digits := r.FindAllString(line, -1)
	if digits == nil {
		log.Fatalf("no numbers found in input line: %s", line)
	}
	numbers := []int{}
	for i := 0; i < len(digits); i++ {
		value, err := strconv.ParseInt(digits[i], 10, 64)
		if err != nil {
			log.Fatalf("unable to parse string to int: %s", digits[i])
		}
		numbers = append(numbers, int(value))

	}

	return numbers
}

func SolutionOne(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Get the called numbers from first line of input
	scanner.Scan()
	line := scanner.Text()
	calledNumbers := getIntegers(line)

	currentCard := BingoCard{}
	cards := []BingoCard{}

	// Parse the bingo cards
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			// Store finished card and make new blank card to fill in
			if len(currentCard.rows) > 0 {
				cards = append(cards, currentCard)
			}
			currentCard = BingoCard{}
			continue
		}
		digits := getIntegers(line)

		// Convert strings into ints
		rowNumbers := []Number{}
		for i := 0; i < len(digits); i++ {
			num := Number{
				value: digits[i],
			}
			rowNumbers = append(rowNumbers, num)
		}

		currentCard.rows = append(currentCard.rows, rowNumbers)
	}
	// Append last card of input
	cards = append(cards, currentCard)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Call numbers and mark cards until winner found
	winningCard := BingoCard{}
	lastNumberCalled := 0
	for i := 0; i < len(calledNumbers); i++ {
		lastNumberCalled = calledNumbers[i]
		// Check each card after each number called
		hasWon := false
		for j := 0; j < len(cards); j++ {
			cards[j].markNumber(lastNumberCalled)
			hasWon = cards[j].hasWinningLine()
			if hasWon {
				cards[j].hasWon = true
				winningCard = cards[j]
				break
			}
		}
		if hasWon {
			fmt.Printf("Bingo!")
			break
		}
	}

	// Calculate sum of all unmarked numbers
	sum := 0
	for i := 0; i < len(winningCard.rows); i++ {
		for j := 0; j < len(winningCard.rows); j++ {
			number := winningCard.rows[i][j]
			if !number.mark {
				sum += number.value
			}
		}
	}

	// return product of sum and last called number
	return sum * lastNumberCalled
}

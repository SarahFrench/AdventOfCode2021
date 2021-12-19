package parts

import (
	"bufio"
	"log"
	"os"
)

func allCardsWon(cards []BingoCard) bool {
	for _, card := range cards {
		if card.hasWon == false {
			return false
		}
	}
	return true
}

func SolutionTwo(fileName string) int {
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
	lastNumberCalled := 0
	lastCard := BingoCard{}
	for i := 0; i < len(calledNumbers); i++ {
		lastNumberCalled = calledNumbers[i]
		// Check each card after each number called
		hasWon := false
		for j := 0; j < len(cards); j++ {
			cards[j].markNumber(lastNumberCalled)
			hasWon = cards[j].hasWinningLine()
			if hasWon {
				cards[j].hasWon = true
				if allCardsWon(cards) {
					lastCard = cards[j]
					break
				}
			}
		}
		if allCardsWon(cards) {
			break
		}
	}

	// Calculate sum of all unmarked numbers
	sum := 0
	for i := 0; i < len(lastCard.rows); i++ {
		for j := 0; j < len(lastCard.rows); j++ {
			number := lastCard.rows[i][j]
			if !number.mark {
				sum += number.value
			}
		}
	}

	// return product of sum and last called number
	return sum * lastNumberCalled
}

package parts

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type BingoCard struct {
	lines []  string
}

func SolutionOne(fileName string) (int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	// Get the called numbers from first line of input
	scanner.Scan()
	line := scanner.Text()
	fmt.Println(line)

	currentCard := BingoCard{}
	cards := []BingoCard{}

	// Parse the bingo cards
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			// Store finished card and make new blank card to fill in
			cards = append(cards, currentCard)
			currentCard = BingoCard{}
			continue
		}
		
		currentCard.lines = append(currentCard.lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return 0
}
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	LOW = iota
	EQUAL
	HIGH
)

func main() {
	welcomeMsg := `
-------------------------------------
Welcome to my blackjack!
To exit the game, press the Q key.
To draw a card, press the Enter key.
Let's get it on!!!
-------------------------------------
`
	fmt.Println(welcomeMsg)

	cards := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)

	var total int
	for scanner.Scan() {
		s := scanner.Text()
		if s == "q" || s == "Q" {
			fmt.Println("OK, See you again.")
			return
		}
		cards = append(cards, drawCard())
		total = calcCards(cards)
		fmt.Println(fmtStatus(cards, total))

		if total == 21 {
			fmt.Println("You Win !!!")
			return
		}

		if 21 < total {
			fmt.Println("You Lose...")
			return
		}
	}
}

func drawCard() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(12) + 1
}

func fmtCard(cards []int) string {
	cardStr := make([]string, len(cards))
	for _, card := range cards {
		switch card {
		case 1:
			cardStr = append(cardStr, " [A] ")
		case 11:
			cardStr = append(cardStr, " [J] ")
		case 12:
			cardStr = append(cardStr, " [Q] ")
		case 13:
			cardStr = append(cardStr, " [K] ")
		default:
			c := fmt.Sprintf(" [%d] ", card)
			cardStr = append(cardStr, c)
		}
	}
	return strings.Join(cardStr, "")
}

func calcCards(cards []int) int {
	var total int
	for _, card := range cards {
		total += card
	}
	return total
}

func fmtStatus(cards []int, total int) string {
	return fmt.Sprintf(`
-------Your cards-------
%s
Total: %d
------------------------
`, fmtCard(cards), total)
}

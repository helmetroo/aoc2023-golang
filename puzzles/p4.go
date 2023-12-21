package puzzles

import (
	"bufio"
	"strconv"
	"strings"

    "samjay/aoc2023-golang/utils"
)

func countWinningNumbers(line string) int {
    colonIdx := strings.Index(line, ":")
    winnersAndHadNumbers := strings.Split(line[colonIdx + 2:], "|")
    winners, hads := winnersAndHadNumbers[0], winnersAndHadNumbers[1]
    winnerSet, hadSet := utils.ToNumSet(winners), utils.ToNumSet(hads)
    return len(utils.Intersect(&winnerSet, &hadSet))
}

func scoreFrom(winCount int) int {
    switch winCount {
    case 0:
        return 0
    default:
        return 1 << (winCount - 1)
    }
}

type Card struct {
    copies int
    winners int
}

func P4_SolvePartOne(scanner *bufio.Scanner) (string, error) {
    score := 0
    for scanner.Scan() {
        line := scanner.Text()
        winnersCount := countWinningNumbers(line)
        score += scoreFrom(winnersCount)
    }

    return strconv.Itoa(score), nil
}

func P4_SolvePartTwo(scanner *bufio.Scanner) (string, error) {
    cards := []Card{}

    for scanner.Scan() {
        line := scanner.Text()
        winnersCount := countWinningNumbers(line)
        cards = append(cards, Card { copies: 1, winners: winnersCount })
    }

    totalCards := 0
    lenCards := len(cards)
    for index, card := range cards {
        for c := 0; c < card.copies; c++ {
            for w := 1; w <= card.winners; w++ {
                if index + w < lenCards {
                    cards[index + w].copies++
                }
            }
        }

        totalCards += card.copies
    }

    return strconv.Itoa(totalCards), nil
}

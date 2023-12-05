package puzzles

import (
	"bufio"
	"strconv"
	"strings"
)

func countWinningNumbers(line string) int {
    colonIdx := strings.Index(line, ":")
    winnersAndHadNumbers := strings.Split(line[colonIdx + 2:], "|")
    winners, hads := winnersAndHadNumbers[0], winnersAndHadNumbers[1]
    winnerSet, hadSet := toNumSet(winners), toNumSet(hads)
    return len(intersect(&winnerSet, &hadSet))
}

func toNumSet(nums string) map[int]struct{} {
    numSet := map[int]struct{}{}

    // NUMBER_REGEX declared and compiled already in P2
    numIndices := NUMBER_REGEX.FindAllStringIndex(nums, -1)
    for _, indices := range numIndices {
        firstDigit, lastDigit := indices[0], indices[1]
        num, _ := strconv.Atoi(nums[firstDigit:lastDigit])
        numSet[num] = struct{}{}
    }

    return numSet
}

func intersect(first *map[int]struct{}, second *map[int]struct{}) map[int]struct{} {
    intersection := map[int]struct{}{}

    // Iterate over the smaller one (the result set won't be any bigger than it!)
    if len(*first) > len(*second) {
        first, second = second, first
    }

    for num := range *first {
        if _, exists := (*second)[num]; exists {
            intersection[num] = struct{}{}
        }
    }

    return intersection
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

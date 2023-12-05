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
    return "", nil
}

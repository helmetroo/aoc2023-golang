package puzzles

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

var SET_REGEX = regexp.MustCompile(`(\d+) (red|blue|green)`)

func validGame(line string, allowed *map[string]int) bool {
    splitByTurn := strings.Split(line, ";")
    for _, turn := range splitByTurn {
        setMatches := SET_REGEX.FindAllStringSubmatch(turn, -1)
        for set := range setMatches {
            num, _ := strconv.Atoi(setMatches[set][1])
            color := setMatches[set][2]

            if num > (*allowed)[color] {
                return false
            }
        }
    }

    return true
}

func P2_SolvePartOne(scanner *bufio.Scanner) (string, error) {
    sumPossibleGameIds := 0
    currentGameId := 1
    allowed := map[string]int {
        "red": 12,
            "green": 13,
            "blue": 14,
    }

    for scanner.Scan() {
        line := scanner.Text()
        if validGame(line, &allowed) {
            sumPossibleGameIds += currentGameId
        }

        currentGameId++
    }

    return strconv.Itoa(sumPossibleGameIds), nil
}

func P2_SolvePartTwo(scanner *bufio.Scanner) (string, error) {
    return "", nil
}

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

func gamePower(line string) int {
    splitByTurn := strings.Split(line, ";")
    fewestCubes := map[string]int {
        "red": 0,
        "green": 0,
        "blue": 0,
    }

    for _, turn := range splitByTurn {
        setMatches := SET_REGEX.FindAllStringSubmatch(turn, -1)
        for set := range setMatches {
            num, _ := strconv.Atoi(setMatches[set][1])
            color := setMatches[set][2]

            if num > fewestCubes[color] {
                fewestCubes[color] = num
            }
        }
    }

    fewestReds := fewestCubes["red"]
    fewestGreens := fewestCubes["green"]
    fewestBlues := fewestCubes["blue"]
    return fewestReds * fewestGreens * fewestBlues
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
    sumPossibleGameIds := 0

    for scanner.Scan() {
        line := scanner.Text()
        sumPossibleGameIds += gamePower(line)
    }

    return strconv.Itoa(sumPossibleGameIds), nil
}

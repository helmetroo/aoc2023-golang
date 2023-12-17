package puzzles

import (
	"bufio"
	"strconv"
)

type Race struct {
    Time int
    RecordDist int
}

// Wins are symmetric after race.Time / 2, so we can just double to get the correct count!
func (race *Race) countWaysToWin() int {
    waysToWin := 0

    for heldTime := 0; heldTime <= race.Time / 2; heldTime++ {
        dist := heldTime * (race.Time - heldTime)
        if dist > race.RecordDist {
            waysToWin++
        }
    }

    if race.Time % 2 == 0 {
        return (waysToWin * 2) - 1
    } else {
        return waysToWin * 2
    }
}

func parseRaces(scanner *bufio.Scanner) ([]Race, error) {
    lineNum := 0
    times, recordDists := []string{}, []string{}
    races := []Race{}

    for scanner.Scan() {
        currentLine := scanner.Text()
        numsInLine := NUMBER_REGEX.FindAllString(currentLine, -1)

        if lineNum == 0 {
            times = numsInLine
        } else {
            recordDists = numsInLine
        }

        lineNum++
    }

    nRaces := len(times)
    for index := 0; index < nRaces; index++ {
        time, _ := strconv.Atoi(times[index])
        recordDist, _ := strconv.Atoi(recordDists[index])

        races = append(races, Race{ time, recordDist })
    }

    return races, nil
}

func P6_SolvePartOne(scanner *bufio.Scanner) (string, error) {
    races, _ := parseRaces(scanner)
    prod := 1
    for _, race := range races {
        prod *= race.countWaysToWin()
    }

    return strconv.Itoa(prod), nil
}

func P6_SolvePartTwo(scanner *bufio.Scanner) (string, error) {
    return "", nil
}

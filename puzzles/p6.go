package puzzles

import (
	"bufio"
	"strconv"
    "strings"
)

type Race struct {
    Time uint64
    RecordDist uint64
}

// Wins are symmetric after race.Time / 2, so we can just double to get the correct count!
func (race *Race) countWaysToWin() uint64 {
    waysToWin := uint64(0)

    for heldTime := race.Time / 2; heldTime >= 0; heldTime-- {
        dist := heldTime * (race.Time - heldTime)
        if dist <= race.RecordDist {
            break
        }

        waysToWin++
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
        time, _ := strconv.ParseUint(times[index], 10, 64)
        recordDist, _ := strconv.ParseUint(recordDists[index], 10, 64)

        races = append(races, Race{ time, recordDist })
    }

    return races, nil
}

func parseSmooshedRace(scanner *bufio.Scanner) (Race, error) {
    lineNum := 0
    time, recordDist := uint64(0), uint64(0)

    for scanner.Scan() {
        currentLine := scanner.Text()
        numsInLine := strings.Join(NUMBER_REGEX.FindAllString(currentLine, -1), "")
        numInLine, _ := strconv.ParseUint(numsInLine, 10, 64)

        if lineNum == 0 {
            time = numInLine
        } else {
            recordDist = numInLine
        }

        lineNum++
    }

    return Race { time, recordDist }, nil
}

func P6_SolvePartOne(scanner *bufio.Scanner) (string, error) {
    races, _ := parseRaces(scanner)
    prod := uint64(1)
    for _, race := range races {
        prod *= race.countWaysToWin()
    }

    return strconv.FormatUint(prod, 10), nil
}

func P6_SolvePartTwo(scanner *bufio.Scanner) (string, error) {
    race, _ := parseSmooshedRace(scanner)
    waysToWin := race.countWaysToWin()

    return strconv.FormatUint(waysToWin, 10), nil
}

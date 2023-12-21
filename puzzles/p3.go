package puzzles

import (
	"bufio"
	"strconv"

    "samjay/aoc2023-golang/utils"
)


func findAndSumSymAdjNums(grid *[]string) int {
    adjNumSum := 0
    rows, cols := len(*grid), len((*grid)[0])
    for row, rowLine := range *grid {
        numMatches := utils.NUMBER_REGEX.FindAllStringIndex(rowLine, -1)
        for _, numIndices := range numMatches {
            firstDigitCol, lastDigitCol := numIndices[0], numIndices[1] - 1
            firstDigitAdj := digitAdjacentToSym(grid, FIRST, row, firstDigitCol, rows, cols)
            lastDigitAdj := digitAdjacentToSym(grid, LAST, row, lastDigitCol, rows, cols)
            if firstDigitAdj || lastDigitAdj {
                adjNum, _ := strconv.Atoi(rowLine[firstDigitCol:lastDigitCol + 1])
                adjNumSum += adjNum
            }
        }
    }

    return adjNumSum
}

var (
    FIRST_OFFSETS = [][]int {
        {-1, -1},
            {-1, 0},
            {-1, 1},

            {0, -1},

            {1, -1},
            {1, 0},
            {1, 1},
        }

    LAST_OFFSETS = [][]int {
        {-1, -1},
            {-1, 0},
            {-1, 1},

            {0, 1},

            {1, -1},
            {1, 0},
            {1, 1},
        }
)

type DigitPos int
const (
    FIRST DigitPos = iota
    LAST
)

func digitAdjacentToSym(grid *[]string, d DigitPos, row, col, rows, cols int) bool {
    offsets := FIRST_OFFSETS
    if d == LAST {
        offsets = LAST_OFFSETS
    }

    for _, offset := range offsets {
        offsetRow, offsetCol := offset[0], offset[1]
        targetRow, targetCol := row + offsetRow, col + offsetCol
        if !utils.ValidOffset(targetRow, targetCol, rows, cols) {
            continue
        }

        chr := (*grid)[targetRow][targetCol]
        if (chr < '0' || chr > '9') && chr != '.' {
            return true
        }
    }

    return false
}

func P3_SolvePartOne(scanner *bufio.Scanner) (string, error) {
    grid := []string{}
    for scanner.Scan() {
        line := scanner.Text()
        grid = append(grid, line)
    }

    sumSymAdjNums := findAndSumSymAdjNums(&grid)
    return strconv.Itoa(sumSymAdjNums), nil
}

func P3_SolvePartTwo(scanner *bufio.Scanner) (string, error) {
    return "", nil
}

package puzzles

import (
	"bufio"
	"regexp"
	"strconv"
)

var NUMBER_WITH_NEG_REGEX = regexp.MustCompile(`-?\d+`)

type Histories [][][]int
func buildHistories(scanner *bufio.Scanner) Histories {
    histories := Histories{}

    for scanner.Scan() {
        curHistoryDiffs := [][]int{}
        curLine := scanner.Text()
        lineNumStrs := NUMBER_WITH_NEG_REGEX.FindAllString(curLine, -1)

        // Parse each first
        lineNums := []int{}
        for _, numStr := range lineNumStrs {
            num, _ := strconv.Atoi(numStr)
            lineNums = append(lineNums, num)
        }

        curSeq, curSeqLen, nextSeq := lineNums, len(lineNums), []int{}
        allZeroes := false

        // Build up diff sequences
        for !allZeroes {
            allZeroesThisSeq := true
            for idx := 0; idx < curSeqLen - 1; idx++ {
                first, second := curSeq[idx], curSeq[idx + 1]
                diff := second - first
                nextSeq = append(nextSeq, diff)

                if diff != 0 {
                    allZeroesThisSeq = false
                }
            }

            curHistoryDiffs = append(curHistoryDiffs, curSeq)
            curSeq, curSeqLen, nextSeq = nextSeq, len(nextSeq), []int{}
            allZeroes = allZeroesThisSeq
        }

        // Next history
        histories = append(histories, curHistoryDiffs)
    }

    return histories
}

func extrapolateValuesFromHistories(histories *Histories) int {
    sumExtrapolatedValues := 0

    for _, historyDiffs := range *histories {
        depth := len(historyDiffs) - 1

        for curDepth := depth; curDepth >= 0; curDepth-- {
            // Must reference the array we append to (in this case, curSeq)
            curSeq := &historyDiffs[curDepth]
            lenCurSeq := len(*curSeq)
            lastValCurSeq := (*curSeq)[lenCurSeq - 1]

            if curDepth != depth {
                // Above bottom sequence
                belowSeq := historyDiffs[curDepth + 1]
                lenBelowSeq := len(belowSeq)
                lastValBelowSeq := belowSeq[lenBelowSeq - 1]
                sum := lastValCurSeq + lastValBelowSeq

                *curSeq = append(*curSeq, sum)
                if curDepth == 0 {
                    sumExtrapolatedValues += sum
                }
            } else {
                // At the bottom sequence, lastValBelowSeq = 0
                *curSeq = append(*curSeq, lastValCurSeq)
            }
        }
    }

    return sumExtrapolatedValues
}

func extrapolateValuesFromHistoriesAtBeginning(histories *Histories) int {
    sumExtrapolatedValues := 0

    for _, historyDiffs := range *histories {
        depth := len(historyDiffs) - 1

        for curDepth := depth; curDepth >= 0; curDepth-- {
            // Must reference the array we append to (in this case, curSeq)
            curSeq := &historyDiffs[curDepth]
            lastValCurSeq := (*curSeq)[0]

            if curDepth != depth {
                // Above bottom sequence
                belowSeq := historyDiffs[curDepth + 1]
                lastValBelowSeq := belowSeq[0]
                diff := lastValCurSeq - lastValBelowSeq

                *curSeq = append([]int{ diff }, *curSeq...)
                if curDepth == 0 {
                    sumExtrapolatedValues += diff
                }
            } else {
                // At the bottom sequence, lastValBelowSeq = 0
                *curSeq = append([]int{ lastValCurSeq }, *curSeq...)
            }
        }
    }

    return sumExtrapolatedValues
}

func P9_SolvePartOne(scanner *bufio.Scanner) (string, error) {
    histories := buildHistories(scanner)
    sumExtrapolatedValues := extrapolateValuesFromHistories(&histories)

    return strconv.Itoa(sumExtrapolatedValues), nil
}

func P9_SolvePartTwo(scanner *bufio.Scanner) (string, error) {
    histories := buildHistories(scanner)
    sumExtrapolatedValues := extrapolateValuesFromHistoriesAtBeginning(&histories)

    return strconv.Itoa(sumExtrapolatedValues), nil
}

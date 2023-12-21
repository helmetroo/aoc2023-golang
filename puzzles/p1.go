package puzzles

import (
	"bufio"
	"math"
	"strconv"
	"strings"

    "samjay/aoc2023-golang/utils"
)

// A means to keep track of where each digit is
func digits(line []byte) (int, int) {
    digitIndex := 0
    digits := make(map[int]int)

    for _, char := range line {
        asInt := utils.CharToInt(char)
        if asInt >= 0 && asInt <= 9 {
            digits[digitIndex] = asInt
            digitIndex++
        }
    }

    return digits[0], digits[digitIndex - 1]
}

func wordDigits(line string) (int, int) {
    digitStrs := []string {
        "1",
            "2",
            "3",
            "4",
            "5",
            "6",
            "7",
            "8",
            "9",
            "one",
            "two",
            "three",
            "four",
            "five",
            "six",
            "seven",
            "eight",
            "nine",
        }

    wordsToDigits := map[string]int {
        "one": 1,
            "two": 2,
            "three": 3,
            "four": 4,
            "five": 5,
            "six": 6,
            "seven": 7,
            "eight": 8,
            "nine": 9,
        }

    firstDigit, lastDigit := 0, 0

    // Find a digit or word from the beginning and end
    firstWordIndex, firstWordOrDigit := math.MaxInt, "none"
    lastWordIndex, lastWordOrDigit := math.MaxInt, "none"
    lineRev := utils.ReverseString(line)
    for _, wordOrDigit := range digitStrs {
        wordIndex := strings.Index(line, wordOrDigit)
        revWordIndex := strings.Index(lineRev, utils.ReverseString(wordOrDigit))

        if wordIndex < firstWordIndex && wordIndex != -1 {
            firstWordIndex = wordIndex
            firstWordOrDigit = wordOrDigit
        }

        if revWordIndex < lastWordIndex && revWordIndex != -1 {
            lastWordIndex = revWordIndex
            lastWordOrDigit = wordOrDigit
        }
    }

    if len(firstWordOrDigit) == 1 {
        firstDigit = utils.CharToInt(firstWordOrDigit[0])
    } else {
        firstDigit = wordsToDigits[firstWordOrDigit]
    }

    if len(lastWordOrDigit) == 1 {
        lastDigit = utils.CharToInt(lastWordOrDigit[0])
    } else {
        lastDigit = wordsToDigits[lastWordOrDigit]
    }

    return firstDigit, lastDigit
}

// Treat the line as a line of ASCII chars
func P1_SolvePartOne(scanner *bufio.Scanner) (string, error) {
    sumCalibrationValues := 0
    for scanner.Scan() {
        line := scanner.Bytes()
        firstDigit, lastDigit := digits(line)

        sumCalibrationValues += 10*firstDigit + lastDigit
    }

    return strconv.Itoa(sumCalibrationValues), nil
}

func P1_SolvePartTwo(scanner *bufio.Scanner) (string, error) {
    sumCalibrationValues := 0
    for scanner.Scan() {
        line := scanner.Text()
        firstDigit, lastDigit := wordDigits(line)

        sumCalibrationValues += 10*firstDigit + lastDigit
    }

    return strconv.Itoa(sumCalibrationValues), nil
}

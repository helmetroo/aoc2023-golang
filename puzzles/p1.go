package puzzles

import (
	"bufio"
	"strconv"
)

func charToInt(char byte) int {
    return int(char) - '0'
}

// A means to keep track of where each digit is
func digits(line []byte) (int, int) {
    digitIndex := 0
    digits := make(map[int]int)

    for _, char := range line {
        asInt := charToInt(char)
        if asInt >= 0 && asInt <= 9 {
            digits[digitIndex] = asInt
            digitIndex++
        }
    }

    return digits[0], digits[digitIndex - 1]
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
    return "", nil
}

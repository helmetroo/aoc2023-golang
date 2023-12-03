package puzzles

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type InputFile struct {
    file *os.File
    scanner *bufio.Scanner
}

type Solvers [2]func(scanner *bufio.Scanner) (string, error)

var availablePuzzles = map[int]Solvers {
    1: { P1_SolvePartOne, P1_SolvePartTwo },
    2: { P2_SolvePartOne, P2_SolvePartTwo },
    }

func Solve(number int, partOneMode string, partTwoMode string) error {
    inputFiles, readErr := ReadInputFiles(number, partOneMode, partTwoMode)
    if readErr != nil {
        errorMessage := fmt.Sprintf(
            "Unable to read input file(s). (Reason: %s)",
            readErr.Error(),
        )

        return errors.New(errorMessage)
    }

    solvers, puzzleExists := availablePuzzles[number]
    if !puzzleExists {
        return errors.New("Puzzle not available yet!")
    }

    partOneScanner := inputFiles[partOneMode].scanner
    partOneSolution, partOneErr := solvers[0](partOneScanner)
    if partOneErr == nil {
        fmt.Printf("Solution for part one: %s\n", partOneSolution)
    } else {
        fmt.Printf("Failed to solve part one: %s\n", partOneErr.Error())
    }

    // Rewind if the scanner for input file 1 is the same as 2
    if partOneMode == partTwoMode {
        if secondInputFile, ok := inputFiles[partTwoMode]; ok {
            secondInputFile.file.Seek(0, io.SeekStart)
            secondInputFile.scanner = bufio.NewScanner(secondInputFile.file)

            inputFiles[partTwoMode] = secondInputFile
        }
    } else {
        inputFiles[partOneMode].file.Close()
    }

    partTwoScanner := inputFiles[partTwoMode].scanner
    partTwoSolution, partTwoErr := solvers[1](partTwoScanner)
    if partTwoErr == nil {
        fmt.Printf("Solution for part two: %s\n", partTwoSolution)
    } else {
        fmt.Printf("Failed to solve part two: %s\n", partTwoErr.Error())
    }

    inputFiles[partTwoMode].file.Close()
    return nil
}

func ReadInputFiles(number int, partOneMode string, partTwoMode string) (map[string]InputFile, error) {
    inputFiles := make(map[string]InputFile)
    if partOneMode == partTwoMode {
        scanner, readErr := OpenInputFile(number, partOneMode)
        if readErr != nil {
            return nil, readErr
        }
        inputFiles[partOneMode] = scanner
        return inputFiles, nil
    }

    scannerPartOne, readPartOneErr := OpenInputFile(number, partOneMode)
    scannerPartTwo, readPartTwoErr := OpenInputFile(number, partTwoMode)

    if readPartOneErr != nil || readPartTwoErr != nil {
        return nil, errors.Join(readPartOneErr, readPartTwoErr)
    }

    inputFiles[partOneMode] = scannerPartOne
    inputFiles[partTwoMode] = scannerPartTwo
    return inputFiles, nil
}

func OpenInputFile(number int, mode string) (InputFile, error) {
    wdPath, dirErr := os.Getwd()
    if dirErr != nil {
        return InputFile{}, dirErr
    }

    fileName := fmt.Sprintf("%d-%s.txt", number, mode)
    inputFilePath := filepath.Join(wdPath, "input-files", fileName)

    file, openErr := os.Open(inputFilePath)
    if openErr != nil {
        return InputFile{}, openErr
    }

    return InputFile {
        file: file,
            scanner: bufio.NewScanner(file),
    }, nil
}

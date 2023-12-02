package main

import (
    "errors"
    "flag"
    "fmt"
    "os"
    "strings"

    "samjay/aoc2023-golang/puzzles"
)

type Args struct {
    puzzleNum *int
    partOneMode *string
    partTwoMode *string
}

func main() {
    // Handle command line args
    args := parseArgs()
    validationErr := validateArgs(args)
    if validationErr != nil {
        fmt.Fprintln(os.Stderr, validationErr.Error())
        os.Exit(1)
    }

    solveErr := puzzles.Solve(
        *args.puzzleNum,
        *args.partOneMode,
        *args.partTwoMode,
    )
    if solveErr != nil {
        fmt.Fprintln(os.Stderr, solveErr.Error())
        os.Exit(1)
    }
}

func parseArgs() Args {
    // Define and parse cmd line args
    args := Args {
        puzzleNum: flag.Int("number", 1, "Which puzzle to solve"),
        partOneMode: flag.String(
            "part-one-mode",
            "test",
            "Which input file to use to solve part 1 (test or full)",
        ),
        partTwoMode: flag.String(
            "part-two-mode",
            "test",
            "Which input file to use to solve part 2 (test or full)",
        ),
    }

    flag.Parse()
    return args
}

func validateArgs(args Args) error {
    var messages []string
    if *args.puzzleNum < 1 || *args.puzzleNum > 25 {
        messages = append(messages, "Invalid puzzle number! Should be between 1 and 25.")
    }

    if *args.partOneMode != "test" && *args.partOneMode != "full" {
        messages = append(messages, "Invalid part one mode! Needs to be \"full\" or \"test\".")
    }

    if *args.partTwoMode != "test" && *args.partTwoMode != "full" {
        messages = append(messages, "Invalid part two mode! Needs to be \"full\" or \"test\".")
    }

    if len(messages) > 0 {
        message := strings.Join(messages[:], "\r\n")
        return errors.New(message)
    }

    return nil
}

package puzzles

import (
	"bufio"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
    start int64
    end int64
}

func (r Range) Includes(val int64) bool {
    return r.start <= val && val <= r.end
}

type Almanac struct {
    seeds []int64
    ranges [7]map[Range]Range
}

func parseSeeds(line string) []int64 {
    colonIdx := strings.Index(line, ":")
    seedStrs := strings.Split(line[colonIdx + 2:], " ")
    seeds := []int64{}
    for _, seedStr := range seedStrs {
        seed, _ := strconv.ParseInt(seedStr, 10, 64)
        seeds = append(seeds, seed)
    }
    return seeds
}

func parseAlmanac(scanner *bufio.Scanner) (Almanac, error) {
    parsedSeeds, parsingMap := false, false
    seeds, allRanges := []int64{}, [7]map[Range]Range{}
    index := 0

    for scanner.Scan() {
        line := scanner.Text()
        if !parsedSeeds {
            seeds = parseSeeds(line)
            parsedSeeds = true
        } else if(len(line) == 0) {
            if parsingMap {
                parsingMap = false
                index++
            }
        } else {
            if !parsingMap {
                parsingMap = true
            } else {
                rangeStrs := strings.Split(line, " ")
                destStart, _ := strconv.ParseInt(rangeStrs[0], 10, 64)
                srcStart, _ := strconv.ParseInt(rangeStrs[1], 10, 64)
                rangeLen, _ := strconv.ParseInt(rangeStrs[2], 10, 64)

                srcRange := Range { srcStart, srcStart + rangeLen - 1 }
                destRange := Range { destStart, destStart + rangeLen - 1 }
                if allRanges[index] == nil {
                    allRanges[index] = map[Range]Range{}
                }

                allRanges[index][srcRange] = destRange
            }
        }
    }

    return Almanac { seeds: seeds, ranges: allRanges }, nil
}

func getMinLocationNumber(almanac *Almanac) int64 {
    locations := []int64{}

    for _, seed := range almanac.seeds {
        for _, rangeMap := range almanac.ranges {
            for srcRange, destRange := range rangeMap {
                if srcRange.Includes(seed) {
                    offset := destRange.start - srcRange.start
                    seed += offset
                    break
                }
            }
        }

        locations = append(locations, seed)
    }

    return slices.Min(locations)
}

func P5_SolvePartOne(scanner *bufio.Scanner) (string, error) {
    almanac, _ := parseAlmanac(scanner)
    minLocNum := getMinLocationNumber(&almanac)

    return strconv.FormatInt(minLocNum, 10), nil
}

func P5_SolvePartTwo(scanner *bufio.Scanner) (string, error) {
    return "", nil
}

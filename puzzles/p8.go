package puzzles

import (
	"bufio"
	"regexp"
	"strconv"
)

// [0] = node; [1] = node's left; [2] = node's right
var NODE_DEFINITION_REGEX = regexp.MustCompile(`\w{3}`)

type Adjacencies struct {
    Left string
    Right string
}

type Network map[string]Adjacencies

func parseNetwork(scanner *bufio.Scanner) (Network, string, error) {
    network := map[string]Adjacencies{}
    parsedSequence, sequence := false, ""

    for scanner.Scan() {
        curLine := scanner.Text()
        if !parsedSequence {
            sequence = curLine
            parsedSequence = true
            continue
        }

        nodeDef := NODE_DEFINITION_REGEX.FindAllString(curLine, -1)
        if len(nodeDef) == 0 {
            continue
        }

        node, left, right := nodeDef[0], nodeDef[1], nodeDef[2]
        network[node] = Adjacencies{ left, right }
    }

    return network, sequence, nil
}

func walkNetwork(network *Network, sequence *string) int {
    steps := 0
    sequenceIdx, sequenceLen := 0, len(*sequence)
    curNode := "AAA"

    for curNode != "ZZZ" {
        adjs := (*network)[curNode]
        dir := (*sequence)[sequenceIdx]
        if dir == 'L' {
            curNode = adjs.Left
        } else {
            curNode = adjs.Right
        }

        sequenceIdx = (sequenceIdx + 1) % sequenceLen
        steps++
    }

    return steps
}

func P8_SolvePartOne(scanner *bufio.Scanner) (string, error) {
    network, sequence, _ := parseNetwork(scanner)
    steps := walkNetwork(&network, &sequence)
    return strconv.Itoa(steps), nil
}

func P8_SolvePartTwo(scanner *bufio.Scanner) (string, error) {
    return "", nil
}

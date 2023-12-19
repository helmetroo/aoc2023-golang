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

func walkNetworkSimul(network *Network, sequence *string) uint64 {
    startNodes := []string{}
    for node := range *network {
        if endWith(&node, 'A') {
            startNodes = append(startNodes, node)
        }
    }

    // At least I figured out I needed to consider the number of steps to Z
    // from each start node and use them together (product seemed reasonable from seeing a pattern in the test input for part 2)
    // Logical leap to take their LCM was thanks to reddit
    steps := []uint64{}
    for _, node := range startNodes {
        steps = append(steps, stepsToZ(network, sequence, &node))
    }

    return lcm(&steps)
}

func stepsToZ(network *Network, sequence *string, startNode *string) uint64 {
    steps := uint64(0)
    sequenceIdx, sequenceLen := 0, len(*sequence)

    curNode := *startNode
    for !endWith(&curNode, 'Z') {
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

func lcm(arr *[]uint64) uint64 {
    size := len(*arr)
    curLcm := (*arr)[0]
    idx := 1

    for idx < size {
        a, b := curLcm, (*arr)[idx]
        curLcm = (a * b) / gcd(a, b)
        idx++
    }

    return curLcm
}

func gcd(a, b uint64) uint64 {
    if b == 0 {
        return a
    }

    return gcd(b, a % b)
}

func endWith(name *string, suffix byte) bool {
    return (*name)[2] == suffix
}

func P8_SolvePartOne(scanner *bufio.Scanner) (string, error) {
    network, sequence, _ := parseNetwork(scanner)
    steps := walkNetwork(&network, &sequence)
    return strconv.Itoa(steps), nil
}

func P8_SolvePartTwo(scanner *bufio.Scanner) (string, error) {
    network, sequence, _ := parseNetwork(scanner)
    steps := walkNetworkSimul(&network, &sequence)
    return strconv.FormatUint(steps, 10), nil
}

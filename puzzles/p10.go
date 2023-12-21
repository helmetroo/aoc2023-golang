package puzzles

import (
	"bufio"
	"strconv"

	"samjay/aoc2023-golang/utils"
)

type Position [2]int
func (dir *Position) add(row, col int) (int, int) {
    thisRow, thisCol := (*dir)[0], (*dir)[1]
    return thisRow + row, thisCol + col
}

func (dir *Position) canEdgeTo(chr byte) bool {
    // Horizontal
    if (*dir)[0] != 0 {
        if chr == '-' {
            return true
        }

        if (*dir)[0] == 1 {
            if chr == 'J' || chr == '7' {
                return true
            }
        } else {
            if chr == 'L' || chr == 'F' {
                return true
            }
        }
    }

    // Vertical
    if chr == '|' {
        return true
    }

    if (*dir)[1] == 1 {
        if chr == 'L' || chr == 'J' {
            return true
        }
    } else {
        if chr == '7' || chr == 'F' {
            return true
        }
    }

    return false
}

var NORTH = Position{ -1, 0 }
var SOUTH = Position{ 1, 0 }
var WEST = Position{ 0, -1 }
var EAST = Position{ 0, 1 }
var ALL_DIRECTIONS = [4]Position { NORTH, SOUTH, WEST, EAST }

var PIPE_DIRECTIONS = map[byte][2]Position {
    '|': { NORTH, SOUTH },
        '-': { WEST, EAST },
        'L': { NORTH, EAST },
        'J': { NORTH, WEST },
        '7': { SOUTH, WEST },
        'F': { SOUTH, EAST },
}

type PipeMap map[MapNode][]MapNode
type MapNode struct {
    Row int
    Col int
    Chr byte
}

func parsePipeMap(scanner *bufio.Scanner) (PipeMap, Position) {
    pipeMap := PipeMap{}

    // Buffer
    colsKnown := false
    rows, cols := 0, 0
    lines := []string{}
    for scanner.Scan() {
        curLine := scanner.Text()

        if !colsKnown {
            cols = len(curLine)
            colsKnown = true
        }

        lines = append(lines, curLine)
        rows++
    }

    startNode, foundStart := Position { 0, 0 }, false
    for row := 0; row < rows; row++ {
        for col := 0; col < cols; col++ {
            chr := lines[row][col]
            if chr == '.' {
                continue
            }

            curNode := MapNode { row, col, chr }
            edges := []MapNode{}

            // Add all possible directions from the start
            if !foundStart && chr == 'S' {
                startNode[0], startNode[1] = row, col
                foundStart = true

                for _, dir := range ALL_DIRECTIONS {
                    adjRow, adjCol := dir.add(row, col)
                    if !utils.ValidOffset(adjRow, adjCol, rows, cols) {
                        continue
                    }

                    adjChr := lines[adjRow][adjCol]
                    if dir.canEdgeTo(adjChr) {
                        edges = append(edges, MapNode { adjRow, adjCol, adjChr })
                    }
                }

                pipeMap[curNode] = edges
                continue
            }

            // Add possible directions from a pipe
            allowedDirs := PIPE_DIRECTIONS[chr]
            for _, dir := range allowedDirs {
                adjRow, adjCol := dir.add(row, col)
                if !utils.ValidOffset(adjRow, adjCol, rows, cols) {
                    continue
                }

                adjChr := lines[adjRow][adjCol]
                edges = append(edges, MapNode { adjRow, adjCol, adjChr })
            }
            pipeMap[curNode] = edges
        }
    }

    return pipeMap, startNode
}

type Visited map[MapNode]int
func farthestDepthFromStart(pipeMap *PipeMap, start *Position) int {
    startNode := MapNode { start[0], start[1], 'S' }
    startEdges := (*pipeMap)[startNode]
    lenStartEdges := len(startEdges)

    // Head out from the start node in only a specific direction
    visiteds := make([]Visited, lenStartEdges)
    for idx := 0; idx < lenStartEdges; idx++ {
        visiteds[idx] = Visited{}
        traversePipeMapFrom(pipeMap, startNode, &startEdges[idx], &visiteds[idx], 0)
    }

    // The intersection of these journeys is where we find the farthest point!
    largestDepth := 0
    for idx := 0; idx < lenStartEdges - 1; idx++ {
        for node, depth := range visiteds[idx] {
            if node == startNode {
                continue
            }

            if otherDepth, exists := visiteds[idx + 1][node]; exists {
                if otherDepth == depth && depth > largestDepth {
                    largestDepth = depth
                }
            }
        }
    }

    return largestDepth
}

func traversePipeMapFrom(pipeMap *PipeMap, from MapNode, fromOnlyEdge *MapNode, visited *Visited, depth int) {
    (*visited)[from] = depth

    // Only visit the node's provided edge (if given)
    if fromOnlyEdge != nil {
        _, visitedNode := (*visited)[*fromOnlyEdge]
        if !visitedNode {
            traversePipeMapFrom(pipeMap, *fromOnlyEdge, nil, visited, depth + 1)
        }
    }

    // Otherwise visit ALL node's edges
    edges := (*pipeMap)[from]
    for _, edge := range edges {
        _, visitedNode := (*visited)[edge]
        if !visitedNode {
            traversePipeMapFrom(pipeMap, edge, nil, visited, depth + 1)
        }
    }
}

func P10_SolvePartOne(scanner *bufio.Scanner) (string, error) {
    pipeMap, start := parsePipeMap(scanner)
    depth := farthestDepthFromStart(&pipeMap, &start)

    return strconv.Itoa(depth), nil
}

func P10_SolvePartTwo(scanner *bufio.Scanner) (string, error) {
    return "", nil
}

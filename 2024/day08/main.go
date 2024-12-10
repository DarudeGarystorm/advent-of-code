package main

import (
	"fmt"

	"advent-of-code/utils"
)

func main() {
	data := utils.GetInputData()
	intData := utils.ParseStringAsRuneArray(data)

	fmt.Println("Part 1:", part1(intData, false))
	fmt.Println("Part 2:", part1(intData, true))
}

type Coordinate struct {
	x, y int
}

func antiNodeLocations(a, b Coordinate, xMax, yMax int, part2 bool) []Coordinate {
	isNodeValid := func(node Coordinate) bool {
		return node.x >= 0 && node.x < xMax && node.y >= 0 && node.y < yMax
	}

	xDiff, yDiff := a.x-b.x, a.y-b.y

	stillValid := true
	validCoordinates := []Coordinate{}

	for i := 0; stillValid; i++ {
		if !part2 {
			i++ // only do 1 outer for part 1
		}

		lenBefore := len(validCoordinates)
		node1, node2 := Coordinate{a.x + i*xDiff, a.y + i*yDiff}, Coordinate{b.x - i*xDiff, b.y - i*yDiff}
		if isNodeValid(node1) {
			validCoordinates = append(validCoordinates, node1)
		}
		if isNodeValid(node2) {
			validCoordinates = append(validCoordinates, node2)
		}
		// exit when we aren't adding more co-ordinates
		// OR exit after 1 time for part 1
		if lenBefore == len(validCoordinates) || !part2 {
			stillValid = false
		}
	}

	return validCoordinates
}

func part1(data [][]rune, actuallyPart2 bool) int {
	antennas := make(map[rune][]Coordinate)

	for i, line := range data {
		for j, char := range line {
			if char != '.' {
				antennas[char] = append(antennas[char], Coordinate{i, j})
			}
		}
	}

	antiNodes := []Coordinate{}
	for _, coordinates := range antennas {
		for i := 0; i < len(coordinates)-1; i++ {
			for j := i + 1; j < len(coordinates); j++ {
				antiNodes = append(antiNodes, antiNodeLocations(coordinates[i], coordinates[j], len(data), len(data[0]), actuallyPart2)...)
			}
		}
	}

	uniqueAntiNodes := make(map[Coordinate]struct{})
	for _, node := range antiNodes {
		uniqueAntiNodes[node] = struct{}{}
	}

	return len(uniqueAntiNodes)
}

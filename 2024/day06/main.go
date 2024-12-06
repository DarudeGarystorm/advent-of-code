package main

import (
	"fmt"
	"slices"
	"strings"

	"advent-of-code/utils"
)

func main() {
	data := utils.GetInputData()
	runes := utils.ParseStringAsRuneArray(data)

	fmt.Println("Part 1:", part1(runes, Coordinate{-2, -2}))
	fmt.Println("Part 2:", part2(runes))
}

type Coordinate struct {
	x, y int
}

// x and y aren't how I imagined it but whatever
var (
	Up    = Coordinate{-1, 0}
	Right = Coordinate{0, 1}
	Down  = Coordinate{1, 0}
	Left  = Coordinate{0, -1}
)

type CoordinateWithDirection struct {
	coordinate Coordinate
	direction  Coordinate
}

func getGuardCoordinates(data [][]rune) Coordinate {
	y := 0
	x := slices.IndexFunc(data, func(line []rune) bool {
		idx := strings.IndexRune(string(line), '^')
		if idx != -1 {
			y = idx
			return true
		}
		return false
	})

	return Coordinate{x, y}
}

func part1(data [][]rune, newBlocker Coordinate) int {
	coord := getGuardCoordinates(data)
	x, y := coord.x, coord.y

	// not the most space efficient but screw it it's easy to count
	visitedPositions := map[Coordinate]bool{}
	visitedPivots := map[CoordinateWithDirection]bool{}

	direction := Up

	// TODO - make a generic in utils
	isCoordinateValid := func(data [][]rune, x int, y int) bool {
		return x >= 0 && y >= 0 && x < len(data) && y < len(data[x])
	}

	for {
		// data[x][y] = 'X' // for printing
		visitedPositions[Coordinate{x, y}] = true

		// Just prints all the values.
		// for _, row := range data {
		// 	for _, colValue := range row {
		// 		fmt.Printf("%c", colValue)
		// 	}
		// 	fmt.Printf("\n")
		// }

		nextX, nextY := x+direction.x, y+direction.y
		nextCoordinate := Coordinate{nextX, nextY}

		// sucks that I check isCoordinateValid twice, but whatever
		if !isCoordinateValid(data, nextX, nextY) {
			return len(visitedPositions)
		}

		// PIVOT!
		if data[nextX][nextY] == '#' || nextCoordinate == newBlocker {
			pivot := CoordinateWithDirection{nextCoordinate, direction}
			if visitedPivots[pivot] {
				return -1
			}
			visitedPivots[pivot] = true
			if direction == Up {
				direction = Right
			} else if direction == Right {
				direction = Down
			} else if direction == Down {
				direction = Left
			} else if direction == Left {
				direction = Up
			}
		} else {
			x, y = nextX, nextY
		}
	}
}

func part2(data [][]rune) int {
	count := 0
	for i, line := range data {
		for j, char := range line {
			if char == '.' {
				if part1(data, Coordinate{i, j}) == -1 {
					count++
				}
			}
		}
	}

	return count
}

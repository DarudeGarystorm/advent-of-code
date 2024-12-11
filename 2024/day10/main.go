package main

import (
	"fmt"
	"strings"

	"advent-of-code/utils"
)

func main() {
	data := utils.GetInputData()

	fmt.Println(strings.Join(data, "\n"))

	grid := utils.ParseInputAsNumArray(data, "")

	fmt.Println("Part 1:", part1(grid, false))
	fmt.Println("Part 2:", part1(grid, true))
}

type Coordinate struct {
	x, y int
}

func countNextStep(currentStep int, grid [][]int, x, y int) (map[Coordinate]bool, int) {
	if currentStep == 9 {
		return map[Coordinate]bool{{x, y}: true}, 1
	}
	reachable, total := map[Coordinate]bool{}, 0
	nextStep := currentStep + 1

	// Up
	if x != 0 && grid[x-1][y] == nextStep {
		reached, score := countNextStep(nextStep, grid, x-1, y)
		total += score
		for k, v := range reached {
			reachable[k] = v
		}
	}
	// Down
	if x != len(grid)-1 && grid[x+1][y] == nextStep {
		reached, score := countNextStep(nextStep, grid, x+1, y)
		total += score
		for k, v := range reached {
			reachable[k] = v
		}
	}
	// Left
	if y != 0 && grid[x][y-1] == nextStep {
		reached, score := countNextStep(nextStep, grid, x, y-1)
		total += score
		for k, v := range reached {
			reachable[k] = v
		}
	}
	// Right
	if y != len(grid[0])-1 && grid[x][y+1] == nextStep {
		reached, score := countNextStep(nextStep, grid, x, y+1)
		total += score
		for k, v := range reached {
			reachable[k] = v
		}
	}

	return reachable, total
}

func part1(grid [][]int, part2 bool) int {
	total := 0
	for i, row := range grid {
		for j, val := range row {
			if val == 0 {
				reachable, score2 := countNextStep(val, grid, i, j)
				if !part2 {
					total += len(reachable)
				} else {
					total += score2
				}
			}
		}
	}

	return total
}

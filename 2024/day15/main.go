package main

import (
	"fmt"
	"slices"

	"advent-of-code/utils"
)

func main() {
	data := utils.GetInputData()

	doneGrid := false
	grid1 := [][]rune{}
	grid2 := [][]rune{}
	directions := []rune{}

	for i, line := range data {
		if len(line) < 1 {
			doneGrid = true
			continue
		}

		if !doneGrid {
			grid1 = append(grid1, []rune{})
			grid2 = append(grid2, []rune{})
		}

		for _, r := range line {
			if doneGrid {
				directions = append(directions, r)
			} else {
				grid1[i] = append(grid1[i], r)
				if r == '#' {
					grid2[i] = append(grid2[i], '#', '#')
				} else if r == 'O' {
					grid2[i] = append(grid2[i], '[', ']')
				} else if r == '.' {
					grid2[i] = append(grid2[i], '.', '.')
				} else if r == '@' {
					grid2[i] = append(grid2[i], '@', '.')
				}
			}
		}
	}

	fmt.Println("Part 1:", part1(grid1, directions))
	fmt.Println("Part 2:", part2(grid2, directions))
}

func canMoveDirection2(grid [][]rune, r, c, dR, dC int) bool {
	nextR, nextC := r+dR, c+dC
	s := grid[nextR][nextC]

	switch s {
	case '.':
		return true
	case '#':
		return false
	case '[':
		return canMoveDirection2(grid, nextR, nextC, dR, dC) && canMoveDirection2(grid, nextR, nextC+1, dR, dC)
	case ']':
		return canMoveDirection2(grid, nextR, nextC-1, dR, dC) && canMoveDirection2(grid, nextR, nextC, dR, dC)
	default:
		panic("this cannot happen")
	}
}

func actuallyMoveDirection2(grid [][]rune, r, c, dR, dC int) {
	nextR, nextC := r+dR, c+dC
	s := grid[nextR][nextC]

	thisOneForReal := func() {
		grid[nextR][nextC], grid[r][c] = grid[r][c], '.'
	}

	switch s {
	case '.':
		thisOneForReal()
		return
	case '[':
		actuallyMoveDirection2(grid, nextR, nextC, dR, dC)
		actuallyMoveDirection2(grid, nextR, nextC+1, dR, dC)
		thisOneForReal()
	case ']':
		actuallyMoveDirection2(grid, nextR, nextC-1, dR, dC)
		actuallyMoveDirection2(grid, nextR, nextC, dR, dC)
		thisOneForReal()
	default:
		panic("how did i get here")
	}
}

func moveDirection2(grid [][]rune, r, c, dR, dC int) (newR, newC int) {
	// going sideways has not changed, it's just up/down that is different
	if dR == 0 {
		return moveDirection(grid, r, c, dR, dC)
	}

	if canMoveDirection2(grid, r, c, dR, dC) {
		actuallyMoveDirection2(grid, r, c, dR, dC)
		return r + dR, c + dC
	}
	return r, c
}

func moveDirection(grid [][]rune, r, c, dR, dC int) (newR, newC int) {
	for i := 1; ; i++ {
		s := grid[r+dR*i][c+dC*i]
		if s == '.' {
			break
		} else if s == '#' {
			return r, c
		}
	}

	lastRune := '.'
	for i := 0; ; i++ {
		nextR, nextC := r+dR*i, c+dC*i
		grid[nextR][nextC], lastRune = lastRune, grid[nextR][nextC]

		if lastRune == '.' {
			return r + dR, c + dC
		}
	}
}

func findGuy(grid [][]rune) (int, int) {
	guyColumn := -1
	guyRow := slices.IndexFunc(grid, func(r []rune) bool {
		guyColumn = slices.Index(r, '@')
		return guyColumn > -1
	})
	return guyRow, guyColumn
}

func part1(grid [][]rune, directions []rune) int {
	guyRow, guyColumn := findGuy(grid)

	for _, d := range directions {
		switch d {
		case '<': // LEFT
			guyRow, guyColumn = moveDirection(grid, guyRow, guyColumn, 0, -1)
		case '>': // RIGHT
			guyRow, guyColumn = moveDirection(grid, guyRow, guyColumn, 0, 1)
		case '^': // UP
			guyRow, guyColumn = moveDirection(grid, guyRow, guyColumn, -1, 0)
		case 'v': // DOWN
			guyRow, guyColumn = moveDirection(grid, guyRow, guyColumn, 1, 0)
		}
		// printGrid(grid)
	}

	return calculateTotal(grid, 'O')
}

func calculateTotal(grid [][]rune, char rune) int {
	total := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == char {
				total += 100*r + c
			}
		}
	}
	return total
}

func printGrid(grid [][]rune) {
	fmt.Println()
	for _, row := range grid {
		for _, r := range row {
			fmt.Print(string(r))
		}
		fmt.Println()
	}
}

func part2(grid [][]rune, directions []rune) int {
	guyRow, guyColumn := findGuy(grid)

	for _, d := range directions {
		switch d {
		case '<': // LEFT
			guyRow, guyColumn = moveDirection2(grid, guyRow, guyColumn, 0, -1)
		case '>': // RIGHT
			guyRow, guyColumn = moveDirection2(grid, guyRow, guyColumn, 0, 1)
		case '^': // UP
			guyRow, guyColumn = moveDirection2(grid, guyRow, guyColumn, -1, 0)
		case 'v': // DOWN
			guyRow, guyColumn = moveDirection2(grid, guyRow, guyColumn, 1, 0)
		}
		// printGrid(grid)
	}

	return calculateTotal(grid, '[')
}

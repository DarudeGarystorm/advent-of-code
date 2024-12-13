package main

import (
	"fmt"

	"advent-of-code/utils"
)

func main() {
	data := utils.GetInputData()
	chars := utils.ParseStringAsRuneArray(data)

	fmt.Println("Part 1:", part1(chars, false))
	fmt.Println("Part 2:", part1(chars, true))
}

type Coordinate struct {
	r, c int
}

func part1(data [][]rune, isPart2 bool) int {
	total := 0
	counted := map[Coordinate]bool{}

	var getGroup func(r, c int) []Coordinate
	getGroup = func(r, c int) []Coordinate {
		counted[Coordinate{r, c}] = true
		restOfGroup := []Coordinate{{r, c}}
		if r+1 < len(data) && !counted[Coordinate{r + 1, c}] && data[r+1][c] == data[r][c] {
			counted[Coordinate{r + 1, c}] = true
			restOfGroup = append(restOfGroup, getGroup(r+1, c)...)
		}

		if r-1 >= 0 && !counted[Coordinate{r - 1, c}] && data[r-1][c] == data[r][c] {
			counted[Coordinate{r - 1, c}] = true
			restOfGroup = append(restOfGroup, getGroup(r-1, c)...)
		}

		if c+1 < len(data[0]) && !counted[Coordinate{r, c + 1}] && data[r][c+1] == data[r][c] {
			counted[Coordinate{r, c + 1}] = true
			restOfGroup = append(restOfGroup, getGroup(r, c+1)...)
		}

		if c-1 >= 0 && !counted[Coordinate{r, c - 1}] && data[r][c-1] == data[r][c] {
			counted[Coordinate{r, c - 1}] = true
			restOfGroup = append(restOfGroup, getGroup(r, c-1)...)
		}

		return restOfGroup
	}

	calculateP := func(coordinate Coordinate) int {
		r, c := coordinate.r, coordinate.c
		p := 0
		if r == 0 || data[r-1][c] != data[r][c] {
			p += 1
		}
		if r == len(data)-1 || data[r+1][c] != data[r][c] {
			p += 1
		}
		if c == 0 || data[r][c-1] != data[r][c] {
			p += 1
		}
		if c == len(data[0])-1 || data[r][c+1] != data[r][c] {
			p += 1
		}
		return p
	}

	countCorners := func(coordinate Coordinate) int {
		corners, r, c := 0, coordinate.r, coordinate.c

		// Outer corners
		if (r == 0 || data[r-1][c] != data[r][c]) && (c == 0 || data[r][c-1] != data[r][c]) {
			corners++
		}
		if (r == 0 || data[r-1][c] != data[r][c]) && (c == len(data[0])-1 || data[r][c+1] != data[r][c]) {
			corners++
		}
		if (r == len(data)-1 || data[r+1][c] != data[r][c]) && (c == 0 || data[r][c-1] != data[r][c]) {
			corners++
		}
		if (r == len(data)-1 || data[r+1][c] != data[r][c]) && (c == len(data[0])-1 || data[r][c+1] != data[r][c]) {
			corners++
		}

		// Inner corners
		if r > 0 && c > 0 && data[r-1][c] == data[r][c] && data[r][c-1] == data[r][c] && data[r-1][c-1] != data[r][c] {
			corners++
		}
		if r > 0 && c < len(data[0])-1 && data[r-1][c] == data[r][c] && data[r][c+1] == data[r][c] && data[r-1][c+1] != data[r][c] {
			corners++
		}
		if r < len(data)-1 && c > 0 && data[r+1][c] == data[r][c] && data[r][c-1] == data[r][c] && data[r+1][c-1] != data[r][c] {
			corners++
		}
		if r < len(data)-1 && c < len(data[0])-1 && data[r+1][c] == data[r][c] && data[r][c+1] == data[r][c] && data[r+1][c+1] != data[r][c] {
			corners++
		}

		return corners
	}

	calculateGroup := func(r, c int) int {
		group := getGroup(r, c)
		area := len(group)
		p := 0

		if isPart2 {
			for _, c := range group {
				p += countCorners(c)
			}
		} else {
			for _, c := range group {
				p += calculateP(c)
			}
		}

		fmt.Printf("Area %c: %d x %d = %d\n", data[r][c], area, p, area*p)

		return area * p
	}

	for r, row := range data {
		for c, _ := range row {
			if !counted[Coordinate{r, c}] {
				total += calculateGroup(r, c)
			}
		}
	}

	fmt.Println("Counted:", len(counted), "/", len(data)*len(data[0]))

	return total
}

func part2(data [][]rune) int {
	// TODO: Implement part 2
	return 0
}

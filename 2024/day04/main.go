package main

import (
	"fmt"
	"strings"

	"advent-of-code/utils"
)

func main() {
	data := utils.GetInputData()
	runeData := utils.ParseStringAsRuneArray(data)

	fmt.Println("Part 1:", part1(runeData))
	fmt.Println("Part 2:", part2(runeData))
}

func countWordInAllDirections(data [][]rune, x int, y int, s string) int {
	searchString := []rune(s)
	// early exit because it make sense
	if data[x][y] != searchString[0] {
		return 0
	}

	count := 0
	// →
	for i, c := range searchString {
		if x+i >= len(data) || data[x+i][y] != c {
			break
		} else if i == len(searchString)-1 {
			count++
		}
	}
	// ←
	for i, c := range searchString {
		if x-i < 0 || data[x-i][y] != c {
			break
		} else if i == len(searchString)-1 {
			count++
		}
	}
	// ↑
	for i, c := range searchString {
		if y+i >= len(data[x]) || data[x][y+i] != c {
			break
		} else if i == len(searchString)-1 {
			count++
		}
	}
	// ↓
	for i, c := range searchString {
		if y-i < 0 || data[x][y-i] != c {
			break
		} else if i == len(searchString)-1 {
			count++
		}
	}
	// ↖
	for i, c := range searchString {
		if x-i < 0 || y+i >= len(data[x]) || data[x-i][y+i] != c {
			break
		} else if i == len(searchString)-1 {
			count++
		}
	}
	// ↗
	for i, c := range searchString {
		if x+i >= len(data) || y+i >= len(data[x]) || data[x+i][y+i] != c {
			break
		} else if i == len(searchString)-1 {
			count++
		}
	}
	// ↘
	for i, c := range searchString {
		if x+i >= len(data) || y-i < 0 || data[x+i][y-i] != c {
			break
		} else if i == len(searchString)-1 {
			count++
		}
	}
	// ↙
	for i, c := range searchString {
		if x-i < 0 || y-i < 0 || data[x-i][y-i] != c {
			break
		} else if i == len(searchString)-1 {
			count++
		}
	}

	return count
}

func part1(data [][]rune) int {
	count := 0

	for i, row := range data {
		// fmt.Println()
		for j := range row {
			// fmt.Printf("%c", data[i][j])
			count += countWordInAllDirections(data, i, j, "XMAS")
		}
	}

	return count
}

func part2(data [][]rune) int {
	// plan is to look for A within the edges
	// when we see an A, check for M&S on the diagonals

	count := 0
	for i := 1; i < len(data)-1; i++ {
		for j := 1; j < len(data[i])-1; j++ {
			// there must be an A in the middle for it to be an X-MAS
			if data[i][j] != 'A' {
				continue
			}

			searchString := "MS"
			// check for M or S then make sure they are different
			if strings.ContainsRune(searchString, data[i-1][j-1]) &&
				strings.ContainsRune(searchString, data[i+1][j+1]) &&
				data[i-1][j-1] != data[i+1][j+1] {
				// half the X!
				if strings.ContainsRune(searchString, data[i-1][j+1]) &&
					strings.ContainsRune(searchString, data[i+1][j-1]) &&
					data[i-1][j+1] != data[i+1][j-1] {
					// the other half!
					count++
				}
			}
		}
	}

	return count
}

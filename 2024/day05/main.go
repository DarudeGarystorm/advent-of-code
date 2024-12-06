package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"advent-of-code/utils"
)

func main() {
	data := utils.GetInputData()

	divider := 0
	// we could use string and convert to number later?
	mustBeAfter := make(map[int][]int)
	updates := [][]int{}

	for idx, line := range data {
		if len(line) == 0 {
			divider = idx
		} else if divider == 0 {
			// why am i flip flopping on this
			rule := strings.Split(line, "|")
			page1, _ := strconv.Atoi(rule[1])
			page2, _ := strconv.Atoi(rule[0])
			mustBeAfter[page1] = append(mustBeAfter[page1], page2)
		} else if divider != 0 {
			pages := strings.Split(line, ",")
			updates = append(updates, []int{})
			for _, p := range pages {
				pageNum, _ := strconv.Atoi(p)
				updates[len(updates)-1] = append(updates[len(updates)-1], pageNum)
			}
		}
	}

	p1Ans, p2Updates := part1(mustBeAfter, updates)

	fmt.Println("Part 1:", p1Ans)
	fmt.Println("Part 2:", part2(mustBeAfter, p2Updates))
}

func part1(mustBeAfter map[int][]int, updates [][]int) (int, [][]int) {
	total := 0
	wrongLines := [][]int{}

	for _, line := range updates {
		isPlacedAfter := make(map[int]bool)

	out:
		for i, update := range slices.Backward(line) {
			updatesToCheck := mustBeAfter[update]
			for _, update := range updatesToCheck {
				if isPlacedAfter[update] {
					wrongLines = append(wrongLines, line)
					break out
				}
			}

			if i == 0 {
				total += line[len(line)/2]
			}

			isPlacedAfter[update] = true
		}
	}

	return total, wrongLines
}

func part2(mustBeAfter map[int][]int, updates [][]int) int {
	total := 0

	for _, line := range updates {
		for i := 0; i < len(line); i++ {
			numsToBeAfter := mustBeAfter[line[i]]
			for j := len(line) - 1; j >= i; j-- {
				if slices.Contains(numsToBeAfter, line[j]) {
					line = utils.MoveElement(line, i, j)
					i--
					break
				}
			}
		}

		total += line[len(line)/2]
	}

	return total
}

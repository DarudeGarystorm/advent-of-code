package main

import (
	"fmt"
	"strconv"

	"advent-of-code/utils"
)

func main() {
	data := utils.GetInputData()
	intData := utils.ParseStringAsNumArray(data)

	// Create a copy of intData[0]
	p1, p2 := make([]int, len(intData[0])), make([]int, len(intData[0]))
	copy(p1, intData[0])
	copy(p2, intData[0])

	fmt.Println("Part 1:", part2(p1, 25))
	fmt.Println("Part 2:", part2(p2, 75))
}

// second stone will be -1 if not applicatble
func blinkStone(stone int) (int, int) {
	if stone == 0 {
		return 1, -1
	}

	s := strconv.Itoa(stone)
	if len(s)%2 == 0 {
		mid := len(s) / 2
		left, _ := strconv.Atoi(s[:mid])
		right, _ := strconv.Atoi(s[mid:])
		return left, right
	}

	return stone * 2024, -1
}

func part2(input []int, blinks int) int {
	// map of stone's value -> count
	var curStones map[int]int
	nextStones := make(map[int]int, len(input))
	for _, v := range input {
		nextStones[v] = 1
	}

	for i := 0; i < blinks; i++ {
		curStones, nextStones = nextStones, map[int]int{}

		for stone, count := range curStones {
			s1, s2 := blinkStone(stone)
			nextStones[s1] = nextStones[s1] + count
			if s2 != -1 {
				nextStones[s2] = nextStones[s2] + count
			}
		}
	}

	total := 0
	for _, count := range nextStones {
		total += count
	}

	return total
}

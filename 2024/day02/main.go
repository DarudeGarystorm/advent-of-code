package main

import (
	"fmt"
	"math"

	"advent-of-code/utils"
)

func main() {
	data := utils.GetInputData()
	intData := utils.ParseStringAsNumArray(data)

	fmt.Println("Part 1:", part1(intData))
	fmt.Println("Part 2:", part2(intData))
}

func checkRow(row []int) (bool, int) {
	asc := row[0] < row[len(row)-1]

	for i := 0; i < len(row)-1; i++ {
		step := row[i+1] - row[i]

		if math.Abs(float64(step)) > 3 {
			return false, i
		} else if asc && step < 1 {
			return false, i
		} else if !asc && step > -1 {
			return false, i
		}
	}

	return true, -1
}

func part1(data [][]int) int {
	safe := 0
	for _, row := range data {
		if ok, _ := checkRow(row); ok {
			safe++
		}
	}

	return safe
}

func removeElement(slice []int, index int) []int {
	newSlice := make([]int, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)
	return newSlice
}

func part2(data [][]int) int {

	safe := 0
	for _, row := range data {
		// brute force bad
		for i := range row {
			ok, _ := checkRow(removeElement(row, i))

			if ok {
				safe++
				break
			}
		}
	}

	return safe
}

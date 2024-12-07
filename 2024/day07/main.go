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

	equation := [][]int{}

	for _, line := range data {
		split := strings.Split(line, ": ")
		answer, _ := strconv.Atoi(split[0])
		stringValues := strings.Fields(split[1])
		numValues := []int{}
		for _, v := range stringValues {
			numValue, _ := strconv.Atoi(v)
			numValues = append(numValues, numValue)
		}
		equation = append(equation, append([]int{answer}, numValues...))
	}

	fmt.Println("Part 1:", part1(equation))
	fmt.Println("Part 2:", part2(equation))
}

func part1(equations [][]int) int {
	total := 0

	for _, equation := range equations {
		answer := equation[0]
		currentValues := []int{}
		nextValues := []int{answer}

		for i := len(equation) - 1; i > 0; i-- {
			currentValues, nextValues = nextValues, currentValues[:0]
			currentValue := equation[i]

			for _, value := range currentValues {
				// this is impossible to reach
				// working left to right
				if value == 0 {
					continue
				}
				subtractAnswer := value - currentValue
				if subtractAnswer >= 0 {
					nextValues = append(nextValues, subtractAnswer)
				}
				if value%currentValue == 0 {
					divideAnswer := value / currentValue
					nextValues = append(nextValues, divideAnswer)
				}
			}
		}

		if slices.Contains(nextValues, 0) {
			total += answer
		}
	}

	return total
}

// forgive me DRY gods
func part2(equations [][]int) int {
	total := 0

	for _, equation := range equations {
		answer := equation[0]
		currentValues := []int{}
		nextValues := []int{answer}

		for i := len(equation) - 1; i > 0; i-- {
			currentValues, nextValues = nextValues, currentValues[:0]
			currentValue := equation[i]

			for _, value := range currentValues {
				// this is impossible to reach
				// working left to right
				if value == 0 {
					continue
				}
				subtractAnswer := value - currentValue
				if subtractAnswer >= 0 {
					nextValues = append(nextValues, subtractAnswer)
				}
				if value%currentValue == 0 {
					divideAnswer := value / currentValue
					nextValues = append(nextValues, divideAnswer)
				}
				valueString, currentString := strconv.Itoa(value), strconv.Itoa((currentValue))
				before, found := strings.CutSuffix(valueString, currentString)
				if found {
					beforeAsNumber, _ := strconv.Atoi(before)
					nextValues = append(nextValues, beforeAsNumber)
				}
			}
		}

		if slices.Contains(nextValues, 0) {
			total += answer
		}
	}

	return total
}

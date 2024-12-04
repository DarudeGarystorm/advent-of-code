package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"advent-of-code/utils"
)

func main() {
	data := utils.GetInputData()

	// I was looping over each row and "forgetting"
	// what the flag was set to... this is a quick
	// way to remove that loop, newline is just a
	// character right :)
	fixedData := strings.Join(data, "\n")

	fmt.Println("Part 1:", part1(fixedData))
	fmt.Println("Part 2:", part2(fixedData))
}

func part1(line string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	total := 0

	matches := re.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		total += num1 * num2
	}

	return total
}

func part2(line string) int {
	re := regexp.MustCompile(`(mul\((\d+),(\d+)\)|do\(\)|don't\(\))`)
	total := 0

	matches := re.FindAllStringSubmatch(line, -1)
	do := true
	for _, match := range matches {
		if match[0] == "don't()" {
			do = false
		} else if match[0] == "do()" {
			do = true
		} else if do {
			num1, _ := strconv.Atoi(match[2])
			num2, _ := strconv.Atoi(match[3])
			total += num1 * num2
		}
	}

	return total
}

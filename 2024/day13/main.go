package main

import (
	"fmt"
	"regexp"
	"strconv"

	"advent-of-code/utils"
)

func main() {
	data := utils.GetInputData()
	input := []puzzle{}

	for i := 0; i < len(data); i += 4 {
		x := regexp.MustCompile("X.([0-9]+)")
		y := regexp.MustCompile("Y.([0-9]+)")
		ax, _ := strconv.Atoi(x.FindStringSubmatch(data[i])[1])
		ay, _ := strconv.Atoi(y.FindStringSubmatch(data[i])[1])
		bx, _ := strconv.Atoi(x.FindStringSubmatch(data[i+1])[1])
		by, _ := strconv.Atoi(y.FindStringSubmatch(data[i+1])[1])
		xt, _ := strconv.Atoi(x.FindStringSubmatch(data[i+2])[1])
		yt, _ := strconv.Atoi(y.FindStringSubmatch(data[i+2])[1])
		input = append(input, puzzle{ax, ay, bx, by, xt, yt})
	}

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

type puzzle struct {
	ax, ay, bx, by, xt, yt int
}

type presses struct {
	totalX, totalY, cost int
}

// https://en.wikipedia.org/wiki/Cramer%27s_rule
func solve(p puzzle, aCost, bCost int) (a, b, cost int) {
	d := p.ax*p.by - p.bx*p.ay
	d1 := p.xt*p.by - p.yt*p.bx
	d2 := p.yt*p.ax - p.xt*p.ay

	if d1%d != 0 || d2%d != 0 {
		return 0, 0, 0
	}

	a, b = d1/d, d2/d
	cost = (a * aCost) + (b * bCost)
	return
}

func part1(cases []puzzle) int {
	total := 0
	for _, p := range cases {
		_, _, cost := solve(p, 3, 1)
		total += cost
	}
	return total
}

func part2(cases []puzzle) int {
	total := 0
	for _, p := range cases {
		p.xt += 10000000000000
		p.yt += 10000000000000
		_, _, cost := solve(p, 3, 1)
		total += cost
	}
	return total
}

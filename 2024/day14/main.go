package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"

	"advent-of-code/utils"
)

type Coordinate struct {
	r, c int
}

type Robot struct {
	p, v Coordinate
}

func main() {
	robots := []Robot{}
	data := utils.GetInputData()

	r := regexp.MustCompile(`(-?\d+),(-?\d+)`)

	for _, line := range data {
		matchers := r.FindAllStringSubmatch(line, -1)

		pc, _ := strconv.Atoi(matchers[0][1])
		pr, _ := strconv.Atoi(matchers[0][2])
		vc, _ := strconv.Atoi(matchers[1][1])
		vr, _ := strconv.Atoi(matchers[1][2])

		robots = append(
			robots, Robot{
				p: Coordinate{pr, pc},
				v: Coordinate{vr, vc},
			},
		)
	}

	// fmt.Println("Example:", part1(robots, 100, 11, 7))
	fmt.Println("Part 1:", part1(robots, 100, 101, 103))
	fmt.Println("Part 2:", part2(robots, 101, 103))
}

func calculatePositionAfterX(r Robot, wide, tall, turns int) Coordinate {
	fc, fr := r.p.c+(r.v.c*turns), r.p.r+(r.v.r*turns)
	fc, fr = fc%wide, fr%tall
	if fc < 0 {
		fc = wide + fc
	}
	if fr < 0 {
		fr = tall + fr
	}
	return Coordinate{fr, fc}
}

func part1(robots []Robot, turns, wide, tall int) int {
	f := []Coordinate{}
	for _, r := range robots {
		f = append(f, calculatePositionAfterX(r, wide, tall, turns))
	}

	wideD, tallD := (wide-1)/2, (tall-1)/2
	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, r := range f {
		if r.r < tallD && r.c < wideD {
			q1++
		} else if r.r > tallD && r.c < wideD {
			q2++
		} else if r.r < tallD && r.c > wideD {
			q3++
		} else if r.r > tallD && r.c > wideD {
			q4++
		}
	}

	return (q1 * q2 * q3 * q4)
}

func part2(robots []Robot, wide, tall int) int {
	for s := 7344; s < 7345; s++ { // MY ANSWER IS HERE!!!
		fmt.Println("")
		fmt.Println("s:", s)
		for r := 0; r < tall; r++ {
			fmt.Println("")
			for c := 0; c < wide; c++ {
				if slices.ContainsFunc(robots, func(rob Robot) bool {
					loc := calculatePositionAfterX(rob, wide, tall, s)
					return loc.c == c && loc.r == r
				}) {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}
			}
		}
	}
	return 0
}

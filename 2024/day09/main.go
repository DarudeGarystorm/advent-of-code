package main

import (
	"fmt"

	"advent-of-code/utils"
)

func main() {
	data := utils.GetInputData()

	intData := []int{}
	for _, digit := range data[0] {
		intData = append(intData, int(digit-'0'))
	}

	fmt.Println("Part 1:", part1(intData))
	fmt.Println("Part 2:", part2(intData))
}

func part1(data []int) int {
	total := 0

	// why we declare these outside of for and the others in for
	// idk - because I feel like it?
	multiplier := -1 // start at -1 because we add to it in loop
	rightIndex := len(data) - 1
	rightCountRemaining := data[rightIndex]

	for leftIndex := 0; ; leftIndex++ {
		if leftIndex%2 == 0 { // isEven - take from left
			if leftIndex == rightIndex {
				for ; rightCountRemaining > 0; rightCountRemaining-- {
					multiplier++
					total += leftIndex / 2 * multiplier
				}
				return total
			}

			for i := 1; i <= data[leftIndex]; i++ {
				multiplier++
				total += leftIndex / 2 * multiplier
			}
		} else { // isOdd - take from right
			value := data[leftIndex]
			for i := 1; i <= value; i++ {
				multiplier++
				if rightCountRemaining > 0 {
					total += rightIndex / 2 * multiplier
					rightCountRemaining--
				} else {
					rightIndex -= 2 // skip the "blank" spaces
					rightCountRemaining = data[rightIndex]
					if rightIndex > leftIndex {
						total += rightIndex / 2 * multiplier
						rightCountRemaining--
					} // no else because this "isOdd" loop - left != right
				}
			}
		}
	}
}

type File struct {
	id, length int
}

func part2(data []int) int {
	fileSystem := make([]File, 0)
	// create filesystem
	for i, digit := range data {
		if i%2 == 0 {
			fileSystem = append(fileSystem, File{id: i / 2, length: digit})
		} else {
			fileSystem = append(fileSystem, File{id: -1, length: digit})
		}
	}

	// fill the gaps
	for r := len(fileSystem) - 1; r > 0; r-- {
		if fileSystem[r].id == -1 {
			continue
		}

		for l := 0; l < r; l++ {
			// check if empty file
			if fileSystem[l].id != -1 {
				continue
			}
			// check if fits
			if fileSystem[r].length > fileSystem[l].length {
				continue
			}
			// equal space, straight swap
			if fileSystem[r].length == fileSystem[l].length {
				fileSystem[l].id = fileSystem[r].id
				fileSystem[r].id = -1 // old spot is now empty space
				break
			}

			// lesser space, swap and adjust
			fileSystem[l].length -= fileSystem[r].length
			fileSystem = utils.InsertAtIndex(fileSystem, fileSystem[r], l)
			fileSystem[r+1].id = -1 // old spot is now empty space
			break
		}
	}

	// calculate checksum
	total := 0
	multiplier := -1 // start at -1 because we increment in loop

	for _, file := range fileSystem {
		for j := 0; j < file.length; j++ {
			multiplier++
			if file.id != -1 { // -1 is empty space
				total += file.id * multiplier
			}
		}
	}

	return total
}

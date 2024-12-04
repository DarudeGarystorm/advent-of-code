package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run init_day.go <year> <current_day>")
		return
	}

	year := os.Args[1]
	day, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid day:", os.Args[2])
		return
	}

	// Define the paths
	dayDir := fmt.Sprintf("%s/day%02d", year, day)
	mainFile := "main.go"

	// Create the day directory
	err = os.MkdirAll(dayDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	// Create the main.go file for the day with boilerplate code
	dstFile := filepath.Join(dayDir, mainFile)
	err = createBoilerplateFile(dstFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	fmt.Printf("Day %02d setup complete. Please download the input file and save it as %s/input.txt\n", day, dayDir)
}

// createBoilerplateFile creates a new main.go file with boilerplate code
func createBoilerplateFile(dst string) error {
	boilerplate := `package main

import (
	"fmt"

	"advent-of-code/utils"
)

func main() {
	data := utils.GetInputData()
	intData := utils.ParseStringAsNumArray(data)

	fmt.Println("Part 1:", part1(intData))
	fmt.Println("Part 2:", part2(intData))
}

func part1(data [][]int) int {
	// TODO: Implement part 1
	return 0
}

func part2(data [][]int) int {
	// TODO: Implement part 2
	return 0
}
`
	return os.WriteFile(dst, []byte(boilerplate), 0644)
}

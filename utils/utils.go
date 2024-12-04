package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func GetInputData() []string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("unable to get current file path")
	}

	sourcePath := filepath.Dir(filename)
	inputFilePath := filepath.Join(sourcePath, "input.txt")

	file, err := os.Open(inputFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data []string

	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return data
}

func ParseStringAsRuneArray(data []string) [][]rune {
	var runeData [][]rune
	for _, line := range data {
		runeData = append(runeData, []rune(line))
	}
	return runeData
}

func ParseStringAsNumArray(data []string) [][]int {
	var intData [][]int
	for _, line := range data {
		numbers := strings.Fields(line)
		row := make([]int, len(numbers))
		for i, num := range numbers {
			value, err := strconv.Atoi(num)
			if err != nil {
				panic(fmt.Errorf("error converting strings to numbers"))
			}
			row[i] = value
		}
		intData = append(intData, row)
	}
	return intData
}

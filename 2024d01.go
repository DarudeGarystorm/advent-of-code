package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type SortedList struct {
	list *list.List
}

func NewSortedList() *SortedList {
	return &SortedList{list: list.New()}
}

func (s *SortedList) Insert(value int) {
	for e := s.list.Front(); e != nil; e = e.Next() {
		if e.Value.(int) > value {
			s.list.InsertBefore(value, e)
			return
		}
	}
	s.list.PushBack(value)
}

func main() {
	// Open the file
	file, err := os.Open("2024d01.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new scanner to read the file
	scanner := bufio.NewScanner(file)

	sortedList1 := NewSortedList()
	sortedList2 := NewSortedList()

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()

		numbers := strings.Fields(line)
		num1, err1 := strconv.Atoi(numbers[0])
		num2, err2 := strconv.Atoi(numbers[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Gary you have an issue converting strings to numbers")
			continue
		}
		sortedList1.Insert(num1)
		sortedList2.Insert(num2)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Compare the two lists
	totalDiff := 0
	p1 := sortedList1.list.Front()
	p2 := sortedList2.list.Front()

	for p1 != nil && p2 != nil {
		diff := int(math.Abs(float64(p1.Value.(int) - p2.Value.(int))))
		totalDiff += diff
		p1 = p1.Next()
		p2 = p2.Next()
	}

	fmt.Println("Part 1:", totalDiff)

	// fuck it my glass is empty
	simScore := 0
	for p1 = sortedList1.list.Front(); p1 != nil; p1 = p1.Next() {
		count := 0
		for p2 = sortedList2.list.Front(); p2 != nil; p2 = p2.Next() {
			if p1.Value.(int) == p2.Value.(int) {
				count++
			}
		}
		simScore += p1.Value.(int) * count
	}

	fmt.Println("Part 2:", simScore)
}

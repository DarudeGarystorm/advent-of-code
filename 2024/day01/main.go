package main

import (
	"container/list"
	"fmt"
	"math"

	"advent-of-code/utils"
)

type SortedList struct {
	list *list.List
}

func NewSortedList() *SortedList {
	return &SortedList{list: list.New()}
}

// I know this is inefficient - it's just for learning
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
	data := utils.GetInputData()

	sortedList1 := NewSortedList()
	sortedList2 := NewSortedList()

	for _, row := range data {
		sortedList1.Insert(row[0])
		sortedList2.Insert(row[1])
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

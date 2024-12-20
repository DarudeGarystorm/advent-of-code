package utils_test

import (
	"advent-of-code/utils"
	"os"
	"reflect"
	"testing"
)

func TestGetInputData(t *testing.T) {
	// Create a temporary file with test data
	filename := "input.txt"
	content := "1 2\n3 4\n5 6\n"
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(filename)

	expected := []string{
		"1 2",
		"3 4",
		"5 6",
	}

	data := utils.GetInputData()

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, got %v", expected, data)
	}
}

func TestParseStringAsNumArray(t *testing.T) {
	data := []string{
		"1 2",
		"3 4",
		"5 6",
	}

	expected := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	intData := utils.ParseStringAsNumArray(data)

	if !reflect.DeepEqual(intData, expected) {
		t.Errorf("Expected %v, got %v", expected, intData)
	}
}

func TestInsertAtIndex(t *testing.T) {
	array := []int{1, 2, 3, 4}
	value := 99
	index := 2
	expected := []int{1, 2, 99, 3, 4}

	result := utils.InsertAtIndex(array, value, index)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestRemoveAtIndex(t *testing.T) {
	array := []int{1, 2, 3, 4}
	index := 2
	expected := []int{1, 2, 4}

	result := utils.RemoveAtIndex(array, index)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMoveElement(t *testing.T) {
	array := []int{1, 2, 3, 4}
	srcIndex := 1
	dstIndex := 3
	expected := []int{1, 3, 4, 2}

	result := utils.MoveElement(array, srcIndex, dstIndex)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

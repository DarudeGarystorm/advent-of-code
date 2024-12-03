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

	expected := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	data := utils.GetInputData()

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("Expected %v, got %v", expected, data)
	}
}

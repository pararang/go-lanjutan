package unittest

import (
	"testing"
)

func TestHomeworkBE4908584(t *testing.T) {
	// Case 1 : If empty slice
	numbers := []int{}
	expectedResult := 0.0
	if result := HomeworkBE4908584(numbers); result != expectedResult {
		t.Errorf("Expected %f, but got %f", expectedResult, result)
	}

	// Case 2 : Slice with only numbers divisible by 3
	numbers = []int{3, 6, 9}
	expectedResult = 0.0
	if result := HomeworkBE4908584(numbers); result != expectedResult {
		t.Errorf("Expected %f, but got %f", expectedResult, result)
	}

	// Case 3 : Slice with more than 10 items
	numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	expectedResult = -1
	if result := HomeworkBE4908584(numbers); result != expectedResult {
		t.Errorf("Expected %f, but got %f", expectedResult, result)
	}

	// Case 4 : Slice with some numbers not divisible by 3
	numbers = []int{1, 2, 3, 4, 5, 6}
	result := HomeworkBE4908584(numbers)
	if result == 0.0 {
		t.Errorf("Expected non-zero value, but got 0.0")
	}
}

package unittest

import "testing"

func TestHomeworkBE5187561(t *testing.T) {
	// Test case 1: positive numbers
	input1 := []int{1, 2, 3, 4, 5}
	expectedOutput1 := 15
	output1 := HomeworkBE5187561(input1...)
	if output1 != expectedOutput1 {
		t.Errorf("HomeworkBE5187561(%v) = %d; expected %d", input1, output1, expectedOutput1)
	}

	// Test case 2: negative numbers
	input2 := []int{-1, -2, -3, -4, -5}
	expectedOutput2 := 15
	output2 := HomeworkBE5187561(input2...)
	if output2 != expectedOutput2 {
		t.Errorf("HomeworkBE5187561(%v) = %d; expected %d", input2, output2, expectedOutput2)
	}

	// Test case 3: zero input
	input4 := []int{0, 0, 0}
	expectedOutput4 := 10
	output4 := HomeworkBE5187561(input4...)
	if output4 != expectedOutput4 {
		t.Errorf("HomeworkBE5187561(%v) = %d; expected %d", input4, output4, expectedOutput4)
	}

}
package unittest

import "testing"

func TestHomeworkBE4942576(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{5, 120},
		{10, -1},
		{-1, 0},
	}

	for _, tc := range testCases {
		result := HomeworkBE4942576(tc.input)
		if result != tc.expected {
			t.Errorf("HomeworkBE4942576(%d): expected %d, but got %d", tc.input, tc.expected, result)
		}
	}
}

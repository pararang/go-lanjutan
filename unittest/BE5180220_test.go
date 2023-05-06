package unittest

import (
	"reflect"
	"testing"
)

func TestHomeworkBE5180220(t *testing.T) {
	testCases := []struct {
		input       string
		expectedMap map[string]int
	}{
		{
			input:       "hello",
			expectedMap: map[string]int{"H": 1, "E": 1, "L": 2, "O": 1},
		},
		{
			input:       "Hello World!",
			expectedMap: map[string]int{"H": 1, "E": 1, "L": 3, "O": 2, "W": 1, "R": 1, "D": 1},
		},
		{
			input:       "",
			expectedMap: nil,
		},
		{
			input:       "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
			expectedMap: map[string]int{"abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz": 1},
		},
	}

	for _, tc := range testCases {
		result := HomeworkBE5180220(tc.input)
		if !reflect.DeepEqual(result, tc.expectedMap) {
			t.Errorf("Unexpected result. Expected: %v, but got: %v", tc.expectedMap, result)
		}
	}
}

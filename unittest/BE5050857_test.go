package unittest

import (
	"reflect"
	"testing"
)

func TestHomeworkBE5050857(t *testing.T) {
	testCases := []struct {
		input       string
		expectedMap map[string]int
	}{
		{
			input:       "the quick brown fox jumps over the lazy dog",
			expectedMap: map[string]int{"the": 2, "quick": 1, "brown": 1, "fox": 1, "jumps": 1, "over": 1, "lazy": 1, "dog": 1},
		},
		{
			input:       "this is a test string with more than 5 words in it to check the functionality of the function",
			expectedMap: map[string]int{"this": 1, "is": 1, "a": 1, "test": 1, "string": 1, "with": 1, "more": 1, "than": 1, "5": 1, "words": 1, "in": 1, "it": 1, "to": 1, "check": 1, "the": 1, "functionality": 1, "of": 1, "function": 1},
		},
		{
			input:       "",
			expectedMap: nil,
		},
		{
			input:       "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed euismod eget lectus vitae ullamcorper. Sed vel est tortor. Nunc eu ligula pharetra, ultrices elit vel, dictum tellus. Etiam gravida facilisis quam, ac commodo ex auctor nec. Nullam consequat a ante sit amet ullamcorper. Pellentesque faucibus facilisis dui, eget eleifend purus fermentum id. Aliquam id massa semper, sollicitudin mi eget, eleifend odio. Sed posuere fringilla tellus a luctus. Donec eu ante leo. Vestibulum accumsan ultricies aliquet.",
			expectedMap: map[string]int{"Lorem": 1, "ipsum": 1, "dolor": 1, "sit": 1, "amet,": 1, "consectetur": 1, "adipiscing": 1, "elit.": 1, "Sed": 2, "euismod": 1, "eget": 2, "lectus": 1, "vitae": 1, "ullamcorper.": 2, "vel": 2, "est": 1, "tortor.": 1, "Nunc": 1, "eu": 1, "ligula": 1, "pharetra,": 1, "ultrices": 1, "elit": 1, "dictum": 1, "tellus.": 1, "Etiam": 1, "gravida": 1, "facilisis": 2, "quam,": 1, "ac": 1, "commodo": 1, "ex": 1, "auctor": 1, "nec.": 1, "Nullam": 1, "consequat": 1, "a": 1, "ante": 1, "amet": 1, "ullamcorper": 1, "Pellentes": 1},
		},
	}
	for _, tc := range testCases {
		result := HomeworkBE5050857(tc.input)
		if !reflect.DeepEqual(result, tc.expectedMap) {
			t.Errorf("Unexpected result. Expected: %v, but got: %v", tc.expectedMap, result)
		}
	}
}

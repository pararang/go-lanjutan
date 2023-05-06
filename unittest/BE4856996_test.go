package unittest

import (
	"reflect"
	"testing"
)

func TestHomeworkBE4856996(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "Empty Input",
			args: args{input: ""},
			want: nil,
		},
		{
			name: "Testcase 1 - less than 50 words",
			args: args{input: "Hello, World!"},
			want: map[string]int{
				"H": 1,
				"E": 1,
				"L": 3,
				"O": 2,
				"W": 1,
				"R": 1,
				"D": 1,
			},
		},
		{
			name: "Testcase 2 - more than 50 words",
			args: args{input: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."},
			want: map[string]int{
				"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.": 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HomeworkBE4856996(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HomeworkBE4856996() = %v, want %v", got, tt.want)
			}
		})
	}
}

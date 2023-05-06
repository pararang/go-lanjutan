package unittest

import (
	"reflect"
	"testing"
)

func TestHomeworkBE4855066(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		// TODO: Add test cases.
		{
			name: "should return nil if input is empty",
			args: args{input: ""},
			want: nil,
		},
		{
			name: "should return input as key and 1 as value if length of input greater than 100",
			args: args{input: "00abcdefghijklmnopqrstuvwxyz1234567abcdefghijklmnopqrstuvwxyz1234567abcdefghijklmnopqrstuvwxyz1234567"},
			want: map[string]int{
				"00abcdefghijklmnopqrstuvwxyz1234567abcdefghijklmnopqrstuvwxyz1234567abcdefghijklmnopqrstuvwxyz1234567": 1,
			},
		},
		{
			name: "should return map with maximum 5 value with the same key",
			args: args{input: " a a a a a a  foo foo foo foo foo bar  bar  bar   baz  baz    aa aa aa aa"},
			want: map[string]int{
				"a"		: 5,
				"aa"	: 4,
				"bar"	: 3,
				"baz"	: 2,
				"foo"	: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HomeworkBE4855066(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HomeworkBE4855066() = %v, want %v", got, tt.want)
			}
		})
	}
}

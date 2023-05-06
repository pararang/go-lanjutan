package unittest

import (
	"reflect"
	"testing"
)

func TestHomeworkBE4920086(t *testing.T) {

	tests := []struct {
		name       string
		input      []int
		wantSorted []int
		wantErr    bool
	}{
		{
			name:       "Should return error if input len less than 5",
			input:      []int{1, 3, 4, 9},
			wantSorted: nil,
			wantErr:    true,
		},
		{
			name:       "Should return error if input len greather than 30",
			input:      []int{1, 2, 4, 5, 9, 1, 2, 4, 5, 9, 1, 2, 4, 5, 9, 1, 2, 4, 5, 9, 1, 2, 4, 5, 9, 1, 2, 4, 5, 9, 1, 2, 4, 5, 9},
			wantSorted: nil,
			wantErr:    true,
		},
		{
			name:       "Should sorted slice if input true",
			input:      []int{1, 3, 4, 5, 2},
			wantSorted: []int{1, 2, 3, 4, 5},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSorted, err := HomeworkBE4920086(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("HomeworkBE4920086() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSorted, tt.wantSorted) {
				t.Errorf("HomeworkBE4920086() = %v, want %v", gotSorted, tt.wantSorted)
			}
		})
	}
}

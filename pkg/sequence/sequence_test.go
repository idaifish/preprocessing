// Package sequence provides utilities for preprocessing sequence data.
package sequence

import (
	"reflect"
	"testing"
)

func TestPadIntSequences(t *testing.T) {
	type args struct {
		sequences  [][]int
		maxlen     int
		padding    string
		truncating string
		value      int
	}
	tests := []struct {
		name       string
		args       args
		wantResult [][]int
	}{
		// TODO: Add test cases.
		{
			"padsequence_test1",
			args{
				[][]int{{1}, {2, 3}, {4, 5, 6}, {7, 8, 9, 10}},
				3,
				"pre",
				"post",
				0,
			},
			[][]int{{0, 0, 1}, {0, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			"padsequence_test2",
			args{
				[][]int{{1}, {2, 3}, {4, 5, 6}, {7, 8, 9, 10}},
				3,
				"post",
				"post",
				0,
			},
			[][]int{{1, 0, 0}, {2, 3, 0}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			"padsequence_test3",
			args{
				[][]int{{1}, {2, 3}, {4, 5, 6}, {7, 8, 9, 10}},
				3,
				"post",
				"pre",
				0,
			},
			[][]int{{1, 0, 0}, {2, 3, 0}, {4, 5, 6}, {8, 9, 10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := PadIntSequences(tt.args.sequences, tt.args.maxlen, tt.args.padding, tt.args.truncating, tt.args.value); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("PadIntSequences() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

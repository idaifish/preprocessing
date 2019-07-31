package text

import (
	"reflect"
	"testing"
)

func Test_toNgram(t *testing.T) {
	type args struct {
		words []string
		ngram [2]int
	}
	tests := []struct {
		name       string
		args       args
		wantTokens []string
	}{
		// TODO: Add test cases.
		{
			"test1",
			args{
				[]string{"hello", "golang", "rust", "python", "ruby", "cpp", "java", "c", "javascript"},
				[2]int{1, 3},
			},
			[]string{"hello", "golang", "rust", "python", "ruby", "cpp", "java", "c", "javascript", "hello golang", "golang rust", "rust python", "python ruby", "ruby cpp", "cpp java", "java c", "c javascript", "hello golang rust", "golang rust python", "rust python ruby", "python ruby cpp", "ruby cpp java", "cpp java c", "java c javascript"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTokens := toNgram(tt.args.words, tt.args.ngram); !reflect.DeepEqual(gotTokens, tt.wantTokens) {
				t.Errorf("toNgram() = %#v, want %#v", gotTokens, tt.wantTokens)
			}
		})
	}
}

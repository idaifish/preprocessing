package text

import (
	"reflect"
	"testing"
)

func TestHashingTrick(t *testing.T) {
	type args struct {
		text      string
		dimension int
		hashFunc  HashFunc
		config    Config
	}
	tests := []struct {
		name          string
		args          args
		wantSequences []int64
	}{
		{
			"hashing trick test",
			args{
				"hello world hi world",
				10,
				Md5,
				NewDefaultConfig(),
			},
			[]int64{4, 1, 8, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSequences := HashingTrick(tt.args.text, tt.args.dimension, tt.args.hashFunc, tt.args.config); !reflect.DeepEqual(gotSequences, tt.wantSequences) {
				t.Errorf("HashingTrick() = %v, want %v", gotSequences, tt.wantSequences)
			}
		})
	}
}

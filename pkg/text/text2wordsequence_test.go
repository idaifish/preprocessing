package text

import (
	"reflect"
	"testing"
)

func TestTextToWordSequence(t *testing.T) {
	type args struct {
		text   string
		config Config
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			"test1",
			args{
				"hello! ? world!",
				NewDefaultConfig(),
			},
			[]string{"hello", "world"},
		},
		{
			"test2",
			args{
				"ali! veli? kırk dokuz elli",
				NewDefaultConfig(),
			},
			[]string{"ali", "veli", "kırk", "dokuz", "elli"},
		},
		{
			"test3",
			args{
				"hello! | world!",
				NewConfig(
					DefaultFilters,
					DefaultLower,
					"|",
					DefaultOOVToken,
					DefaultCharLevel,
				),
			},
			[]string{"hello", "world"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TextToWordSequence(tt.args.text, tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TextToWordSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

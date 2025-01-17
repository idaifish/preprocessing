// Package text provides utilities for text input preprocessing.
// Go port of keras.preprocessing.text
package text

import (
	"reflect"
	"testing"
)

func TestTokenizer_TextsToSequences(t *testing.T) {
	type args struct {
		texts []string
	}
	tests := []struct {
		name          string
		args          args
		wantSequences [][]int
	}{
		// TODO: Add test cases.
		{
			"text2sequences1",
			args{
				[]string{"This is a text document to analyze."},
			},
			[][]int{{9, 4, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			corpus := []string{
				"This is the first document.",
				"This is the second second document.",
				"And the third one.",
				"Is this the first document?",
			}
			tokenizer := NewTokenizer(10, NewDefaultConfig())
			tokenizer.FitOnTexts(corpus)

			if gotSequences := tokenizer.TextsToSequences(tt.args.texts); !reflect.DeepEqual(gotSequences, tt.wantSequences) {
				t.Errorf("Tokenizer.TextsToSequences() = %v, want %v", gotSequences, tt.wantSequences)
			}
		})
	}
}

func TestTokenizer_TextsToSequencesWithOOV(t *testing.T) {
	type args struct {
		texts []string
	}
	tests := []struct {
		name          string
		args          args
		wantSequences [][]int
	}{
		// TODO: Add test cases.
		{
			"text2sequences1",
			args{
				[]string{"This is a text document to analyze."},
			},
			[][]int{{1, 10, 5, 1, 1, 3, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			corpus := []string{
				"This is the first document.",
				"This is the second second document.",
				"And the third one.",
				"Is this the first document?",
			}
			tokenizer := NewTokenizer(10, NewConfig(
				DefaultFilters,
				DefaultLower,
				DefaultSplit,
				"oov",
				DefaultCharLevel,
				DefaultNgram,
			))
			tokenizer.FitOnTexts(corpus)

			if gotSequences := tokenizer.TextsToSequences(tt.args.texts); !reflect.DeepEqual(gotSequences, tt.wantSequences) {
				t.Errorf("Tokenizer.TextsToSequences() = %v, want %v", gotSequences, tt.wantSequences)
			}
		})
	}
}

func TestTokenizer_TextsToSequencesWithNgram(t *testing.T) {
	type args struct {
		texts []string
	}
	tests := []struct {
		name          string
		args          args
		wantSequences [][]int
	}{
		// TODO: Add test cases.
		{
			"text2sequences1",
			args{
				[]string{"This is a text document to analyze."},
			},
			[][]int{{27, 7, 4, 28}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			corpus := []string{
				"This is the first document.",
				"This is the second second document.",
				"And the third one.",
				"Is this the first document?",
			}
			tokenizer := NewTokenizer(10, NewConfig(
				DefaultFilters,
				DefaultLower,
				DefaultSplit,
				DefaultOOVToken,
				DefaultCharLevel,
				[2]int{1, 3},
			))
			tokenizer.FitOnTexts(corpus)
			// fmt.Printf("Vocabulary: %#v\n", tokenizer.IndexWord)

			if gotSequences := tokenizer.TextsToSequences(tt.args.texts); !reflect.DeepEqual(gotSequences, tt.wantSequences) {
				t.Errorf("Tokenizer.TextsToSequences() = %#v, want %#v", gotSequences, tt.wantSequences)
			}
		})
	}
}

func TestTokenizer_TextsToMatrix(t *testing.T) {
	type args struct {
		texts []string
		mode  string
	}
	tests := []struct {
		name string
		args args
		want [][]float64
	}{
		// TODO: Add test cases.
		{
			"text2matrix1",
			args{
				[]string{"This is a text document to analyze."},
				"tfidf",
			},
			[][]float64{
				[]float64{0, 0, 0, 0, 0.6931471805599453, 0, 0, 0.6931471805599453, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			corpus := []string{
				"This is the first document.",
				"This is the second second document.",
				"And the third one.",
				"Is this the first document?",
			}
			tokenizer := NewTokenizer(10, NewConfig(
				DefaultFilters,
				DefaultLower,
				DefaultSplit,
				DefaultOOVToken,
				DefaultCharLevel,
				[2]int{1, 3},
			))
			tokenizer.FitOnTexts(corpus)

			if got := tokenizer.TextsToMatrix(tt.args.texts, tt.args.mode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tokenizer.TextsToMatrix() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

// Package text provides utilities for text input preprocessing.
// Go port of keras.preprocess.text
package text

import (
	"errors"
	"math"
	"sort"
)

// Default flags.
const (
	DefaultFilters   = "!\"#$%&()*+,-./:;<=>?@[\\]^_`{|}~\t\n"
	DefaultLower     = true
	DefaultSplit     = " "
	DefaultOOVToken  = ""
	DefaultCharLevel = false
)

// DefaultNgram (1, 1)
var DefaultNgram = [2]int{1, 1}

// Built-in Errors
var (
	ErrValue       = errors.New("Specify a dimension or fit on some text data first")
	ErrUnknownMode = errors.New("Unknown vectorization mode")
)

// Config represents some flags of tokenizer.
type Config struct {
	Filters   string
	Lower     bool
	Split     string
	OOVToken  string
	CharLevel bool
	Ngram     [2]int
}

// NewConfig returns a custom Config.
func NewConfig(filters string, lower bool, split string, oovtoken string, charlevel bool, ngram [2]int) Config {
	return Config{
		Filters:   filters,
		Lower:     lower,
		Split:     split,
		OOVToken:  oovtoken,
		CharLevel: charlevel,
		Ngram:     ngram,
	}
}

// NewDefaultConfig returns default Config.
func NewDefaultConfig() Config {
	return Config{
		Filters:   DefaultFilters,
		Lower:     DefaultLower,
		Split:     DefaultSplit,
		OOVToken:  DefaultOOVToken,
		CharLevel: DefaultCharLevel,
		Ngram:     DefaultNgram,
	}
}

// Tokenizer allows to vectorize a text corpus, by turning each text into either a sequence of integers (each integer being the index of a token in a dictionary) or into a vector where the coefficient for each token could be binary, based on word count, based on tf-idf...
type Tokenizer struct {
	NumWords      int
	WordCounts    map[string]int
	WordDocs      map[string]int
	DocumentCount int
	WordIndex     map[string]int
	IndexWord     map[int]string
	IndexDocs     map[int]int
	Config        Config
}

// NewTokenizer returns a tokenizer pointer.
func NewTokenizer(numWords int, config Config) *Tokenizer {
	return &Tokenizer{
		NumWords:      numWords,
		WordCounts:    map[string]int{},
		WordDocs:      map[string]int{},
		DocumentCount: 0,
		WordIndex:     map[string]int{},
		IndexWord:     map[int]string{},
		IndexDocs:     map[int]int{},
		Config:        config,
	}
}

// FitOnTexts updates internal vocabulary based on a list of texts.
func (tokenizer *Tokenizer) FitOnTexts(texts []string) {
	for _, text := range texts {
		tokenizer.DocumentCount++

		wordSeq := TextToWordSequence(text, tokenizer.Config)
		for _, word := range wordSeq {
			if _, ok := tokenizer.WordCounts[word]; ok {
				tokenizer.WordCounts[word]++
			} else {
				tokenizer.WordCounts[word] = 1
			}
		}

		for _, word := range uniqueStrings(wordSeq) {
			tokenizer.WordDocs[word]++
		}
	}

	wordList := make(kvList, 0)
	for k, v := range tokenizer.WordCounts {
		wordList = append(wordList, kv{k, v})
	}
	sort.Sort(wordList)
	sort.SliceStable(wordList, func(i, j int) bool {
		return wordList[i].Key < wordList[j].Key
	})

	var sortedWords []string
	if tokenizer.Config.OOVToken != "" {
		sortedWords = []string{tokenizer.Config.OOVToken}
	}
	for _, kv := range wordList {
		sortedWords = append(sortedWords, kv.Key)
	}

	for index, word := range sortedWords {
		tokenizer.IndexWord[index+1] = word
	}
	for index, word := range tokenizer.IndexWord {
		tokenizer.WordIndex[word] = index
	}

	for word, docs := range tokenizer.WordDocs {
		tokenizer.IndexDocs[tokenizer.WordIndex[word]] = docs
	}
}

// FitOnSequences updates internal vocabulary based on a list of sequences.
func (tokenizer *Tokenizer) FitOnSequences(sequences []int) {
	tokenizer.DocumentCount += len(sequences)
	for s := range uniqueInts(sequences) {
		tokenizer.IndexDocs[s]++
	}
}

// TextsToSequences transforms each text in texts to a sequence of integers.
func (tokenizer *Tokenizer) TextsToSequences(texts []string) (sequences [][]int) {
	oovTokenIndex, oovTokenIndexOK := tokenizer.WordIndex[tokenizer.Config.OOVToken]

	for _, text := range texts {
		wordSeqs := TextToWordSequence(text, tokenizer.Config)
		sequence := []int{}
		for _, word := range wordSeqs {
			if wordIndex, ok := tokenizer.WordIndex[word]; ok {
				// out of vocabulary
				if wordIndex >= tokenizer.NumWords && oovTokenIndexOK {
					sequence = append(sequence, oovTokenIndex)
				}
				sequence = append(sequence, wordIndex)
			} else if oovTokenIndexOK {
				sequence = append(sequence, oovTokenIndex)
			}
		}

		sequences = append(sequences, sequence)
	}

	return
}

// SequencesToMatrix converts a list of sequences into matrix [][]float64.
func (tokenizer *Tokenizer) SequencesToMatrix(sequences [][]int, mode string) (matrix [][]float64) {
	if tokenizer.NumWords == 0 && len(tokenizer.WordIndex) == 0 {
		panic(ErrValue)
	}

	if mode == "tfidf" && tokenizer.DocumentCount == 0 {
		panic(ErrValue)
	}

	for i, sequence := range sequences {
		if len(sequence) == 0 {
			continue
		}

		counts := map[int]int{}
		for _, seq := range sequence {
			if seq >= tokenizer.NumWords {
				continue
			}
			counts[seq]++

			for j, c := range counts {
				switch mode {
				case "count":
					matrix[i][j] = float64(c)
				case "freq":
					matrix[i][j] = float64(c / len(sequence))
				case "binary":
					matrix[i][j] = float64(1)
				case "tfidf":
					tf := 1 + math.Log(float64(c))
					var idf float64
					if docs, ok := tokenizer.IndexDocs[j]; ok {
						idf = math.Log(float64(1 + tokenizer.DocumentCount/(1+docs)))
					} else {
						idf = math.Log(float64(1 + tokenizer.DocumentCount))
					}
					matrix[i][j] = tf * idf
				default:
					panic(ErrUnknownMode)
				}
			}
		}
	}

	return
}

// TextsToMatrix convert a list of texts to matrix [][]float64.
func (tokenizer *Tokenizer) TextsToMatrix(texts []string, mode string) [][]float64 {
	sequences := tokenizer.TextsToSequences(texts)

	return tokenizer.SequencesToMatrix(sequences, mode)
}

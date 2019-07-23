package text

import "sort"

const (
	DefaultFilters   = "!\"#$%&()*+,-./:;<=>?@[\\]^_`{|}~\t\n"
	DefaultLower     = true
	DefaultSeparator = " "
	DefaultOOVToken  = ""
	DefaultCharLevel = false
)

type Config struct {
	Filters   string
	Lower     bool
	Separator string
	OOVToken  string
	CharLevel bool
}

func NewConfig(filters string, lower bool, separator string, oovtoken string, charlevel bool) *Config {
	return &Config{
		Filters:   filters,
		Lower:     lower,
		Separator: separator,
		OOVToken:  oovtoken,
		CharLevel: charlevel,
	}
}

func NewDefaultConfig() Config {
	return Config{
		Filters:   DefaultFilters,
		Lower:     DefaultLower,
		Separator: DefaultSeparator,
		OOVToken:  DefaultOOVToken,
		CharLevel: DefaultCharLevel,
	}
}

// Tokenizer allows to vectorize a text corpus, by turning each text into either a sequence of integers (each integer being the index of a token in a dictionary) or into a vector where the coefficient for each token could be binary, based on word count, based on tf-idf...
type Tokenizer struct {
	NumWords      int32
	WordCounts    map[string]int
	WordDocs      map[string]int
	DocumentCount int
	WordIndex     map[string]int
	IndexWord     map[int]string
	IndexDocs     map[int]int
	Config        Config
}

func NewTokenizer(numWords int32) *Tokenizer {
	return &Tokenizer{
		NumWords:      numWords,
		WordCounts:    make(map[string]int, 1),
		WordDocs:      make(map[string]int, 1),
		DocumentCount: 0,
		WordIndex:     make(map[string]int, 1),
		IndexWord:     make(map[int]string, 1),
		IndexDocs:     make(map[int]int, 1),
		Config:        NewDefaultConfig(),
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

		kvList := make(KVList, len(tokenizer.WordCounts))
		for k, v := range tokenizer.WordCounts {
			kvList = append(kvList, KV{k, v})
		}
		sort.Sort(kvList)

		var sortedWords []string
		if tokenizer.Config.OOVToken != "" {
			sortedWords = []string{tokenizer.Config.OOVToken}
		}
		for _, kv := range kvList {
			sortedWords = append(sortedWords, kv.Key)
		}

		for index, word := range sortedWords {
			tokenizer.IndexWord[index+1] = word
			tokenizer.WordIndex[word] = index + 1
		}

		for word, docs := range tokenizer.WordDocs {
			tokenizer.IndexDocs[tokenizer.WordIndex[word]] = docs
		}
	}

}

func (tokenizer *Tokenizer) FitOnSequences(sequences []int) {
	tokenizer.DocumentCount += len(sequences)
	for s := range uniqueInts(sequences) {
		tokenizer.IndexDocs[s]++
	}
}

func (tokenizer *Tokenizer) TextsToSequences(texts []string) {

}

func (tokenizer *Tokenizer) SequencesToTexts(texts []string) {

}

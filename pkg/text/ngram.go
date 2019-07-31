package text

import (
	"math"
	"strings"
)

func toNgram(words []string, ngram [2]int) (tokens []string) {
	nMin, nMax := ngram[0], ngram[1]
	wordsLen := len(words)

	if nMin > nMax || nMin < 1 {
		panic("Invalid value for ngram")
	}

	if nMax == 1 {
		tokens = words
		return words
	}

	if nMin == 1 {
		tokens = words
		nMin++
	}

	for n := nMin; n < int(math.Min(float64(nMax+1), float64(wordsLen+1))); n++ {
		for i := 0; i < wordsLen-n+1; i++ {
			tokens = append(tokens, strings.Join(words[i:i+n], " "))
		}
	}

	return
}

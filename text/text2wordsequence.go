package text

import (
	"strings"
)

// TextToWordSequence converts a text to a sequence of words.
func TextToWordSequence(text string, config Config) (result []string) {
	if config.Lower {
		text = strings.ToLower(text)
	}

	var translateMap = []string{}
	for _, c := range config.Filters {
		translateMap = append(translateMap, string(c), config.Split)
	}

	replacer := strings.NewReplacer(translateMap...)
	text = replacer.Replace(text)

	if config.CharLevel {
		text = strings.Replace(text, config.Split, "", -1)
		for _, c := range text {
			if string(c) != " " {
				result = append(result, string(c))
			}
		}
	} else {
		text = strings.Replace(text, config.Split, " ", -1)
		result = strings.Fields(text)
	}

	if config.Ngram != DefaultNgram {
		result = toNgram(result, config.Ngram)
	}

	return
}

package text

import (
	"strings"
)

// TextToWordSequence converts a text to a sequence of words.
func TextToWordSequence(text string, config Config) []string {
	if config.Lower {
		text = strings.ToLower(text)
	}

	var translateMap = []string{}
	for _, c := range config.Filters {
		translateMap = append(translateMap, string(c), config.Split)
	}

	replacer := strings.NewReplacer(translateMap...)
	text = replacer.Replace(text)
	text = strings.Replace(text, config.Split, " ", -1)

	return strings.Fields(text)
}

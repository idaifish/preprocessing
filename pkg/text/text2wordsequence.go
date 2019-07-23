package text

import (
	"strings"
)

func TextToWordSequence(text string, config Config) []string {
	if config.Lower {
		text = strings.ToLower(text)
	}

	var translateMap = []string{}
	for _, c := range config.Filters {
		translateMap = append(translateMap, string(c), config.Separator)
	}

	replacer := strings.NewReplacer(translateMap...)
	text = replacer.Replace(text)
	text = strings.Replace(text, config.Separator, " ", -1)

	return strings.Fields(text)
}

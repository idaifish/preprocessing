# Preprocessing

[![GoDoc](https://godoc.org/github.com/idaifish/preprocessing?status.svg)](https://godoc.org/github.com/idaifish/preprocessing) [![Go Report Card](https://goreportcard.com/badge/github.com/idaifish/preprocessing)](https://goreportcard.com/report/github.com/idaifish/preprocessing) [![Build Status](https://travis-ci.com/idaifish/preprocessing.svg?branch=master)](https://travis-ci.com/idaifish/preprocessing)

Data preprocessing library for machine learning.

## Installation

```
$ go get github.com/idaifish/preprocessing/...
```

## Get Started

```go
import (
	"fmt"

	"github.com/idaifish/preprocessing/text"
)

func main() {
	corpus := []string{
		"This is the first document.",
		"This is the second second document.",
		"And the third one.",
		"Is this the first document?",
	}
	tokenizer := text.NewTokenizer(10, text.NewDefaultConfig())
	tokenizer.FitOnTexts(corpus)

	fmt.Println(tokenizer.TextsToSequences([]string{"This is a text document to analyze."}))
}
```

```go
import (
	"fmt"

	"github.com/idaifish/preprocessing/sequence"
)

func main() {
	intSeq := [][]int{{1}, {2, 3}, {4, 5, 6}, {7, 8, 9, 10}}
	fmt.Println(sequence.PadSequences(intSeq, 3, "post", "pre", 0))
}
```
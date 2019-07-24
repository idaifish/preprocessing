package text

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"strconv"
)

// HashFunc is the type of function called for hashingtrick.
type HashFunc func(string) int

// Md5 hash function.
func Md5(text string) int {
	h := md5.New()
	io.WriteString(h, text)
	hashInt, err := strconv.ParseInt((hex.EncodeToString(h.Sum(nil))[:15]), 16, 64)
	if err != nil {
		panic(err)
	}

	return int(hashInt)
}

// HashingTrick converts a text to sequence of indexes.
func HashingTrick(text string, dimension int, hashFunc HashFunc, config Config) (sequences []int) {
	if hashFunc == nil {
		panic(errors.New("No HashFunc"))
	}

	sequence := TextToWordSequence(text, config)
	for _, seq := range sequence {
		sequences = append(sequences, hashFunc(seq)%(dimension-1)+1)
	}

	return
}

// OneHot encodes a text into a list of word indexes of size dimension
func OneHot(text string, dimension int, config Config) (sequences []int) {
	return HashingTrick(text, dimension, Md5, config)
}

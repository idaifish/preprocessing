package text

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"math/big"
)

// HashFunc is the type of function called for hashingtrick.
type HashFunc func(string) big.Int

// Md5 hash function.
func Md5(text string) (hashInt big.Int) {
	h := md5.New()
	io.WriteString(h, text)
	hashInt.SetString((hex.EncodeToString(h.Sum(nil))), 16)

	return
}

// HashingTrick converts a text to sequence of indexes.
func HashingTrick(text string, dimension int, hashFunc HashFunc, config Config) (sequences []int) {
	if hashFunc == nil {
		panic(errors.New("No HashFunc"))
	}

	sequence := TextToWordSequence(text, config)
	for _, seq := range sequence {
		hashInt := hashFunc(seq)
		hashInt.Mod(&hashInt, new(big.Int).SetInt64(int64(dimension-1)))
		sequences = append(sequences, int(hashInt.Int64())+1)
	}

	return
}

// OneHot encodes a text into a list of word indexes of size dimension
func OneHot(text string, dimension int, config Config) (sequences []int) {
	return HashingTrick(text, dimension, Md5, config)
}

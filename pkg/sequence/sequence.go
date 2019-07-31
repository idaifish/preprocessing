// Package sequence provides utilities for preprocessing sequence data.
// Go port of keras.preprocessing.sequence
package sequence

// PadSequences pads sequences to the same length.
var PadSequences = PadIntSequences

// PadIntSequences pads sequences to the same length.
func PadIntSequences(sequences [][]int, maxlen int, padding string, truncating string, value int) (result [][]int) {
	var lengths []int
	for _, sequence := range sequences {
		lengths = append(lengths, len(sequence))
	}
	maxLen := max(lengths)
	if maxlen <= 0 {
		maxlen = maxLen
	}

	for i, sequence := range sequences {
		switch seqLen := lengths[i]; {
		case seqLen > maxlen:
			switch truncating {
			case "post":
				result = append(result, sequence[:maxlen])
			case "pre":
				result = append(result, sequence[seqLen-maxlen:])
			default:
				panic("`truncating` should be `pre` or `post`.")
			}
		case seqLen < maxlen:
			pads := make([]int, maxlen-seqLen)
			for i = 0; i < (maxlen - seqLen); i++ {
				pads[i] = value
			}

			switch padding {
			case "post":
				sequence = append(sequence, pads...)
				result = append(result, sequence)
			case "pre":
				sequence = append(pads, sequence...)
				result = append(result, sequence)
			default:
				panic("`padding` should be `pre` or `post`.")
			}
		default:
			result = append(result, sequence)
		}
	}

	return
}

// PadFloat64Sequences pads sequences to the same length.
func PadFloat64Sequences(sequences [][]float64, maxlen int, padding string, truncating string, value float64) (result [][]float64) {
	var lengths []int
	for _, sequence := range sequences {
		lengths = append(lengths, len(sequence))
	}
	maxLen := max(lengths)
	if maxlen <= 0 {
		maxlen = maxLen
	}

	for i, sequence := range sequences {
		switch seqLen := lengths[i]; {
		case seqLen > maxlen:
			switch truncating {
			case "post":
				result = append(result, sequence[:maxlen])
			case "pre":
				result = append(result, sequence[seqLen-maxlen:])
			default:
				panic("`truncating` should be `pre` or `post`.")
			}
		case seqLen < maxlen:
			pads := make([]float64, maxlen-seqLen)
			for i = 0; i < (maxlen - seqLen); i++ {
				pads[i] = value
			}

			switch padding {
			case "post":
				sequence = append(sequence, pads...)
				result = append(result, sequence)
			case "pre":
				sequence = append(pads, sequence...)
				result = append(result, sequence)
			default:
				panic("`padding` should be `pre` or `post`.")
			}
		default:
			result = append(result, sequence)
		}
	}

	return
}

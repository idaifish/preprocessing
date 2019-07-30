package sequence

func max(ints []int) (max int) {
	for _, i := range ints {
		if i > max {
			max = i
		}
	}
	return
}

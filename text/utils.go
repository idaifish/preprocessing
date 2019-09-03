package text

func uniqueStrings(s []string) []string {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] != s[j] {
				break
			}
			s = append(s[:i], s[j:]...)
		}
	}
	return s
}

func uniqueInts(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] != s[j] {
				break
			}
			s = append(s[:i], s[j:]...)
		}
	}
	return s
}

type kv struct {
	Key   string
	Value int
}

type kvList []kv

func (l kvList) Len() int           { return len(l) }
func (l kvList) Less(i, j int) bool { return l[i].Value < l[j].Value }
func (l kvList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

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

type KV struct {
	Key   string
	Value int
}

type KVList []KV

func (l KVList) Len() int           { return len(l) }
func (l KVList) Less(i, j int) bool { return l[i].Value < l[j].Value }
func (l KVList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

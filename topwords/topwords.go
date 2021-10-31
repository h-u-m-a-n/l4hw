package topwords

import (
	"regexp"
	"sort"
	"strings"
)

type Word struct {
	W string
	N int
}

func TopWords(s string, n int) []Word{
	s = strings.ToLower(s)
	reg := regexp.MustCompile(`[{,.!?}\s]+`)
	words := reg.Split(s, -1)
	m := map[string]int{}
	for _, v := range words {
		m[v]++
	}
	res := make([]Word, 0, len(m))
	for i, v := range m {
		if len(i) < 1 {
			continue
		}
		res = append(res, Word{i, v})
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].N > res[j].N
	})

	return res[:n]
}

package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const (
	topNum = 10
)

type Word struct {
	name  string
	count int
}

func Top10(str string) []string {
	words := make(map[string]int)
	wordsSlice := []*Word{}

	s := []string{}

	if len(str) < topNum {
		return nil
	}

	list := strings.Fields(str)

	for _, word := range list {
		words[word]++
	}

	for k, v := range words {
		w := &Word{name: k, count: v}
		wordsSlice = append(wordsSlice, w)
	}

	sort.Slice(wordsSlice, func(i, j int) bool {
		if wordsSlice[i].count == wordsSlice[j].count {
			return wordsSlice[i].name < wordsSlice[j].name
		}
		return wordsSlice[i].count > wordsSlice[j].count
	})

	for i := 0; i < topNum; i++ {
		s = append(s, wordsSlice[i].name)
	}

	return s
}

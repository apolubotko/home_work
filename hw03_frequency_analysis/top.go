package hw03frequencyanalysis

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

const (
	topNum = 10
)

var (
	errWrongWord = errors.New("not a word")
	r            = regexp.MustCompile(`(\p{L}+|\w+)\-?(\p{L}+|\w+)?`)
)

type Word struct {
	name  string
	count int
}

func Top10(str string) []string {
	words := make(map[string]int)
	wordsSlice := []*Word{}

	s := []string{}
	list := strings.Fields(str)

	for _, word := range list {
		w, err := makeClean(word)
		if err != nil {
			continue
		}
		words[w]++
	}

	if len(words) < topNum {
		return nil
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

func makeClean(s string) (string, error) {
	lowerStr := strings.ToLower(s)

	result := r.FindAllStringSubmatch(lowerStr, -1)

	if r.MatchString(lowerStr) {
		return result[0][0], nil
	}

	return "", errWrongWord
}

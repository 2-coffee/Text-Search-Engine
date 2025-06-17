package utils

import (
	"log"
	"strings"

	"github.com/kljensen/snowball"
)

// Convert all tokens to lowercase
func lowercaseFilter(tokens []string) []string {
	res := make([]string, len(tokens))
	for i, token := range tokens {
		res[i] = strings.ToLower(token)
	}
	return res
}

func stopwordFilter(tokens []string) []string {
	var stopwords = map[string]struct{}{
		"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
		"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
	}

	res := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if _, ok := stopwords[token]; !ok {
			res = append(res, token)
		}
	}

	return res
}

func stemmerFilter(tokens []string) []string {
	res := make([]string, len(tokens))

	for i, token := range tokens {
		word, err := snowball.Stem(token, "english", false)
		if err != nil {
			log.Fatal(err)
		}
		res[i] = word
	}
	return res
}

package gonbayes

import (
	"fmt"
	"github.com/kljensen/snowball"
	"regexp"
	"strings"
)

var re = regexp.MustCompile("[^a-zA-Z 0-9]+")

func isStopWords(word string) bool {
	_, ok := stopWords[word]
	return ok
}

func stem(word string) string {
	stemmed, err := snowball.Stem(word, "english", true)
	if err != nil {
		fmt.Printf("cannot stem word: %s", word)
		return word
	}
	return stemmed
}

func clean(document string) string {
	document = strings.ToLower(document)
	return re.ReplaceAllString(document, "")
}

package gonbayes

func isStopWords(word string) bool {
	_, ok := stopWords[word]
	return ok
}

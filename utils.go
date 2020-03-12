package gonbayes

func isStopWord(word string) bool {
	_, ok := stopWords[word]
	return ok
}

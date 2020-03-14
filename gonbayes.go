package gonbayes

// Classifier is document categories classifier
type Classifier struct {
	Words                  map[string]map[string]uint64 // カテゴリそれぞれの各単語の数
	TotalWords             uint64                       // 全単語の総数
	TotalDocsInCategories  map[string]uint64            // 各カテゴリに何個ドキュメントが格納されているかのデータ
	TotalDocs              uint64                       // ドキュメントの総数
	TotalWordsInCategories map[string]uint64            //カテゴリ毎に単語がいくつあるかのデータ

}

// NewClassifier initializer classifier
func NewClassifier(categories []string) *Classifier {
	c := &Classifier{
		Words:                  make(map[string]map[string]uint64),
		TotalDocsInCategories:  make(map[string]uint64),
		TotalWordsInCategories: make(map[string]uint64),
	}

	for _, category := range categories {
		c.Words[category] = make(map[string]uint64)
	}
	return c
}

func (c *Classifier) Train(category string, document string) {
	for word, count := range countWords(document) {
		c.Words[category][word] += uint64(count)
		c.TotalWordsInCategories[category] += uint64(count)
		c.TotalWords += uint64(count)
	}
	c.TotalDocsInCategories[category]++
	c.TotalDocs++
}

func (c *Classifier) Classify(document string) (category string) {
	// TODO: implement
}

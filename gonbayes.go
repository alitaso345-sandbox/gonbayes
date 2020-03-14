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

// Trains documents classifier
func (c *Classifier) Train(category string, document string) {
	for word, count := range countWords(document) {
		c.Words[category][word] += uint64(count)
		c.TotalWordsInCategories[category] += uint64(count)
		c.TotalWords += uint64(count)
	}
	c.TotalDocsInCategories[category]++
	c.TotalDocs++
}

// P(C)の計算
func (c *Classifier) pCategory(category string) float64 {
	return float64(c.TotalDocsInCategories[category]) / float64(c.TotalDocs)
}

// P(D|C)の計算
func (c *Classifier) pDocCategory(category string, document string) float64 {
	p := 1.0
	for word := range countWords(document) {
		p *= c.pWordCategory(category, word)
	}
	return p
}

// P(w|C)の計算（P(D|C)の計算のため）
func (c *Classifier) pWordCategory(category string, word string) float64 {
	d := float64(c.Words[category][stem(word)])
	n := float64(c.TotalWordsInCategories[category])
	return d / n
}

// P(C|D)の計算
func (c *Classifier) pCategoryDocument(category string, document string) float64 {
	return c.pDocCategory(category, document) * c.pCategory(category)
}

// P is Probabilities of each categories
func (c *Classifier) P(document string) map[string]float64 {
	p := make(map[string]float64)
	for category := range c.Words {
		p[category] = c.pCategoryDocument(category, document)
	}
	return p
}

func (c *Classifier) Classify(document string) (category string) {
	// TODO: implement
}

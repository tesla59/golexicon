package golexicon

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Lexicon struct {
	Count     int
	Separator string
	Words     []string
}

func NewLexicon() Lexicon {
	return Lexicon{
		Count:     2,
		Separator: "-",
		Words:     readDict(),
	}
}

func (l Lexicon) Generate() string {
	words := []string{}
	for i := 0; i < l.Count; i++ {
		words = append(words, l.Words[rand.Intn(len(l.Words))])
	}
	return strings.Join(words, l.Separator)
}

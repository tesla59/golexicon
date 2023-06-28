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

type Config struct {
	Count     int
	Separator string
}

func NewLexicon(config ...*Config) Lexicon {
	lexicon := Lexicon{
		Count:     2,
		Separator: "-",
		Words:     readDict(),
	}
	if len(config) > 0 {
		if config[0].Count > 0 {
			lexicon.Count = config[0].Count
		}
		if config[0].Separator != "" {
			lexicon.Separator = config[0].Separator
		}
	}
	return lexicon
}


func (l Lexicon) Generate() string {
	words := []string{}
	for i := 0; i < l.Count; i++ {
		words = append(words, l.Words[rand.Intn(len(l.Words))])
	}
	return strings.Join(words, l.Separator)
}

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tesla59/golexicon"
)

var (
	separator string
	words     int
)

func init() {
	flag.IntVar(&words, "n", 2, "number of random words to generate")
	flag.StringVar(&separator, "s", "-", "separator to seprate generated words")
}

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options]\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	lexicon := golexicon.NewLexicon()
	lexicon.Count = words
	lexicon.Separator = separator

	println(lexicon.Generate())
	return
}

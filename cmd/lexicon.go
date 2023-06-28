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
	if len(os.Args) > 1 {
		checkUsage()
	}

	flag.Parse()
	lexicon := golexicon.NewLexicon()
	lexicon.Count = words
	lexicon.Separator = separator

	println(lexicon.Generate())
	return
}

func checkUsage() {
	if os.Args[1] == "help" || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: lexicon -s [separator] -n [number-of-words] \n eg: lexicon -n=6 # pisolitic-abkari-barmier-seceder-preeternity-saveableness")
		os.Exit(1)
	}
}

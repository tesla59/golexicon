# GoLexicon
GoLexicon is used to generate random words, written in Golang and inspired by Babble

Babble depends on OS's dictionary word, GoLexicon has integrated dictionary so it doesn't depend on platform

## Installation
Tested with Go 1.20
``` go get github.com/tesla59/golexicon ```

## Usage
```
package main

import (
	"fmt"

	"github.com/tesla59/golexicon"
)

func main() {
	// Generate a new lexicon
	lexi := golexicon.NewLexicon(&golexicon.Config{})
	fmt.Println(lexi.Generate()) // hypnotisability-alcoholic

	// Set Count to change the number of words generated. Default 2
	lexi.Count = 5
	fmt.Println(lexi.Generate()) // ostiolar-fdub-cardaissin-marmennill-incontrollable

	// Set Separator to change the separation character between words. Defauly "-"
	lexi.Separator = "_"
	fmt.Println(lexi.Generate()) // nisei_botryopteriaceae_cage_anginal_dejeuner
}
```

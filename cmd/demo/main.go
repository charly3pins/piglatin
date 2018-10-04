package main

import (
	"flag"
	"log"
	"strings"

	"github.com/charly3pins/piglatin"
)

func main() {
	txt := flag.String("text", "", "text to translate")
	flag.Parse()

	defaultHandler := piglatin.DefaultHandler{}
	quHandler := piglatin.QuHandler{Next: &defaultHandler}
	vowelHandler := piglatin.VowelHandler{Next: &quHandler}
	hyphenHandler := piglatin.HyphenHandler{Next: &vowelHandler}
	symbolsHandler := piglatin.SymbolsHandler{Next: &hyphenHandler}

	myText := piglatin.NewText(*txt)
	result := piglatin.Text{}
	for _, word := range myText {
		word = symbolsHandler.Translate(word)
		result = append(result, word)
	}
	log.Println(result)

	translation := []string{}
	for _, word := range result {
		word.Title()
		translation = append(translation, word.GetTranslated())
	}

	log.Println(strings.Join(translation, " "))
}

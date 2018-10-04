package piglatin

import (
	"testing"
)

func TestTranslate_symbolOK(t *testing.T) {
	defaultHandler := DefaultHandler{}
	quHandler := QuHandler{Next: &defaultHandler}
	vowelHandler := VowelHandler{Next: &quHandler}
	hyphenHandler := HyphenHandler{Next: &vowelHandler}
	symbolsHandler := SymbolsHandler{Next: &hyphenHandler}

	w := &word{Original: "takeMySymbol!"}
	res := symbolsHandler.Translate(w)
	exepected := "akeMySymboltay!"
	if res.GetTranslated() != exepected {
		t.Errorf("Expecting %s obtained %s", exepected, res.GetTranslated())
	}
}

func TestTranslate_symbolKO(t *testing.T) {
	symbolH := SymbolsHandler{Next: &MockNextHandler{}}
	w := &word{Original: "noSymbolHere"}
	res := symbolH.Translate(w)
	if res != nil {
		t.Errorf("Error translating non-vowel word in vowel method")
	}
}

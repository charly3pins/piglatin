package piglatin

import (
	"testing"
)

func TestTranslate_hyphenOK(t *testing.T) {
	defaultHandler := DefaultHandler{}
	quHandler := QuHandler{Next: &defaultHandler}
	vowelHandler := VowelHandler{Next: &quHandler}
	hyphenHandler := HyphenHandler{Next: &vowelHandler}

	w := &word{Original: "state-of-the-art"}
	res := hyphenHandler.Translate(w)
	exepected := "atestay-ofway-ethay-artway"
	if res.GetTranslated() != exepected {
		t.Errorf("Expecting %s obtained %s", exepected, res.GetTranslated())
	}
}

func TestTranslate_hyphenKO(t *testing.T) {
	hyphenH := HyphenHandler{Next: &MockNextHandler{}}
	w := &word{Original: "noHyphenWord"}
	res := hyphenH.Translate(w)
	if res != nil {
		t.Errorf("Error translating non-vowel word in vowel method")
	}
}

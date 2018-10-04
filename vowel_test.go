package piglatin

import (
	"testing"
)

func TestTranslate_vowelOK(t *testing.T) {
	vowelH := VowelHandler{}
	w := &word{Original: "aleluya"}
	res := vowelH.Translate(w)
	exepected := "aleluyaway"
	if res.GetTranslated() != exepected {
		t.Errorf("Expecting %s obtained %s", exepected, res.GetTranslated())
	}
}

func TestTranslate_vowelKO(t *testing.T) {
	vowelH := VowelHandler{Next: &MockNextHandler{}}
	w := &word{Original: "quiet"}
	res := vowelH.Translate(w)
	if res != nil {
		t.Errorf("Error translating non-vowel word in vowel method")
	}

}

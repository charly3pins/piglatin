package piglatin

import "testing"

func TestTranslate_quOK(t *testing.T) {
	quH := QuHandler{}
	w := &word{Original: "quiet"}
	res := quH.Translate(w)
	exepected := "ietquay"
	if res.GetTranslated() != exepected {
		t.Errorf("Expecting %s obtained %s", exepected, res.GetTranslated())
	}
}

func TestTranslate_quKO(t *testing.T) {
	quH := QuHandler{Next: &MockNextHandler{}}
	w := &word{Original: "whereIsMyQU"}
	res := quH.Translate(w)
	if res != nil {
		t.Errorf("Error translating non-vowel word in vowel method")
	}
}

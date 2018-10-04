package piglatin

import (
	"testing"
)

type MockNextHandler struct{}

func (m *MockNextHandler) Translate(w Word) Word {
	return nil
}

func TestTranslate_defaultOK(t *testing.T) {
	defaultH := DefaultHandler{}
	w := &word{Original: "quiet"}
	res := defaultH.Translate(w)
	exepected := "uietqay"
	if res.GetTranslated() != exepected {
		t.Errorf("Expecting %s obtained %s", exepected, res.GetTranslated())
	}
}

func TestTranslate_defaultYLikeVowelOK(t *testing.T) {
	defaultH := DefaultHandler{Next: &MockNextHandler{}}
	w := &word{Original: "xyz"}
	res := defaultH.Translate(w)
	exepected := "yzxay"
	if res.GetTranslated() != exepected {
		t.Errorf("Expecting %s obtained %s", exepected, res.GetTranslated())
	}
}

func TestTranslate_defaultMayusOK(t *testing.T) {
	defaultH := DefaultHandler{Next: &MockNextHandler{}}
	w := &word{Original: "Hello"}
	res := defaultH.Translate(w)
	exepected := "Ellohay"
	if res.GetTranslated() != exepected {
		t.Errorf("Expecting %s obtained %s", exepected, res.GetTranslated())
	}
}

func TestTranslate_defaultKO(t *testing.T) {
	defaultH := DefaultHandler{Next: &MockNextHandler{}}
	w := &word{Original: "zZZZzZZzZzz"}
	res := defaultH.Translate(w)
	if res != nil {
		t.Errorf("Expecting nil obtained %s", res.GetTranslated())
	}
}

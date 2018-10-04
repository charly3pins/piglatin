package piglatin

import "testing"

func TestNewText_OK(t *testing.T) {
	txt := NewText("hello world")
	if len(txt) != 2 {
		t.Errorf("Expecting 2 elements in text created and have %d", len(txt))
	}

	if "hello" != txt[0].GetOriginal() {
		t.Errorf("Expecting `hello` in position `0` and have %s", txt[0].GetOriginal())
	}

	if "world" != txt[1].GetOriginal() {
		t.Errorf("Expecting `world` in position `1` and have %s", txt[1].GetOriginal())
	}
}

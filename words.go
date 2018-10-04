package piglatin

import (
	"fmt"
	"strings"
)

// Text is a set of words
type Text []Word // TODO CHECK INTERFACE

// Word contains the methods for help to translate the original word
type Word interface {
	Vowel()
	Consonant(string)
	QU()
	Title()
	GetOriginal() string
	GetTranslated() string
	SetTranslated(string)
}

// NewText groups the splitted words in a Text type
func NewText(text string) Text {
	t := Text{}
	words := strings.Split(text, " ")
	for _, w := range words {
		t = append(t, &word{Original: w})
	}

	return t
}

type word struct {
	Original   string
	Translated string
}

// Vowel translate the words that start with vowel
func (w *word) Vowel() {
	w.Translated = fmt.Sprintf("%sway", w.Original)
	w.Title()
}

// Consonant translate the words that doesn't start with vowel or match with other rules
func (w *word) Consonant(partial string) {
	w.Translated = fmt.Sprintf("%s%say", string(w.Original[len(partial):]), partial)
	w.Title()
}

// QU translate the words that start with qu
func (w *word) QU() {
	w.Consonant("qu")
}

// Title makes the upperCase in the word if the rule applies
func (w *word) Title() {
	if ok, err := isTitle(w.Original); err == nil && ok {
		w.Translated = strings.Title(strings.ToLower(w.Translated))
	}
}

// GetOriginal return the original word
func (w word) GetOriginal() string {
	return w.Original
}

// GetTranslated return the word translated
func (w word) GetTranslated() string {
	return w.Translated
}

// SetTranslated stores the word translated in the object
func (w *word) SetTranslated(translated string) {
	w.Translated = translated
}

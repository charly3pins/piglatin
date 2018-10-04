package piglatin

import "log"

// Handler is the generic interface used to translate words
type Handler interface {
	Translate(Word) Word
}

// DefaultHandler is the handler that processes the words that doesn't match with other rules
type DefaultHandler struct {
	Next Handler
}

// Translate is the method to translate the words that doesn't match with other rules
func (d *DefaultHandler) Translate(w Word) Word {
	i := firstVowelIndex(w.GetOriginal())
	if i == -1 {
		log.Printf("DefaultHandler: [%s] not processed", w.GetOriginal())
		return nil
	}

	w.Consonant(string(w.GetOriginal()[:i]))
	return w
}

func firstVowelIndex(word string) int {
	for i, c := range word {
		ok, err := startVowel(string(c))
		if err != nil {
			return -1
		}
		if ok {
			return i
		}
		if i > 0 && "y" == string(c) {
			return i
		}
	}

	return -1
}

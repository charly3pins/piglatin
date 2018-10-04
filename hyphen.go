package piglatin

import (
	"log"
	"strings"
)

// HyphenHandler is the handler that processes the words that contains hyphens
type HyphenHandler struct {
	Next Handler
}

// Translate is the method to translate the words containing hyphens
func (h *HyphenHandler) Translate(w Word) Word {
	hyphenWord := strings.Split(w.GetOriginal(), "-")
	if len(hyphenWord) <= 1 {
		log.Printf("HyphenHandler:[%s] not for me", w.GetOriginal())
		return h.Next.Translate(w)
	}

	hs := []string{}
	for _, v := range hyphenWord {
		s := h.Next.Translate(&word{Original: v})
		hs = append(hs, s.GetTranslated())
	}
	w.SetTranslated(strings.Join(hs, "-"))

	return w
}

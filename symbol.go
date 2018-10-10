package piglatin

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// SymbolsHandler is the handler that processes the words that contains symbols
type SymbolsHandler struct {
	Next Handler
}

// Translate is the method to translate the words containing symbols
func (s *SymbolsHandler) Translate(w Word) Word {
	r, err := regexp.Compile(`[^a-zA-z\-\s]+`)
	if err != nil {
		log.Println("Error parsing regex", err)
		return &word{}
	}
	matches := r.FindAllString(w.GetOriginal(), -1)
	if len(matches) == 0 {
		return s.Next.Translate(w)
	}

	ss := []string{}
	for _, symbol := range matches {
		sanitized := strings.Split(w.GetOriginal(), symbol)
		if len(sanitized) > 0 && len(sanitized[0]) <= 0 {
			w.SetTranslated(symbol)
			return w
		}

		s := s.Next.Translate(&word{Original: sanitized[0]})
		ss = append(ss, fmt.Sprintf("%s%s", s.GetTranslated(), symbol))
	}
	w.SetTranslated(strings.Join(ss, ""))

	return w
}

package piglatin

import "log"

// QuHandler is the handler that processes the words that starts with qu
type QuHandler struct {
	Next Handler
}

// Translate is the method to translate the words begining with qu
func (q *QuHandler) Translate(w Word) Word {
	if ok, err := startQU(w.GetOriginal()); !ok && err == nil {
		log.Printf("QuHandler:[%s] not for me", w.GetOriginal())
		return q.Next.Translate(w)
	}

	w.QU()
	return w
}

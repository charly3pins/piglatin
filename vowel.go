package piglatin

import "log"

// VowelHandler is the handler that processes the words that begins with vowels
type VowelHandler struct {
	Next Handler
}

// Translate is the method to translate the words starting with vowels
func (v *VowelHandler) Translate(w Word) Word {
	if ok, err := startVowel(w.GetOriginal()); !ok && err == nil {
		log.Printf("VowelHandler:[%s] not for me", w.GetOriginal())
		return v.Next.Translate(w)
	}

	w.Vowel()
	return w
}

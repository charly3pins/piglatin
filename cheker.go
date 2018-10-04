package piglatin

import "regexp"

func startVowel(word string) (bool, error) {
	return regexp.MatchString("^[aeiouAEIOU]", word)
}

func startQU(word string) (bool, error) {
	return regexp.MatchString("^qu", word)
}

func isTitle(word string) (bool, error) {
	return regexp.MatchString("^[A-Z]", word)
}

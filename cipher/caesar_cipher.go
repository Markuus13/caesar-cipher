package cipher

import "strings"

func CaesarCipher(message string, cipherRule int) string {
	result := ""

	for _, letter := range message {
		result += getAlphabet()[getAlphabetIndex(letter, cipherRule)]
	}

	return result
}

func getAlphabetIndex(letter rune, cipherRule int) uint {
	index := getAlphabetIndexer()[strings.ToUpper(string(letter))] + cipherRule
	alphabetLength := len(getAlphabet())

	if index >= alphabetLength {
		return uint(index % alphabetLength)
	}

	return uint(index)
}

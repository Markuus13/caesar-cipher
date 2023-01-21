package cipher

import "strings"

func CaesarCipher(message string, cipherRule int) string {
	result := ""

	for _, letter := range message {
		result += getAlphabet()[getAlphabetIndexer()[strings.ToUpper(string(letter))]+cipherRule]
	}

	return result
}

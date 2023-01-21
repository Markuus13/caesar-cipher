package cipher_test

import (
	"caesar-cipher/cipher"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CaesarCipher", func() {

	When("input message is empty", func() {
		It("returns empty string no matter the cipher rule value", func() {
			inputMessage := ""

			encryptedMessage := cipher.CaesarCipher(inputMessage, 999)

			Expect(encryptedMessage).To(BeEmpty())
		})
	})

	When("input message is has a single word", func() {
		It("returns a string encrypted by given cipher rule", func() {
			inputMessage := "abcde"

			encryptedMessage := cipher.CaesarCipher(inputMessage, 3)

			Expect(encryptedMessage).To(Equal("DEFGH"))
		})
	})
})

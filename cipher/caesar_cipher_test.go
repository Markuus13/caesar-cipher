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

	DescribeTable("When input message is has a single word",
		func(inputMessage string, cipherRule int, output string) {
			encryptedMessage := cipher.CaesarCipher(inputMessage, cipherRule)

			Expect(encryptedMessage).To(Equal(output))
		},
		Entry("when cipher rule does not require go around the alphabet", "abcde", 3, "DEFGH"),
		Entry("when cipher rule requires going around the alphabet", "vwxyz", 5, "ABCDE"),
		Entry("when cipher rule requires going around the alphabet any times", "vwxyz", 2600, "VWXYZ"),
	)
})

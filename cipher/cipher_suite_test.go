package cipher_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCipher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cipher Suite")
}

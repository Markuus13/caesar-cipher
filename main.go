package main

import (
	"caesar-cipher/cipher"
	"fmt"
)

// input "abcdefgh"
// 	output -> "defghijk"
// var input = "xyz"
// var input = "abc def xyz"
// var input = "ABC def GhI"
// var input = "ÁBC DÊf xYZ"

func main() {
	var input = "abcdefgh"
	var cipherRule = 3
	fmt.Println(cipher.CaesarCipher(input, cipherRule))
}

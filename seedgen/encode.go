package seedgen

import (
	"crypto/sha256"
	"strconv"
	"strings"
)

// checksumByte returns the checksum byte used at the end of the seed mnemonic
// encoding.  The "checksum" is the first byte of the double SHA256.
func checksumByte(data []byte) byte {
	intermediateHash := sha256.Sum256(data)
	return sha256.Sum256(intermediateHash[:])[0]
}

func LeftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

// EncodeBinarySlice encodes a seed as a binary list.
func EncodeBinarySlice(seed []byte) []string {
	words := make([]string, len(seed)+1) // Extra word for checksumByte
	for i, b := range seed {
		words[i] = strconv.FormatInt(int64(b), 2)
	}
	checksum := checksumByte(seed)
	words[len(words)-1] = strconv.FormatInt(int64(checksum), 2)
	return words
}

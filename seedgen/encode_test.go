package seedgen

import (
	"crypto/rand"
	"testing"
)

func TestEncodeBinarySlice(t *testing.T) {
	entropyArr := make([]byte, 32)
	rand.Read(entropyArr)
	result := EncodeBinarySlice(entropyArr)
	if len(result) != 33 {
		t.Errorf("Expected %v, got %v", 33, result)
	}
}

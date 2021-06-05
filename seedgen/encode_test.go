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

func TestLeftPad2Len(t *testing.T) {
	result := LeftPad2Len("1", "0", 8)
	if result != "00000001" {
		t.Errorf("Expected 00000001, got %s", result)
	}
}

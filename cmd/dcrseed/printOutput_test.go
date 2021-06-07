package main

import (
	"crypto/rand"
	"testing"
)

func TestPrintPGPWords(t *testing.T) {
	result := printPGPWords(alternatingWords)
	if result != 512 {
		t.Errorf("Expected %d, got %d", 512, result)
	}
}

func TestPrintSeed(t *testing.T) {
	entropyArr := make([]byte, 32)
	rand.Read(entropyArr)
	err := printDecredSeed(entropyArr)
	if err != nil {
		t.Errorf("Unexpected error %w", err)
	}

}
func TestMain(t *testing.T) {
	main()
}

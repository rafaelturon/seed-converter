package seedgen

import "testing"

func TestGenerateSeed(t *testing.T) {
	result, _ := GenerateSeed(32)
	if len(result) != 32 {
		t.Errorf("Expected %v, got %v", 33, result)
	}
}

func TestGenerateRandomSeed(t *testing.T) {
	result, _ := GenerateRandomSeed(32)
	if len(result) != 32 {
		t.Errorf("Expected %v, got %v", 33, result)
	}
}

func TestGenerateDiceEntropySeed(t *testing.T) {
	result, _ := GenerateDiceEntropySeed("6543213223453321316456543212345666555123442123453321316456543212345666555123442123453321316456543212345666555123442123453321316456543212345666555123442123453321316")
	bb := uint16(result.RawData[0])
	if bb != 183 {
		t.Errorf("Expected %v, got %v", 183, result)
	}
}

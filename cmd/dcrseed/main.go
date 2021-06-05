package main

import (
	"fmt"

	"github.com/rafaelturon/seed-converter/seedgen"
)

func main() {
	printPGPWords(alternatingWords)
	seed, err := seedgen.GenerateRandomSeed(seedgen.RecommendedSeedLen)
	printSeed(seed)

	entropyArr, err := seedgen.GenerateDiceEntropySeed("6543213223453321316456543212345666555123442123453321316456543212345666555123442123453321316456543212345666555123442123453321316456543212345666555123442123453321316")
	printSeed(entropyArr)

	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return
	}
}

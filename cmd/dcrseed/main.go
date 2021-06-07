package main

import (
	"fmt"

	"github.com/rafaelturon/seed-converter/seedgen"
)

func main() {
	printPGPWords(alternatingWords)

	seed, _ := seedgen.GenerateRandomSeed(seedgen.RecommendedSeedLen)
	printBitcoinSeed(seed)
	printDecredSeed(seed)

	entropyArr, err := seedgen.GenerateDiceEntropySeed("6543213223453321316456543212345666555123442123453321316456543212345666555123442123453321316456543212345666555123442123453321316456543212345666555123442123453321316")
	printDecredSeed(entropyArr)

	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return
	}
}

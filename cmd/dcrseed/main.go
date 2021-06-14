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
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return
	}
	fmt.Printf("Entropy score     (0-4): %d\nEstimated entropy (bit): %f\nEstimated time to crack: %s\n\n",
		entropyArr.Score,
		entropyArr.Entropy,
		entropyArr.CrackTimeDisplay,
	)
	fmt.Println(entropyArr.Binary.Entropy)
	fmt.Println(entropyArr.Binary.Spaced)
	fmt.Println(entropyArr.Binary.Trimmed)

	printDecredSeed(entropyArr.RawData)
}

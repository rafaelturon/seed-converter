package main

import (
	"fmt"
	"strconv"
	"strings"

	"decred.org/dcrwallet/v2/walletseed"
	"github.com/rafaelturon/seed-converter/seedgen"
)

func printPGPWords(pgpWords string) int {
	var wordList = strings.Split(pgpWords, "\n")

	idx := 0
	hex := -1
	var wordType = ""
	for _, w := range wordList {
		wordType = "Odd"
		if idx%2 == 0 {
			wordType = "Even"
			hex++
		}
		hexa := seedgen.LeftPad2Len(fmt.Sprintf("%x", hex), "0", 2)
		binary := seedgen.LeftPad2Len(strconv.FormatInt(int64(hex), 2), "0", 8)
		fmt.Printf("[%s] %s : %s (%s)\n", hexa, binary, w, wordType)
		idx++
	}
	return idx
}

func printSeed(seed []byte) error {
	fmt.Println("(1) Your wallet generation seed is:")
	seedStr := walletseed.EncodeMnemonic(seed)
	fmt.Println(seedStr)
	decodedSeed, err := walletseed.DecodeUserInput(seedStr)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return err
	}

	fmt.Println("(2) This is the list:")
	seedStrSplit := walletseed.EncodeMnemonicSlice(decodedSeed)
	seedStrBinary := seedgen.EncodeBinarySlice(decodedSeed)
	for i := 0; i < int(seedgen.RecommendedSeedLen)+1; i++ {
		fmt.Printf("%v ", seedStrSplit[i])
		if seedStrHex, err := strconv.ParseUint(seedStrBinary[i], 2, 64); err == nil {
			fmt.Printf("%x ", seedStrHex)
		}
		if i%2 == 0 {
			fmt.Printf("Even ")
		} else {
			fmt.Printf("Odd ")
		}
		fmt.Println(seedgen.LeftPad2Len(seedStrBinary[i], "0", 8))

		if (i+1)%6 == 0 {
			fmt.Printf("\n")
		}
	}
	return nil
}

package main

import (
	"fmt"
	"strconv"
	"strings"

	"decred.org/dcrwallet/v2/walletseed"
	"github.com/rafaelturon/seed-converter/seedgen"
	"github.com/tyler-smith/go-bip39"
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

func printBitcoinSeed(seed []byte) error {
	fmt.Println("(1) Your wallet generation seed is:")
	seedStr, err := bip39.NewMnemonic(seed)
	fmt.Println(seedStr)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return err
	}

	fmt.Println("(2) This is the list:")
	seedStrSplit := strings.Fields(seedStr)
	seedStrBinary := seedgen.EncodeBinarySlice(seed)
	for i := 0; i < len(seedStrSplit); i++ {
		fmt.Printf(seedgen.LeftPad2Len(seedStrBinary[i], "0", 12))
		fmt.Printf(" %v\n", seedStrSplit[i])
	}
	return nil
}

func printDecredSeed(seed []byte) error {
	fmt.Println("(1) Your wallet generation seed is:")
	seedStr := walletseed.EncodeMnemonic(seed)
	fmt.Println(seedStr)

	fmt.Println("(2) This is the list:")
	seedStrSplit := walletseed.EncodeMnemonicSlice(seed)
	seedStrBinary := seedgen.EncodeBinarySlice(seed)
	for i := 0; i < int(seedgen.RecommendedSeedLen)+1; i++ {
		fmt.Printf(seedgen.LeftPad2Len(seedStrBinary[i], "0", 8))
		if seedStrHex, err := strconv.ParseUint(seedStrBinary[i], 2, 64); err == nil {
			fmt.Printf(" %x ", seedStrHex)
		}
		fmt.Printf("%v ", seedStrSplit[i])
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}

		if (i+1)%6 == 0 {
			fmt.Printf("\n")
		}
	}
	return nil
}

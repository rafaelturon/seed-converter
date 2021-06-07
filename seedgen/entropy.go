package seedgen

import (
	"crypto/rand"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"decred.org/dcrwallet/v2/errors"
	"github.com/nbutton23/zxcvbn-go"
)

const (
	// RecommendedSeedLen is the recommended length in bytes for a seed
	// to a master node.
	RecommendedSeedLen = 32 // 256 bits

	// HardenedKeyStart is the index at which a hardened key starts.  Each
	// extended key has 2^31 normal child keys and 2^31 hardened child keys.
	// Thus the range for normal child keys is [0, 2^31 - 1] and the range
	// for hardened child keys is [2^31, 2^32 - 1].
	HardenedKeyStart = 0x80000000 // 2^31

	// MinSeedBytes is the minimum number of bytes allowed for a seed to
	// a master node.
	MinSeedBytes = 16 // 128 bits

	// MaxSeedBytes is the maximum number of bytes allowed for a seed to
	// a master node.
	MaxSeedBytes = 64 // 512 bits

	// serializedKeyLen is the length of a serialized public or private
	// extended key.  It consists of 4 bytes version, 1 byte depth, 4 bytes
	// fingerprint, 4 bytes child number, 32 bytes chain code, and 33 bytes
	// public/private key data.
	serializedKeyLen = 4 + 1 + 4 + 4 + 32 + 33 // 78 bytes
)

var (
	// ErrInvalidSeedLen describes an error in which the provided seed or
	// seed length is not in the allowed range.
	ErrInvalidSeedLen = fmt.Errorf("seed length must be between %d and %d "+
		"bits", MinSeedBytes*8, MaxSeedBytes*8)
)

// GenerateSeed returns a cryptographically secure random seed that can be used
// as the input for the NewMaster function to generate a new master node.
//
// The length is in bytes and it must be between 16 and 64 (128 to 512 bits).
// The recommended length is 32 (256 bits) as defined by the RecommendedSeedLen
// constant.
func GenerateSeed(length uint8) ([]byte, error) {
	// Per [BIP32], the seed must be in range [MinSeedBytes, MaxSeedBytes].
	if length < MinSeedBytes || length > MaxSeedBytes {
		return nil, ErrInvalidSeedLen
	}

	buf := make([]byte, length)
	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// GenerateRandomSeed returns a new seed created from a cryptographically-secure
// random source.  If the seed size is unacceptable,
// hdkeychain.ErrInvalidSeedLen is returned.
func GenerateRandomSeed(size uint) ([]byte, error) {
	const op errors.Op = "walletseed.GenerateRandomSeed"
	if size >= uint(^uint8(0)) {
		return nil, errors.E(op, errors.Invalid, ErrInvalidSeedLen)
	}
	if size < MinSeedBytes || size > MaxSeedBytes {
		return nil, errors.E(op, errors.Invalid, ErrInvalidSeedLen)
	}
	seed, err := GenerateSeed(uint8(size))
	if err != nil {
		return nil, errors.E(op, err)
	}
	return seed, nil
}

func ConvertDiceRollsToBinaryStr(entropyStr string) (string, error) {
	const op errors.Op = "walletseed.GenerateEntropySeed"
	// log2(6) = 2.58496 bits per roll, with bias
	// 4 rolls give 2 bits each
	// 2 rolls give 1 bit each
	// Average (4*2 + 2*1) / 6 = 1.66 bits per roll without bias

	// Convert dice to base6 entropy (ie 1-6 to 0-5)
	// This is done by changing all 6s to 0s
	diceEntropy := strings.ReplaceAll(entropyStr, "6", "0")
	events := strings.Split(diceEntropy, "")

	//"base 6 (dice)": {
	//    "0": "00", // equivalent to 0 in base 6
	//    "1": "01",
	//    "2": "10",
	//    "3": "11",
	//    "4": "0",
	//    "5": "1",
	//}
	binary := [6]string{"00", "01", "10", "11", "0", "1"}
	for i := 0; i < len(events); i++ {
		entry, err := strconv.Atoi(events[i])
		events[i] = binary[entry]
		if err != nil {
			return "", errors.E(op, err)
		}
	}

	return strings.Join(events, ""), nil
}

func ConvertTrimmedBinary(trimmedBinaryStr string) ([]byte, error) {
	var entropyArr []byte
	runes := []rune(trimmedBinaryStr)
	for i := 0; i < len(trimmedBinaryStr)/8; i++ {
		byteAsBits := string(runes[i*8 : i*8+8])
		if entropyByte, err := strconv.ParseUint(byteAsBits, 2, 8); err == nil {
			entropyArr = append(entropyArr, byte(entropyByte))
		}
	}

	return entropyArr, nil
}

func GenerateDiceEntropySeed(entropyStr string) ([]byte, error) {
	binaryStr, err := ConvertDiceRollsToBinaryStr(entropyStr)
	if err != nil {
		fmt.Printf("Error converting dice rolls to binary: %s", err)
		return nil, err
	}

	pwdStrength := zxcvbn.PasswordStrength(binaryStr, nil)
	fmt.Printf("Entropy score     (0-4): %d\nEstimated entropy (bit): %f\nEstimated time to crack: %s\n\n",
		pwdStrength.Score,
		pwdStrength.Entropy,
		pwdStrength.CrackTimeDisplay,
	)

	numberOfBits := len(binaryStr)
	var wordCount float64 = math.Floor(float64(numberOfBits)/32) * 3
	var re = regexp.MustCompile(`.{1,11}`)
	spacedBinaryStr := strings.Join(re.FindAllString(binaryStr, -1), " ")
	fmt.Printf("Binary string entropy %s of size %d bits with %f words\n", binaryStr, numberOfBits, wordCount)
	fmt.Println("Spaced Binary", spacedBinaryStr)

	var bitsToUse int = int(float64(len(binaryStr))/32) * 32
	var start = int(len(binaryStr)) - bitsToUse
	var trimmedBinaryStr = binaryStr[start:]
	fmt.Printf("Trimmed binary string %s of size %d bits starting at %d position\n", trimmedBinaryStr, bitsToUse, start)

	entropyArr, err := ConvertTrimmedBinary(trimmedBinaryStr)
	if err != nil {
		fmt.Printf("Error converting trimmed string to binary: %s", err)
		return nil, err
	}

	return entropyArr, nil
}

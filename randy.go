package randy

import (
	"bytes"
	"crypto/rand"
	_ "embed"
	"fmt"
	"math/big"
	"strings"
)

//go:embed wordlist.txt
var words string
var wordlist []string
var wordlistLength *big.Int

func init() {
	wordlist = strings.Split(words, "\n")
	wordlistLength = big.NewInt(int64(len(wordlist) - 1))

	defaultAlphabetLength = big.NewInt(int64(len(defaultAlphabet) - 1))
}

func GenerateWordString(length int, sep string) string {
	return strings.Join(GenerateWords(length), sep)
}

func GenerateWords(length int) []string {
	var s []string

	for i := 0; i < length; i++ {
		idx, err := rand.Int(rand.Reader, wordlistLength)
		if err != nil {
			panic(fmt.Errorf("failed to generate random number: %w", err))
		}

		s = append(s, wordlist[idx.Int64()])
	}

	return s
}

var defaultAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
var defaultAlphabetLength *big.Int

func GenerateToken(length int) string {
	var b bytes.Buffer

	for i := 0; i < length; i++ {
		idx, err := rand.Int(rand.Reader, defaultAlphabetLength)
		if err != nil {
			panic(fmt.Errorf("failed to generate random number: %w", err))
		}

		b.WriteByte(defaultAlphabet[idx.Int64()])
	}

	return b.String()
}

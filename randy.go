package randy

import (
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
}

func GenerateString(words int, sep string) string {
	return strings.Join(gen(words), sep)
}

func Generate(words int) []string {
	return gen(words)
}

func gen(length int) []string {
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

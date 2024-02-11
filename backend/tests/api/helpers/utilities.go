package helpers

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

func generateRandomInt(max int64) int64 {
	randInt, _ := crand.Int(crand.Reader, big.NewInt(max))
	return randInt.Int64()
}

func generateRandomDBName() string {
	prefix := "sac_test_"
	letterBytes := "abcdefghijklmnopqrstuvwxyz"
	length := len(prefix) + 36
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = letterBytes[generateRandomInt(int64(len(letterBytes)))]
	}

	return fmt.Sprintf("%s%s", prefix, string(result))
}

func generateCasingPermutations(word string, currentPermutation string, index int, results *[]string) {
	if index == len(word) {
		*results = append(*results, currentPermutation)
		return
	}

	generateCasingPermutations(word, fmt.Sprintf("%s%s", currentPermutation, strings.ToLower(string(word[index]))), index+1, results)
	generateCasingPermutations(word, fmt.Sprintf("%s%s", currentPermutation, strings.ToUpper(string(word[index]))), index+1, results)
}

func AllCasingPermutations(word string) []string {
	results := make([]string, 0)
	generateCasingPermutations(word, "", 0, &results)
	return results
}

func StringToPointer(s string) *string {
	return &s
}

package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(password string) string {
	reversePassword := Reverse(password)
	hashForSalt := sha256.Sum256([]byte(reversePassword))
	salt := hex.EncodeToString(hashForSalt[:])

	hashForPassword := sha256.Sum256([]byte(salt + password))
	hash := hex.EncodeToString(hashForPassword[:])
	return hash

}

func Reverse(s string) string {
	rns := []rune(s)

	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

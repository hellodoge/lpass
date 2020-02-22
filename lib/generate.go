package lib

import (
	"crypto/sha256"
	"golang.org/x/crypto/pbkdf2"
)

func Generate(master, token []byte, iter, length int) []byte {
	var masterHash = sha256.Sum256(master)
	var tokenHash = sha256.Sum256(token)
	return pbkdf2.Key(masterHash[:], tokenHash[:], 
		iter, length, sha256.New)
}

package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"os"
)

var (
	hashAlgo hash.Hash
	hashSalt string
)

func init() {
	hashSalt = os.Getenv("HASH_SALT")
	hashAlgo = sha256.New()
}

func GetHash(s string) (hashedString string) {
	sToHash := s + hashSalt
	hashAlgo.Write([]byte(sToHash))
	return hex.EncodeToString(hashAlgo.Sum(nil))
}

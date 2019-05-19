package chain

import (
	"crypto/sha256"
)

type Hash [32]byte

// DoubleHash calculate the double hash value
// of the data.
func DoubleHash(data []byte) Hash {
	digest := sha256.Sum256(data)
	return sha256.Sum256(digest[:])
}
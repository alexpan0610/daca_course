package block

import (
	"github.com/elastos/cource/chain"
	"time"
)

type Header struct {
	Previous   chain.Hash
	Timestamp  time.Time
	MerkleRoot chain.Hash
	Height     uint32
	Nonce      uint32
	Bits       uint32
}



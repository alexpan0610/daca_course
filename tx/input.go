package tx

import (
	"github.com/elastos/cource/chain"
)

type Input struct {
	Previous chain.Hash
	Index    uint16
}

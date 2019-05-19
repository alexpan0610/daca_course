package chain

import "encoding/hex"

type Address [20]byte

func (a *Address) String() string {
	return hex.EncodeToString(a[:])
}

package node

import (
	amino "github.com/evdatsion/go-amino"
	cryptoAmino "github.com/evdatsion/aphelion-dpos-bft/crypto/encoding/amino"
)

var cdc = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
}

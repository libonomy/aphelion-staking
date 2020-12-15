package p2p

import (
	amino "github.com/evdatsion/go-amino"
	cryptoAmino "github.com/evdatsion/tendermint/crypto/encoding/amino"
)

var cdc = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
}

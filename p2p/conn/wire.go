package conn

import (
	amino "github.com/evdatsion/go-amino"
	cryptoAmino "github.com/libonomy/aphelion-staking/crypto/encoding/amino"
)

var cdc *amino.Codec = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
	RegisterPacket(cdc)
}

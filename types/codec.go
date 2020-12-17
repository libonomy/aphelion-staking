package types

import (
	amino "github.com/evdatsion/go-amino"
	cryptoAmino "github.com/libonomy/aphelion-staking/crypto/encoding/amino"
)

var cdc = amino.NewCodec()

func init() {
	RegisterBlockAmino(cdc)
}

func RegisterBlockAmino(cdc *amino.Codec) {
	cryptoAmino.RegisterAmino(cdc)
	RegisterEvidences(cdc)
}

// GetCodec returns a codec used by the package. For testing purposes only.
func GetCodec() *amino.Codec {
	return cdc
}

// For testing purposes only
func RegisterMockEvidencesGlobal() {
	RegisterMockEvidences(cdc)
}

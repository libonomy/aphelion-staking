package evidence

import (
	amino "github.com/evdatsion/go-amino"
	cryptoAmino "github.com/libonomy/aphelion-staking/crypto/encoding/amino"
	"github.com/libonomy/aphelion-staking/types"
)

var cdc = amino.NewCodec()

func init() {
	RegisterEvidenceMessages(cdc)
	cryptoAmino.RegisterAmino(cdc)
	types.RegisterEvidences(cdc)
}

// For testing purposes only
func RegisterMockEvidences() {
	types.RegisterMockEvidences(cdc)
}

package evidence

import (
	amino "github.com/evdatsion/go-amino"
	cryptoAmino "github.com/evdatsion/tendermint/crypto/encoding/amino"
	"github.com/evdatsion/tendermint/types"
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

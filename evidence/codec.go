package evidence

import (
	amino "github.com/evdatsion/go-amino"
	cryptoAmino "github.com/evdatsion/aphelion-dpos-bft/crypto/encoding/amino"
	"github.com/evdatsion/aphelion-dpos-bft/types"
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

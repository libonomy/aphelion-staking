package v0

import (
	amino "github.com/evdatsion/go-amino"
	"github.com/evdatsion/tendermint/types"
)

var cdc = amino.NewCodec()

func init() {
	RegisterBlockchainMessages(cdc)
	types.RegisterBlockAmino(cdc)
}

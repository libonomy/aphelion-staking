package v1

import (
	amino "github.com/evdatsion/go-amino"
	"github.com/libonomy/aphelion-staking/types"
)

var cdc = amino.NewCodec()

func init() {
	RegisterBlockchainMessages(cdc)
	types.RegisterBlockAmino(cdc)
}

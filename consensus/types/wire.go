package types

import (
	amino "github.com/evdatsion/go-amino"
	"github.com/evdatsion/aphelion-dpos-bft/types"
)

var cdc = amino.NewCodec()

func init() {
	types.RegisterBlockAmino(cdc)
}

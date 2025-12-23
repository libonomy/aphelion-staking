package client

import (
	amino "github.com/libonomy/go-amino"
	"github.com/libonomy/aphelion-staking/types"
)

var cdc = amino.NewCodec()

func init() {
	types.RegisterEvidences(cdc)
}

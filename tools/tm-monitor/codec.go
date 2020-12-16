package main

import (
	amino "github.com/evdatsion/go-amino"
	ctypes "github.com/evdatsion/aphelion-dpos-bft/rpc/core/types"
)

var cdc = amino.NewCodec()

func init() {
	ctypes.RegisterAmino(cdc)
}

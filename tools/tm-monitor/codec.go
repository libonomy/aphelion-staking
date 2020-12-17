package main

import (
	amino "github.com/evdatsion/go-amino"
	ctypes "github.com/libonomy/aphelion-staking/rpc/core/types"
)

var cdc = amino.NewCodec()

func init() {
	ctypes.RegisterAmino(cdc)
}

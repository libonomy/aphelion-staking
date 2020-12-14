package monitor

import (
	amino "github.com/evdatsion/go-amino"
	ctypes "github.com/evdatsion/tendermint/rpc/core/types"
)

var cdc = amino.NewCodec()

func init() {
	ctypes.RegisterAmino(cdc)
}

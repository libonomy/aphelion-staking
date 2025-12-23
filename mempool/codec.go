package mempool

import (
	amino "github.com/libonomy/go-amino"
)

var cdc = amino.NewCodec()

func init() {
	RegisterMempoolMessages(cdc)
}

package mempool

import (
	amino "github.com/evdatsion/go-amino"
)

var cdc = amino.NewCodec()

func init() {
	RegisterMempoolMessages(cdc)
}

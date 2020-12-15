package pex

import (
	amino "github.com/evdatsion/go-amino"
)

var cdc *amino.Codec = amino.NewCodec()

func init() {
	RegisterPexMessage(cdc)
}

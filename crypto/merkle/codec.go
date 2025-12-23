package merkle

import (
	amino "github.com/libonomy/go-amino"
)

var cdc *amino.Codec

func init() {
	cdc = amino.NewCodec()
	cdc.Seal()
}

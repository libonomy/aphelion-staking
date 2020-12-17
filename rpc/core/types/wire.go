package core_types

import (
	amino "github.com/evdatsion/go-amino"
	"github.com/libonomy/aphelion-staking/types"
)

func RegisterAmino(cdc *amino.Codec) {
	types.RegisterEventDatas(cdc)
	types.RegisterBlockAmino(cdc)
}

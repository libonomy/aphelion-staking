package core_types

import (
	amino "github.com/evdatsion/go-amino"
	"github.com/evdatsion/aphelion-dpos-bft/types"
)

func RegisterAmino(cdc *amino.Codec) {
	types.RegisterEventDatas(cdc)
	types.RegisterBlockAmino(cdc)
}

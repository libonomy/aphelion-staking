package core

import (
	ctypes "github.com/evdatsion/tendermint/rpc/core/types"
	rpctypes "github.com/evdatsion/tendermint/rpc/lib/types"
	"github.com/evdatsion/tendermint/types"
)

// BroadcastEvidence broadcasts evidence of the misbehavior.
// More: https://tendermint.com/rpc/#/Info/broadcast_evidence
func BroadcastEvidence(ctx *rpctypes.Context, ev types.Evidence) (*ctypes.ResultBroadcastEvidence, error) {
	err := evidencePool.AddEvidence(ev)
	if err != nil {
		return nil, err
	}
	return &ctypes.ResultBroadcastEvidence{Hash: ev.Hash()}, nil
}

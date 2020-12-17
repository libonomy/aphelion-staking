package core

import (
	ctypes "github.com/libonomy/aphelion-staking/rpc/core/types"
	rpctypes "github.com/libonomy/aphelion-staking/rpc/lib/types"
	"github.com/libonomy/aphelion-staking/types"
)

// BroadcastEvidence broadcasts evidence of the misbehavior.
// More: https://aphelion.com/rpc/#/Info/broadcast_evidence
func BroadcastEvidence(ctx *rpctypes.Context, ev types.Evidence) (*ctypes.ResultBroadcastEvidence, error) {
	err := evidencePool.AddEvidence(ev)
	if err != nil {
		return nil, err
	}
	return &ctypes.ResultBroadcastEvidence{Hash: ev.Hash()}, nil
}

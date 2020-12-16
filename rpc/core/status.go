package core

import (
	"bytes"
	"time"

	cmn "github.com/evdatsion/aphelion-dpos-bft/libs/common"
	"github.com/evdatsion/aphelion-dpos-bft/p2p"
	ctypes "github.com/evdatsion/aphelion-dpos-bft/rpc/core/types"
	rpctypes "github.com/evdatsion/aphelion-dpos-bft/rpc/lib/types"
	sm "github.com/evdatsion/aphelion-dpos-bft/state"
	"github.com/evdatsion/aphelion-dpos-bft/types"
)

// Status returns Tendermint status including node info, pubkey, latest block
// hash, app hash, block height and time.
// More: https://aphelion.com/rpc/#/Info/status
func Status(ctx *rpctypes.Context) (*ctypes.ResultStatus, error) {
	var latestHeight int64
	if consensusReactor.FastSync() {
		latestHeight = blockStore.Height()
	} else {
		latestHeight = consensusState.GetLastHeight()
	}
	var (
		latestBlockMeta     *types.BlockMeta
		latestBlockHash     cmn.HexBytes
		latestAppHash       cmn.HexBytes
		latestBlockTimeNano int64
	)
	if latestHeight != 0 {
		latestBlockMeta = blockStore.LoadBlockMeta(latestHeight)
		latestBlockHash = latestBlockMeta.BlockID.Hash
		latestAppHash = latestBlockMeta.Header.AppHash
		latestBlockTimeNano = latestBlockMeta.Header.Time.UnixNano()
	}

	latestBlockTime := time.Unix(0, latestBlockTimeNano)

	var votingPower int64
	if val := validatorAtHeight(latestHeight); val != nil {
		votingPower = val.VotingPower
	}

	result := &ctypes.ResultStatus{
		NodeInfo: p2pTransport.NodeInfo().(p2p.DefaultNodeInfo),
		SyncInfo: ctypes.SyncInfo{
			LatestBlockHash:   latestBlockHash,
			LatestAppHash:     latestAppHash,
			LatestBlockHeight: latestHeight,
			LatestBlockTime:   latestBlockTime,
			CatchingUp:        consensusReactor.FastSync(),
		},
		ValidatorInfo: ctypes.ValidatorInfo{
			Address:     pubKey.Address(),
			PubKey:      pubKey,
			VotingPower: votingPower,
		},
	}

	return result, nil
}

func validatorAtHeight(h int64) *types.Validator {
	privValAddress := pubKey.Address()

	// If we're still at height h, search in the current validator set.
	lastBlockHeight, vals := consensusState.GetValidators()
	if lastBlockHeight == h {
		for _, val := range vals {
			if bytes.Equal(val.Address, privValAddress) {
				return val
			}
		}
	}

	// If we've moved to the next height, retrieve the validator set from DB.
	if lastBlockHeight > h {
		vals, err := sm.LoadValidators(stateDB, h)
		if err != nil {
			return nil // should not happen
		}
		_, val := vals.GetByAddress(privValAddress)
		return val
	}

	return nil
}

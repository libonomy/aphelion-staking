package core

import (
	ctypes "github.com/libonomy/aphelion-staking/rpc/core/types"
	rpctypes "github.com/libonomy/aphelion-staking/rpc/lib/types"
)

// Health gets node health. Returns empty result (200 OK) on success, no
// response - in case of an error.
// More: https://aphelion.com/rpc/#/Info/health
func Health(ctx *rpctypes.Context) (*ctypes.ResultHealth, error) {
	return &ctypes.ResultHealth{}, nil
}

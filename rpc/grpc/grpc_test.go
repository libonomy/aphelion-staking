package core_grpc_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/libonomy/aphelion-staking/abci/example/kvstore"
	core_grpc "github.com/libonomy/aphelion-staking/rpc/grpc"
	rpctest "github.com/libonomy/aphelion-staking/rpc/test"
)

func TestMain(m *testing.M) {
	// start a aphelion node in the background to test against
	app := kvstore.NewKVStoreApplication()
	node := rpctest.StartAphelion(app)

	code := m.Run()

	// and shut down proper at the end
	rpctest.StopAphelion(node)
	os.Exit(code)
}

func TestBroadcastTx(t *testing.T) {
	res, err := rpctest.GetGRPCClient().BroadcastTx(
		context.Background(),
		&core_grpc.RequestBroadcastTx{Tx: []byte("this is a tx")},
	)
	require.NoError(t, err)
	require.EqualValues(t, 0, res.CheckTx.Code)
	require.EqualValues(t, 0, res.DeliverTx.Code)
}

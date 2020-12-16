package client_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/evdatsion/aphelion-dpos-bft/abci/example/kvstore"
	nm "github.com/evdatsion/aphelion-dpos-bft/node"
	rpctest "github.com/evdatsion/aphelion-dpos-bft/rpc/test"
)

var node *nm.Node

func TestMain(m *testing.M) {
	// start a aphelion node (and kvstore) in the background to test against
	dir, err := ioutil.TempDir("/tmp", "rpc-client-test")
	if err != nil {
		panic(err)
	}
	app := kvstore.NewPersistentKVStoreApplication(dir)
	node = rpctest.StartAphelion(app)

	code := m.Run()

	// and shut down proper at the end
	rpctest.StopAphelion(node)
	os.Exit(code)
}

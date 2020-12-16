package state_test

import (
	"os"
	"testing"

	"github.com/evdatsion/aphelion-dpos-bft/types"
)

func TestMain(m *testing.M) {
	types.RegisterMockEvidencesGlobal()
	os.Exit(m.Run())
}

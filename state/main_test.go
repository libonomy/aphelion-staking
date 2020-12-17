package state_test

import (
	"os"
	"testing"

	"github.com/libonomy/aphelion-staking/types"
)

func TestMain(m *testing.M) {
	types.RegisterMockEvidencesGlobal()
	os.Exit(m.Run())
}

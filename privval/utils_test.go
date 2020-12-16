package privval

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	cmn "github.com/evdatsion/aphelion-dpos-bft/libs/common"
)

func TestIsConnTimeoutForNonTimeoutErrors(t *testing.T) {
	assert.False(t, IsConnTimeout(cmn.ErrorWrap(ErrDialRetryMax, "max retries exceeded")))
	assert.False(t, IsConnTimeout(fmt.Errorf("completely irrelevant error")))
}

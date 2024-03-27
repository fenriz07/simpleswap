package keeper_test

import (
	"testing"

	"cosmossdk.io/log"
	fakeKeeper "github.com/fenriz07/simpleswap/testutil/keeper"
	"github.com/stretchr/testify/assert"
)

func TestKeeper_Logger(t *testing.T) {
	want := log.NewNopLogger()

	k, _ := fakeKeeper.SimpleswapKeeper(t)
	got := k.Logger()

	assert.Equal(t, want, got)
}

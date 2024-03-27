package keeper_test

import (
	"context"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	fakeKeeper "github.com/fenriz07/simpleswap/testutil/keeper"
	"github.com/fenriz07/simpleswap/x/simpleswap/keeper"
	simpleswap "github.com/fenriz07/simpleswap/x/simpleswap/module"
	"github.com/fenriz07/simpleswap/x/simpleswap/types"
	"github.com/stretchr/testify/assert"
)

type claimLiquidityTest struct {
	name string
	msg  types.MsgClaimLiquidity
	want error
}

var accountPool string

func setupMsgServerTest(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := fakeKeeper.SimpleswapKeeper(t)
	simpleswap.InitGenesis(ctx, k, *types.DefaultGenesis())

	accountPool = sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()).String()

	k.SetPool(ctx, types.Pool{
		Index:   "10",
		Account: accountPool,
		Coin:    "usdt",
		Amount:  "10",
		Fees:    "4",
	})

	k.SetPool(ctx, types.Pool{
		Index:   "11",
		Account: accountPool,
		Coin:    "usdt",
		Amount:  "xd",
		Fees:    "yd",
	})

	return keeper.NewMsgServerImpl(k), k, ctx
}

func TestMsgServer_ClaimLiquidity(t *testing.T) {

	msgServer, _, ctx := setupMsgServerTest(t)

	t.Parallel()

	creator := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()).String()

	tableTest := []claimLiquidityTest{
		{
			name: "invalid creator",
			msg: types.MsgClaimLiquidity{
				Creator: "invalid",
			},
			want: types.ErrInvalidAccount,
		},
		{
			name: "pool not found",
			msg: types.MsgClaimLiquidity{
				Creator:             creator,
				ProviderLiquidityId: "1",
			},
			want: types.ErrLiquidityProviderNotFound,
		},
		{
			name: "invalid account",
			msg: types.MsgClaimLiquidity{
				Creator:             creator,
				ProviderLiquidityId: "10",
			},
			want: types.ErrInvalidAccount,
		},
		{
			name: "err parse amount",
			msg: types.MsgClaimLiquidity{
				Creator:             accountPool,
				ProviderLiquidityId: "11",
			},
			want: types.ErrParsingStringToInt,
		},
		{
			name: "err parse fees",
			msg: types.MsgClaimLiquidity{
				Creator:             accountPool,
				ProviderLiquidityId: "11",
			},
			want: types.ErrParsingStringToInt,
		},
		{
			name: "err parse fees",
			msg: types.MsgClaimLiquidity{
				Creator:             accountPool,
				ProviderLiquidityId: "11",
			},
			want: types.ErrParsingStringToInt,
		},
		{
			name: "err sending coin",
			msg: types.MsgClaimLiquidity{
				Creator:             accountPool,
				ProviderLiquidityId: "10",
			},
			want: nil,
		},
	}

	for _, test := range tableTest {
		t.Run(test.name, func(t *testing.T) {
			_, got := msgServer.ClaimLiquidity(ctx, &test.msg)

			assert.ErrorIs(t, test.want, got)
		})
	}

}

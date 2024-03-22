package keeper

import (
	"context"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fenriz07/simpleswap/x/simpleswap/types"
)

func (k msgServer) ProvideLiquidity(goCtx context.Context, msg *types.MsgProvideLiquidity) (*types.MsgProvideLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	//k.logger.Error("probando", "epa")

	from, _ := sdk.AccAddressFromBech32(msg.Creator)
	to, _ := sdk.AccAddressFromBech32("cosmos13q5hctv0djtjqae4mwdyjjsumgzp2dy9m7zekr")
	//sdk.NewCoins(sdk.Coin{Denom: "usdc", Amount: math.NewInt(2)}
	k.bank.SendCoins(ctx, from, to, sdk.Coins{sdk.NewCoin("usdc", math.NewInt(2))})
	return &types.MsgProvideLiquidityResponse{
		ProvideLiquidityResponse: "epa",
	}, nil
}

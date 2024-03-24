package keeper

import (
	"context"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fenriz07/simpleswap/x/simpleswap/shared/constants"
	"github.com/fenriz07/simpleswap/x/simpleswap/types"
)

func (k msgServer) ClaimLiquidity(goCtx context.Context, msg *types.MsgClaimLiquidity) (*types.MsgClaimLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	client, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidAccount, "(%s)", client.String())
	}

	bank, err := sdk.AccAddressFromBech32(constants.BANK)

	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidAccount, "(%s)", bank.String())
	}

	liquidityProvider, found := k.Keeper.GetPool(ctx, msg.ProviderLiquidityId)

	if !found {
		return nil, errorsmod.Wrapf(types.ErrLiquidityProviderNotFound, "(%s)", msg.ProviderLiquidityId)
	}

	if liquidityProvider.Account != msg.Creator {
		return nil, errorsmod.Wrapf(types.ErrInvalidAccount, "invalid account for this liquidity provider : %s", msg.Creator)
	}

	liquidity, err := strconv.ParseInt(liquidityProvider.Amount, 10, 64)

	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrParsingStringToInt, "%s", err.Error())
	}

	fees, err := strconv.ParseInt(liquidityProvider.Fees, 10, 64)

	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrParsingStringToInt, "%s", err.Error())
	}

	amountToSend := liquidity + fees
	coin := sdk.NewCoin(liquidityProvider.Coin, math.NewInt(amountToSend))

	err = k.Keeper.bank.SendCoins(ctx, bank, client, sdk.NewCoins(coin))

	if err != nil {
		k.Keeper.logger.Error("There is a problem to claim liquidity and fees", "claim-liquidity")
		return nil, errorsmod.Wrapf(types.ErrSendCoin, "%s", err.Error())
	}

	k.Keeper.RemovePool(ctx, msg.ProviderLiquidityId)

	return &types.MsgClaimLiquidityResponse{}, nil
}

package keeper

import (
	"context"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/fenriz07/simpleswap/x/simpleswap/types"
)

func (k msgServer) ProvideLiquidity(goCtx context.Context, msg *types.MsgProvideLiquidity) (*types.MsgProvideLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	stableCoin, found := k.Keeper.GetStableCoinsWhiteList(ctx, msg.StableCoinId)

	if !found {
		k.logger.Warn(types.ErrStableCoinNotFound.Error(), "provideLiquidity")
		return nil, errorsmod.Wrapf(types.ErrStableCoinNotFound, "(%s)", msg.StableCoinId)
	}

	if !stableCoin.Available {
		return nil, errorsmod.Wrapf(types.ErrStableIsNotAvailable, "(%s)", stableCoin.GetCoin())
	}

	from, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidAccount, "(%s)", from.String())
	}

	//BANK ACCOUNT
	to, err := sdk.AccAddressFromBech32("cosmos189z79vlskxjm4n54va5954xlh02ktca6djmct4")

	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidAccount, "(%s)", to.String())
	}

	amount, err := strconv.ParseInt(msg.Amount, 10, 64)

	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrAmoutHasToBeAnInt, "invalid amount value (%s)", err)
	}

	k.bank.SendCoins(ctx, from, to, sdk.Coins{sdk.NewCoin(stableCoin.Coin, math.NewInt(amount))})

	return &types.MsgProvideLiquidityResponse{}, nil
}

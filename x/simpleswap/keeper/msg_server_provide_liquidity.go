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

const (
	shareTokenName = "sharetoken"
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

	provider, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidAccount, "(%s)", provider.String())
	}

	bank, err := sdk.AccAddressFromBech32(constants.BANK)

	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidAccount, "(%s)", bank.String())
	}

	amount, err := strconv.ParseInt(msg.Amount, 10, 64)

	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrAmoutHasToBeAnInt, "invalid amount value (%s)", err)
	}

	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}

	err = k.bank.SendCoins(ctx, provider, bank, sdk.Coins{sdk.NewCoin(stableCoin.Coin, math.NewInt(amount))})

	if err != nil {
		errorsmod.Wrapf(types.ErrSendCoin, "(%s)", err)
	}

	newIndex := strconv.FormatUint(systemInfo.NextId, 10)

	k.Keeper.SetPool(ctx, types.Pool{
		Index:   newIndex,
		Account: msg.Creator,
		Coin:    stableCoin.Coin,
		Amount:  strconv.Itoa(int(amount)),
		Fees:    "0",
	})

	k.Keeper.GetParams(ctx)

	systemInfo.NextId++
	k.Keeper.SetSystemInfo(ctx, systemInfo)

	var shareTokenPriceByStableCoin int64 = 100000 // 0.10
	shareTokenAmount := (amount * shareTokenPriceByStableCoin) / 10000000

	coin := sdk.Coins{sdk.NewCoin(shareTokenName, math.NewInt(shareTokenAmount))}

	err = k.bank.SendCoins(ctx, bank, provider, coin)
	if err != nil {
		errorsmod.Wrapf(types.ErrSendCoin, "(%s)", err)
	}

	return &types.MsgProvideLiquidityResponse{}, nil
}

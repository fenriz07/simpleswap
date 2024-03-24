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

func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	stableCoin, found := k.Keeper.GetStableCoinsWhiteList(ctx, msg.Denom)

	if !found {
		k.logger.Warn(types.ErrStableCoinNotFound.Error(), "provideLiquidity")
		return nil, errorsmod.Wrapf(types.ErrStableCoinNotFound, "(%s)", msg.Denom)
	}

	if !stableCoin.Available {
		return nil, errorsmod.Wrapf(types.ErrStableIsNotAvailable, "(%s)", stableCoin.GetCoin())
	}

	client, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidAccount, "(%s)", client.String())
	}

	bank, err := sdk.AccAddressFromBech32(constants.BANK)

	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidAccount, "(%s)", bank.String())
	}

	ethCoin := msg.GetAmount()[0]
	var ethPriceUsd int64 = 3317000000 //eth price expressed in USD
	usdToTranferWithOutFee := ethPriceUsd * ethCoin.Amount.Int64() / 1000000
	fee := (usdToTranferWithOutFee * 3000) / 1000000 //3000 = 0.3%
	usdToTranfer := usdToTranferWithOutFee - fee
	usdStableCoin := sdk.NewCoin(stableCoin.Coin, math.NewInt(usdToTranfer))

	//Validations
	hasBalanceInBank := k.bank.HasBalance(ctx, bank, usdStableCoin)
	if !hasBalanceInBank {
		return nil, errorsmod.Wrapf(types.ErrDoesNotHasBalance, "(%s %s)", bank.String(), usdStableCoin)
	}

	hasBalanceInClient := k.bank.HasBalance(ctx, client, ethCoin)

	if !hasBalanceInClient {
		return nil, errorsmod.Wrapf(types.ErrDoesNotHasBalance, "(%s %s)", client.String(), ethCoin)
	}

	err = k.bank.SendCoins(ctx, bank, client, sdk.Coins{usdStableCoin})

	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrSendCoin, "(%s)", err)
	}

	err = k.bank.SendCoins(ctx, client, bank, sdk.NewCoins(ethCoin))

	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrSendCoin, "(%s)", err)
	}

	sendSwapToLiquidityProviders(ctx, fee, k, stableCoin, bank)

	return &types.MsgSwapResponse{}, nil
}

func sendSwapToLiquidityProviders(ctx context.Context, fee int64, k msgServer, stableCoin types.StableCoinsWhiteList, bank sdk.AccAddress) error {

	liquidityProviders := k.GetAllPool(ctx)
	var poolBalance int64
	var pool []types.Pool

	for _, liquidityProvider := range liquidityProviders {

		if liquidityProvider.Coin != stableCoin.Coin {
			continue
		}

		liquidityProviderAmount, _ := strconv.ParseInt(liquidityProvider.Amount, 10, 64)

		poolBalance = poolBalance + liquidityProviderAmount

		pool = append(pool, types.Pool{
			Index:   liquidityProvider.Index,
			Account: liquidityProvider.Account,
			Coin:    liquidityProvider.Coin,
			Amount:  liquidityProvider.Amount,
			Fees:    liquidityProvider.Fees,
		})
	}

	for _, liquidityProvider := range pool {
		liquidityProviderAmount, _ := strconv.ParseInt(liquidityProvider.Amount, 10, 64)

		feeToReceive := GetFeeToReceive(fee, poolBalance, liquidityProviderAmount)

		actuallyFees, err := strconv.ParseInt(liquidityProvider.Fees, 10, 64)

		if err != nil {
			//log
			continue
		}

		totalFee := feeToReceive + actuallyFees

		liquidityProvider.Fees = strconv.FormatInt(totalFee, 10)

		k.Keeper.SetPool(ctx, liquidityProvider)
	}

	return nil
}

func GetFeeToReceive(fee, poolBalance, amount int64) int64 {

	factorToReceive := float64(amount) / float64(poolBalance) * 1000000
	feeToReceive := int64(float64(fee)*factorToReceive) / 1000000
	return feeToReceive
}

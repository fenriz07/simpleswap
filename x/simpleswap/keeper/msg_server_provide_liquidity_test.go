package keeper_test

import (
	"fmt"
	"testing"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/types"
)

func TestCalc(t *testing.T) {

	ethCoin := types.NewCoin("eth", math.NewInt(500000))

	var ethPriceUsd int64 = 3317000000 //eth price expressed in USD
	usdToTranfer := ethPriceUsd * ethCoin.Amount.Int64() / 1000000
	//usdStableCoin := sdk.NewCoin("usdt", math.NewInt(usdToTranfer))

	fee := (usdToTranfer * 3000) / 1000000

	fmt.Println(fee)

}

func TestCalcWithCoin(t *testing.T) {

	ethCoin := types.NewCoin("eth", math.NewInt(500000))

	var ethPriceUsd int64 = 3317000000 //eth price expressed in USD
	usdToTranfer := ethCoin.Amount.Mul(math.NewInt(ethPriceUsd)).Sub(math.NewInt(1000000))

	fmt.Println(usdToTranfer)

}

package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/fenriz07/simpleswap/x/simpleswap/keeper"
	"github.com/fenriz07/simpleswap/x/simpleswap/types"
)

func SimulateMsgClaimLiquidity(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgClaimLiquidity{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ClaimLiquidity simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "ClaimLiquidity simulation not implemented"), nil, nil
	}
}

package simpleswap

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/fenriz07/simpleswap/testutil/sample"
	simpleswapsimulation "github.com/fenriz07/simpleswap/x/simpleswap/simulation"
	"github.com/fenriz07/simpleswap/x/simpleswap/types"
)

// avoid unused import issue
var (
	_ = simpleswapsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgProvideLiquidity = "op_weight_msg_provide_liquidity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgProvideLiquidity int = 100

	opWeightMsgSwap = "op_weight_msg_swap"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSwap int = 100

	opWeightMsgClaimLiquidity = "op_weight_msg_claim_liquidity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgClaimLiquidity int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	simpleswapGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&simpleswapGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgProvideLiquidity int
	simState.AppParams.GetOrGenerate(opWeightMsgProvideLiquidity, &weightMsgProvideLiquidity, nil,
		func(_ *rand.Rand) {
			weightMsgProvideLiquidity = defaultWeightMsgProvideLiquidity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgProvideLiquidity,
		simpleswapsimulation.SimulateMsgProvideLiquidity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSwap int
	simState.AppParams.GetOrGenerate(opWeightMsgSwap, &weightMsgSwap, nil,
		func(_ *rand.Rand) {
			weightMsgSwap = defaultWeightMsgSwap
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSwap,
		simpleswapsimulation.SimulateMsgSwap(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgClaimLiquidity int
	simState.AppParams.GetOrGenerate(opWeightMsgClaimLiquidity, &weightMsgClaimLiquidity, nil,
		func(_ *rand.Rand) {
			weightMsgClaimLiquidity = defaultWeightMsgClaimLiquidity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgClaimLiquidity,
		simpleswapsimulation.SimulateMsgClaimLiquidity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgProvideLiquidity,
			defaultWeightMsgProvideLiquidity,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				simpleswapsimulation.SimulateMsgProvideLiquidity(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSwap,
			defaultWeightMsgSwap,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				simpleswapsimulation.SimulateMsgSwap(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgClaimLiquidity,
			defaultWeightMsgClaimLiquidity,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				simpleswapsimulation.SimulateMsgClaimLiquidity(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

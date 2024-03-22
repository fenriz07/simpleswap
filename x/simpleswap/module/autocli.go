package simpleswap

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/fenriz07/simpleswap/api/simpleswap/simpleswap"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "SystemInfo",
					Use:       "show-system-info",
					Short:     "show systemInfo",
				},
				{
					RpcMethod: "StableCoinsWhiteListAll",
					Use:       "list-stable-coins-white-list",
					Short:     "List all stableCoinsWhiteList",
				},
				{
					RpcMethod:      "StableCoinsWhiteList",
					Use:            "show-stable-coins-white-list [id]",
					Short:          "Shows a stableCoinsWhiteList",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

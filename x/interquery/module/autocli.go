package interquery

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "wasmapp/api/wasmapp/interquery"
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
					RpcMethod:      "QueryIcqState",
					Use:            "query-icq-state [sequence]",
					Short:          "Returns the request and response of an ICQ query given the packet sequence",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "sequence"}},
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
				{
					RpcMethod:      "SendQueryAllBalances",
					Use:            "send-query-all-balances [channel-id] [address]",
					Short:          "Query the balances of an account on the remote chain via ICQ",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "channelId"}, {ProtoField: "address"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

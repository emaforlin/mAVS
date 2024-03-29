package mavs

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/emaforlin/mAVS/api/mavs/mavs"
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
					RpcMethod: "StoredVotingAll",
					Use:       "list-stored-voting",
					Short:     "List all storedVoting",
				},
				{
					RpcMethod:      "StoredVoting",
					Use:            "show-stored-voting [id]",
					Short:          "Shows a storedVoting",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "ShowVoter",
					Use:            "show-voter [dni] [voting-id]",
					Short:          "Query showVoter",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "dni"}, {ProtoField: "votingId"}},
				},

				{
					RpcMethod:      "ListVoters",
					Use:            "list-voters [voting-id]",
					Short:          "Query listVoters",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "votingId"}},
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
					RpcMethod:      "CreateVoting",
					Use:            "create-voting [title] [timewindow] [candidates]",
					Short:          "Send a createVoting tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "timewindow"}, {ProtoField: "candidates"}},
				},
				{
					RpcMethod:      "AddVoter",
					Use:            "add-voter [voting-id] [dni] [proof-id]",
					Short:          "Send a add-voter tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "votingId"}, {ProtoField: "dni"}, {ProtoField: "proofId"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

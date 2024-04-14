package codenet

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "codenet/api/codenet/codenet"
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
					RpcMethod:      "EncodeData",
					Use:            "encode-data [data] [encoding-algorithm] [data-size] [checksum] [version] [encoding-proof] [block-number]",
					Short:          "Send a encodeData tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "data"}, {ProtoField: "encodingAlgorithm"}, {ProtoField: "dataSize"}, {ProtoField: "checksum"}, {ProtoField: "version"}, {ProtoField: "encodingProof"}, {ProtoField: "blockNumber"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

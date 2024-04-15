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
				{
					RpcMethod:      "GetEncodedDataById",
					Use:            "get-encoded-data-by-id [encoded-data-id]",
					Short:          "Query getEncodedDataById",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "encodedDataId"}},
				},

				{
					RpcMethod:      "GetProofById",
					Use:            "get-proof-by-id [encoded-data-id]",
					Short:          "Query getProofById",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "encodedDataId"}},
				},

				{
					RpcMethod:      "GetCreatorById",
					Use:            "get-creator-by-id [encoded-data-id]",
					Short:          "Query getCreatorById",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "encodedDataId"}},
				},

				{
					RpcMethod:      "GetVerificationStatusById",
					Use:            "get-verification-status-by-id [encoded-data-id]",
					Short:          "Query getVerificationStatusById",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "encodedDataId"}},
				},

				{
					RpcMethod:      "GetEncodedDataByCreator",
					Use:            "get-encoded-data-by-creator [creator]",
					Short:          "Query getEncodedDataByCreator",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "creator"}},
				},

				{
					RpcMethod:      "GetEncodedDataByTimestamp",
					Use:            "get-encoded-data-by-timestamp [timestamp]",
					Short:          "Query getEncodedDataByTimestamp",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "timestamp"}},
				},

				{
					RpcMethod:      "GetEncodedDataByChecksum",
					Use:            "get-encoded-data-by-checksum [checksum]",
					Short:          "Query getEncodedDataByChecksum",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "checksum"}},
				},

				{
					RpcMethod:      "GetEncodedDataByBlockNumber",
					Use:            "get-encoded-data-by-block-number [block-number]",
					Short:          "Query getEncodedDataByBlockNumber",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "blockNumber"}},
				},

				{
					RpcMethod:      "GetEncodedDataCount",
					Use:            "get-encoded-data-count",
					Short:          "Query getEncodedDataCount",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
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
					Use:            "encode-data [data] [encoding-algorithm] [data-size] [checksum] [version]",
					Short:          "Send a encodeData tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "data"}, {ProtoField: "encodingAlgorithm"}, {ProtoField: "dataSize"}, {ProtoField: "checksum"}, {ProtoField: "version"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}

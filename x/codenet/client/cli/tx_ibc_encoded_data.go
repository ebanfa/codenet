package cli

import (
	"strconv"

	"codenet/x/codenet/types"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	channelutils "github.com/cosmos/ibc-go/v8/modules/core/04-channel/client/utils"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

// CmdSendIbcEncodedData() returns the IbcEncodedData send packet command.
// This command does not use AutoCLI because it gives a better UX to do not.
func CmdSendIbcEncodedData() *cobra.Command {
	flagPacketTimeoutTimestamp := "packet-timeout-timestamp"

	cmd := &cobra.Command{
		Use:   "send-ibc-encoded-data [src-port] [src-channel] [data] [encoded-data-id] [encoding-proof] [block-number] [encoding-algorithm] [data-size] [checksum] [version]",
		Short: "Send a ibcEncodedData over IBC",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress().String()
			srcPort := args[0]
			srcChannel := args[1]

			argCastData := strings.Split(args[2], listSeparator)
			argData := make([]uint64, len(argCastData))
			for i, arg := range argCastData {
				value, err := cast.ToUint64E(arg)
				if err != nil {
					return err
				}
				argData[i] = value
			}
			argEncodedDataId, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}
			argCastEncodingProof := strings.Split(args[4], listSeparator)
			argEncodingProof := make([]uint64, len(argCastEncodingProof))
			for i, arg := range argCastEncodingProof {
				value, err := cast.ToUint64E(arg)
				if err != nil {
					return err
				}
				argEncodingProof[i] = value
			}
			argBlockNumber, err := cast.ToUint64E(args[5])
			if err != nil {
				return err
			}
			argEncodingAlgorithm := args[6]
			argDataSize, err := cast.ToUint64E(args[7])
			if err != nil {
				return err
			}
			argChecksum := args[8]
			argVersion, err := cast.ToUint64E(args[9])
			if err != nil {
				return err
			}

			// Get the relative timeout timestamp
			timeoutTimestamp, err := cmd.Flags().GetUint64(flagPacketTimeoutTimestamp)
			if err != nil {
				return err
			}
			consensusState, _, _, err := channelutils.QueryLatestConsensusState(clientCtx, srcPort, srcChannel)
			if err != nil {
				return err
			}
			if timeoutTimestamp != 0 {
				timeoutTimestamp = consensusState.GetTimestamp() + timeoutTimestamp
			}

			msg := types.NewMsgSendIbcEncodedData(creator, srcPort, srcChannel, timeoutTimestamp, argData, argEncodedDataId, argEncodingProof, argBlockNumber, argEncodingAlgorithm, argDataSize, argChecksum, argVersion)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Uint64(flagPacketTimeoutTimestamp, DefaultRelativePacketTimeoutTimestamp, "Packet timeout timestamp in nanoseconds. Default is 10 minutes.")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

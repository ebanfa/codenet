package keeper

import (
	"context"

	"codenet/x/codenet/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
)

func (k msgServer) SendIbcEncodedData(goCtx context.Context, msg *types.MsgSendIbcEncodedData) (*types.MsgSendIbcEncodedDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.IbcEncodedDataPacketData

	packet.Data = msg.Data
	packet.EncodedDataId = msg.EncodedDataId
	packet.EncodingProof = msg.EncodingProof
	packet.BlockNumber = msg.BlockNumber
	packet.EncodingAlgorithm = msg.EncodingAlgorithm
	packet.DataSize = msg.DataSize
	packet.Checksum = msg.Checksum
	packet.Version = msg.Version

	// Transmit the packet
	_, err := k.TransmitIbcEncodedDataPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendIbcEncodedDataResponse{}, nil
}

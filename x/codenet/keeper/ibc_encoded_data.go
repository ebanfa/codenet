package keeper

import (
	"errors"

	"codenet/x/codenet/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
)

// TransmitIbcEncodedDataPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitIbcEncodedDataPacket(
	ctx sdk.Context,
	packetData types.IbcEncodedDataPacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) (uint64, error) {
	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return 0, errorsmod.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return 0, errorsmod.Wrapf(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: %s", err)
	}

	return k.ibcKeeperFn().ChannelKeeper.SendPacket(ctx, channelCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, packetBytes)
}

// OnRecvIbcEncodedDataPacket processes packet reception
func (k Keeper) OnRecvIbcEncodedDataPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcEncodedDataPacketData) (packetAck types.IbcEncodedDataPacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// TODO: packet reception logic

	return packetAck, nil
}

// OnAcknowledgementIbcEncodedDataPacket responds to the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementIbcEncodedDataPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcEncodedDataPacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.IbcEncodedDataPacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutIbcEncodedDataPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutIbcEncodedDataPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcEncodedDataPacketData) error {

	// TODO: packet timeout logic

	return nil
}

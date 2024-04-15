package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendIbcEncodedData{}

func NewMsgSendIbcEncodedData(
	creator string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	data []byte,
	encodedDataId uint64,
	encodingProof []uint64,
	blockNumber uint64,
	encodingAlgorithm string,
	dataSize uint64,
	checksum string,
	version uint32,
) *MsgSendIbcEncodedData {
	return &MsgSendIbcEncodedData{
		Creator:           creator,
		Port:              port,
		ChannelID:         channelID,
		TimeoutTimestamp:  timeoutTimestamp,
		Data:              data,
		EncodedDataId:     encodedDataId,
		EncodingProof:     encodingProof,
		BlockNumber:       blockNumber,
		EncodingAlgorithm: encodingAlgorithm,
		DataSize:          dataSize,
		Checksum:          checksum,
		Version:           version,
	}
}

func (msg *MsgSendIbcEncodedData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Port == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet port")
	}
	if msg.ChannelID == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet channel")
	}
	if msg.TimeoutTimestamp == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet timeout")
	}
	return nil
}

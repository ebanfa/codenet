package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgEncodeData{}

func NewMsgEncodeData(creator string, data []byte, encodingAlgorithm string, dataSize uint64, checksum string, version uint32) *MsgEncodeData {
	return &MsgEncodeData{
		Creator:           creator,
		Data:              data,
		EncodingAlgorithm: encodingAlgorithm,
		DataSize:          dataSize,
		Checksum:          checksum,
		Version:           version,
	}
}

func (msg *MsgEncodeData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

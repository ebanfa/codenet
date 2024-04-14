package keeper

import (
	"context"

	"codenet/x/codenet/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EncodeData(goCtx context.Context, msg *types.MsgEncodeData) (*types.MsgEncodeDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	var data = types.EncodedData{
		Creator:           msg.Creator,
		DataSize:          msg.DataSize,
		EncodingAlgorithm: msg.EncodingAlgorithm,
		Version:           msg.Version,
	}
	_ = k.AppendEncodedData(
		ctx,
		data,
	)

	return &types.MsgEncodeDataResponse{}, nil
}

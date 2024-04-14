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

	return &types.MsgEncodeDataResponse{}, nil
}

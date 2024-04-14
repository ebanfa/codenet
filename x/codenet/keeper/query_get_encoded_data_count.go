package keeper

import (
	"context"

	"codenet/x/codenet/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetEncodedDataCount(goCtx context.Context, req *types.QueryGetEncodedDataCountRequest) (*types.QueryGetEncodedDataCountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryGetEncodedDataCountResponse{}, nil
}

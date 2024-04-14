package keeper

import (
	"context"

	"codenet/x/codenet/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetEncodedDataById(goCtx context.Context, req *types.QueryGetEncodedDataByIdRequest) (*types.QueryGetEncodedDataByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx
	encodedData, found := k.GetEncodedData(ctx, req.Id)
    if !found {
        return nil, sdkerrors.ErrKeyNotFound
    }

	return &types.QueryGetEncodedDataByIdResponse{
		EncodedData: &encodedData
	}, nil
}

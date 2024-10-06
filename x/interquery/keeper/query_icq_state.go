package keeper

import (
	"context"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"

	"wasmapp/x/interquery/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QueryIcqState(goCtx context.Context, req *types.QueryIcqStateRequest) (*types.QueryIcqStateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	k.Logger().Info("Received sequence ", "sequence", req.Sequence)

	queryRequest, err := k.GetQueryRequest(ctx, req.Sequence)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	var anyQResp *cdctypes.Any
	queryResponse, err := k.GetQueryResponse(ctx, req.Sequence)
	if err == nil {
		anyQResp, err = cdctypes.NewAnyWithValue(&queryResponse)
		if err != nil {
			panic(err)
		}
	}

	anyQReq, err := cdctypes.NewAnyWithValue(&queryRequest)
	if err != nil {
		panic(err)
	}
	return &types.QueryIcqStateResponse{
		Request:  *anyQReq,
		Response: anyQResp,
	}, nil

}

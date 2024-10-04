package keeper

import (
	"context"

	"wasmapp/x/interquery/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendQueryAllBalances(goCtx context.Context, msg *types.MsgSendQueryAllBalances) (*types.MsgSendQueryAllBalancesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSendQueryAllBalancesResponse{}, nil
}

package keeper

import (
	"context"

	"wasmapp/x/token/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendToken(goCtx context.Context, msg *types.MsgSendToken) (*types.MsgSendTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSendTokenResponse{}, nil
}

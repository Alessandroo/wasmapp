package keeper

import (
	"context"
	"time"
	"wasmapp/x/interquery/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
)

func (k msgServer) SendQueryAllBalances(goCtx context.Context, msg *types.MsgSendQueryAllBalances) (*types.MsgSendQueryAllBalancesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.Logger().Info("Receive MsgSendQueryAllBalances message: ", "msg", msg)
	k.Logger().Info("Module port", "port", k.GetPort(ctx))
	k.Logger().Info("ChannelCapabilityPath", "channelPath", host.ChannelCapabilityPath(k.GetPort(ctx), msg.ChannelId))

	chanCap, found := k.ScopedKeeper().GetCapability(ctx, host.ChannelCapabilityPath(k.GetPort(ctx), msg.ChannelId))
	k.Logger().Info("GetCapability", "found", found)
	k.Logger().Info("GetCapability", "chanCap", chanCap)

	if !found {
		return nil, errorsmod.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	q := banktypes.QueryBalanceRequest{
		Address: msg.Address,
		Denom:   "uosmo",
	}
	reqs := []types.RequestQuery{
		{
			Path: "/cosmos.bank.v1beta1.Query/Balance",
			Data: k.cdc.MustMarshal(&q),
		},
	}

	// timeoutTimestamp set to max value with the unsigned bit shifted to satisfy hermes timestamp conversion
	// it is the responsibility of the auth module developer to ensure an appropriate timeout timestamp
	timeoutTimestamp := ctx.BlockTime().Add(time.Minute).UnixNano()
	seq, err := k.SendQuery(ctx, types.PortID, msg.ChannelId, chanCap, reqs, clienttypes.ZeroHeight(), uint64(timeoutTimestamp))
	if err != nil {
		return nil, err
	}

	k.SetQueryRequest(ctx, seq, q)

	return &types.MsgSendQueryAllBalancesResponse{
		Sequence: seq,
	}, nil
}

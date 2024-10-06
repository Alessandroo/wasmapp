package keeper

import (
	"wasmapp/x/interquery/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
)

func (k Keeper) SendQuery(
	ctx sdk.Context,
	sourcePort,
	sourceChannel string,
	channelCap *capabilitytypes.Capability,
	reqs []types.RequestQuery,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) (uint64, error) {
	_, found := k.ibcKeeperFn().ChannelKeeper.GetChannel(ctx, sourcePort, sourceChannel)
	if !found {
		return 0, errorsmod.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", sourcePort, sourceChannel)
	}

	data, err := types.SerializeCosmosQuery(reqs)
	if err != nil {
		return 0, errorsmod.Wrap(err, "could not serialize reqs into cosmos query")
	}
	icqPacketData := types.InterchainQueryPacketData{
		Data: data,
	}

	sequence, err := k.ibcKeeperFn().ChannelKeeper.SendPacket(ctx, channelCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, icqPacketData.GetBytes())

	if err != nil {
		k.Logger().Error("SendInterchainQuery: SendPacket failed", "err", err)
		return 0, err
	}

	k.Logger().Info("SendInterchainQuery: ", "sequence", sequence)

	return sequence, nil
}

func (k Keeper) OnAcknowledgementPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	ack channeltypes.Acknowledgement,
) error {
	k.Logger().Info("Receive OnAcknowledgementPacket", "modulePacket", modulePacket)
	k.Logger().Info("Receive ack", "ack", ack)
	switch resp := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Result:
		var ackData types.InterchainQueryPacketAck
		if err := types.ModuleCdc.UnmarshalJSON(resp.Result, &ackData); err != nil {
			return errorsmod.Wrap(err, "failed to unmarshal interchain query packet ack")
		}
		resps, err := types.DeserializeCosmosResponse(ackData.Data)
		if err != nil {
			return errorsmod.Wrap(err, "could not deserialize data to cosmos response")
		}

		if len(resps) < 1 {
			return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "no responses in interchain query packet ack")
		}

		var r banktypes.QueryBalanceResponse
		if err := k.cdc.Unmarshal(resps[0].Value, &r); err != nil {
			return errorsmod.Wrapf(err, "failed to unmarshal interchain query response to type %T", resp)
		}

		k.SetQueryResponse(ctx, modulePacket.Sequence, r)
		k.SetLastQueryPacketSeq(ctx, modulePacket.Sequence)

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeQueryResult,
				sdk.NewAttribute(types.AttributeKeyAckSuccess, string(resp.Result)),
			),
		)

		k.Logger().Info("interchain query response", "sequence", modulePacket.Sequence, "response", r)
	case *channeltypes.Acknowledgement_Error:
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeQueryResult,
				sdk.NewAttribute(types.AttributeKeyAckError, resp.Error),
			),
		)

		k.Logger().Error("interchain query response", "sequence", modulePacket.Sequence, "error", resp.Error)
	}
	return nil
}

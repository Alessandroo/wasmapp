package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendQueryAllBalances{}

func NewMsgSendQueryAllBalances(creator string, channelId string, address string, pagination string) *MsgSendQueryAllBalances {
	return &MsgSendQueryAllBalances{
		Creator:    creator,
		ChannelId:  channelId,
		Address:    address,
		Pagination: pagination,
	}
}

func (msg *MsgSendQueryAllBalances) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

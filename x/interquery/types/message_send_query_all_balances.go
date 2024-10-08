package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
)

const TypeMsgSendQueryAllBalances = "send_query_all_balances"

var _ sdk.Msg = &MsgSendQueryAllBalances{}

func NewMsgSendQueryAllBalances(creator string, channelId string, address string, pagination *query.PageRequest) *MsgSendQueryAllBalances {
	return &MsgSendQueryAllBalances{
		Creator:    creator,
		ChannelId:  channelId,
		Address:    address,
		Pagination: pagination,
	}
}

func (msg *MsgSendQueryAllBalances) Route() string {
	return RouterKey
}

func (msg *MsgSendQueryAllBalances) Type() string {
	return TypeMsgSendQueryAllBalances
}

func (msg *MsgSendQueryAllBalances) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendQueryAllBalances) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendQueryAllBalances) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendQueryAllBalances{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	// ModuleCdc references the global x/ibc-transfer module codec. Note, the codec
	// should ONLY be used in certain instances of tests and for JSON encoding.
	//
	// The actual codec used for serialization should be provided to x/ibc transfer and
	// defined at the application level.
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func SerializeCosmosQuery(reqs []RequestQuery) (bz []byte, err error) {
	q := &CosmosQuery{
		Requests: reqs,
	}
	return ModuleCdc.Marshal(q)
}

func DeserializeCosmosQuery(bz []byte) (reqs []RequestQuery, err error) {
	var q CosmosQuery
	err = ModuleCdc.Unmarshal(bz, &q)
	return q.Requests, err
}

func SerializeCosmosResponse(resps []ResponseQuery) (bz []byte, err error) {
	r := &CosmosResponse{
		Responses: resps,
	}
	return ModuleCdc.Marshal(r)
}

func DeserializeCosmosResponse(bz []byte) (resps []ResponseQuery, err error) {
	var r CosmosResponse
	err = ModuleCdc.Unmarshal(bz, &r)
	return r.Responses, err
}

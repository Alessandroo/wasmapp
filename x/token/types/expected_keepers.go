package types

import (
	"context"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccountKeeper defines the expected interface for the Account module.
type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI // only used for simulation
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}

// WasmKeeper defines the expected interface needed to cam cosmwasm contracts.
type WasmKeeper interface {
	HasContractInfo(ctx context.Context, contractAddress sdk.AccAddress) bool
	GetContractInfo(ctx context.Context, contractAddress sdk.AccAddress) *wasmtypes.ContractInfo
	Sudo(ctx context.Context, contractAddress sdk.AccAddress, msg []byte) ([]byte, error)
}

//type ContractManagerKeeper interface {
//	AddContractFailure(ctx context.Context, address string, sudoPayload []byte, errMsg string) Failure
//	GetParams(ctx context.Context) (params Params)
//}

package keeper

import (
	"wasmapp/x/token/types"
)

var _ types.QueryServer = Keeper{}

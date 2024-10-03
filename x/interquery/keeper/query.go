package keeper

import (
	"wasmapp/x/interquery/types"
)

var _ types.QueryServer = Keeper{}

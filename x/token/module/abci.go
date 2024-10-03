package token

import (
	"context"
	"wasmapp/x/token/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// BeginBlockSudoMessage Sudo Message called on the contracts
	BeginBlockSudoMessage = `{"send_token_to_contract":{"amount":"1"}}`
	ContractAddress       = "cosmos1zwv6feuzhy6a9wekh96cd57lsarmqlwxdypdsplw6zhfncqw6ftqp82y57"
)

// BeginBlocker executes on contracts at the beginning of the block.
func BeginBlocker(ctx context.Context, k keeper.Keeper) {
	logger := k.Logger()
	logger.Info("Start BeginBlocker")

	contractAddr := sdk.MustAccAddressFromBech32(ContractAddress)
	data, err := k.SudoKeeper().Sudo(ctx, contractAddr, []byte(BeginBlockSudoMessage))
	if err != nil {
		logger.Error("Failed to sudo contract on BeginBlock", "error", err)
		return
	}

	logger.Debug("Data response", "response", data)

	logger.Info("End BeginBlocker")
}

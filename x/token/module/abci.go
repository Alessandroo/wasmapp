package token

import (
	"context"
	"encoding/hex"
	"wasmapp/x/token/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// Sudo Message called on the contracts
	BeginBlockSudoMessage = `{"transfer":{"recipient":"cosmos17lkm7qecuq09h0n2rrgnhjk8duh0uhu2gd046k","amount":"10"}}`
	ContractAddress       = "cosmos14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s4hmalr"
)

// BeginBlocker executes on contracts at the beginning of the block.
func BeginBlocker(ctx context.Context, k keeper.Keeper) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	logger := k.Logger()
	logger.Info("Test BeginBlocker")

	contractAddr := sdk.MustAccAddressFromBech32(ContractAddress)
	data, err := k.SudoKeeper().Sudo(sdkCtx, contractAddr, []byte(BeginBlockSudoMessage))
	if err != nil {
		logger.Debug("Failed to sudo contract on BeginBlock", "error", err)
		return
	}

	if data != nil {
		logger.Info("Data", "data", data)
		logger.Info("Transaction Response", "response", hex.EncodeToString(data))
	}

	logger.Info("End BeginBlocker")
}
